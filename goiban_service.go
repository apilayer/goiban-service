package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/fourcube/goiban-data"
	"github.com/fourcube/goiban-data-loader/loader"

	"github.com/fourcube/goiban"
	m "github.com/fourcube/goiban-service/metrics"
	"github.com/julienschmidt/httprouter"
	"github.com/pmylund/go-cache"
	"github.com/rs/cors"
)

/**
Handles requests and serves static pages.

route							description
--------------------			--------------------------------------------------------
/validate/{iban} 				Tries to validate {iban} and returns a HTTP response
								in JSON. See goiban.ValidationResult for details of the
								data returned.

/*								Renders static content from the "./static" folder
*/
var (
	c   = cache.New(5*time.Minute, 30*time.Second)
	err error

	metrics      *m.KeenMetrics
	inmemMetrics = m.NewInmemMetricsRegister()
	repo         data.BankDataRepository
	// Flags
	dataPath   string
	staticPath string
	mysqlURL   string
	port       string
	help       bool
	web        bool
)

func init() {
	flag.StringVar(&dataPath, "dataPath", "", "Base path of the bank data")
	flag.StringVar(&staticPath, "staticPath", "", "Base path of the static web content")
	flag.StringVar(&mysqlURL, "dbUrl", "", "Database connection string")

	flag.StringVar(&port, "port", "8080", "HTTP Port or interface to listen on")
	flag.BoolVar(&help, "h", false, "Show usage")
	flag.BoolVar(&web, "w", false, "Serve staticPath folder")
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	if web && staticPath == "" {
		// Try to serve from the package src directory
		path := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "fourcube", "goiban-service", "static")
		f, err := os.Open(path)
		defer f.Close()

		if err != nil {
			log.Fatalf("Cannot serve static content from %s: %v. Please set a correct folder with the -staticPath option.", path, err)
		}

		staticPath = path
	}

	listen()
}

func listen() {
	if mysqlURL != "" {
		log.Printf("Using SQL data store.")
		repo = data.NewSQLStore("mysql", mysqlURL)
	} else {
		log.Printf("Using in-memory data store.")
		repo = data.NewInMemoryStore()

		if dataPath != "" {
			loader.SetBasePath(dataPath)
		}
		loader.LoadBundesbankData(loader.DefaultBundesbankPath(), repo)
		loader.LoadAustriaData(loader.DefaultAustriaPath(), repo)
		loader.LoadBelgiumData(loader.DefaultBelgiumPath(), repo)
		loader.LoadLiechtensteinData(loader.DefaultLiechtensteinPath(), repo)
		loader.LoadLuxembourgData(loader.DefaultLuxembourgPath(), repo)
		loader.LoadNetherlandsData(loader.DefaultNetherlandsPath(), repo)
		loader.LoadSwitzerlandData(loader.DefaultSwitzerlandPath(), repo)
	}

	router := httprouter.New()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET"},
	})

	router.GET("/validate/:iban", validationHandler)
	router.GET("/countries", countryCodeHandler)
	router.GET("/calculate/:countryCode/:bankCode/:accountNumber", calculateIBAN)
	router.GET("/v2/calculate/:countryCode/:bankCode/:accountNumber", calculateAndValidateIBAN)
	router.Handler("GET", "/metrics", http.Handler(inmemMetrics))

	//Only host the static template when the ENV is 'Live' or 'Test'
	if web {
		log.Printf("Serving static content from folder %s.", staticPath)
		router.NotFound = http.FileServer(http.Dir(staticPath))
	}

	listeningInfo := "Listening on %s"
	handler := corsHandler.Handler(router)

	if strings.ContainsAny(port, ":") {
		if web {
			listeningInfo = fmt.Sprintf(listeningInfo, "%s (serving static content from '/').")
		}
		log.Printf(listeningInfo, port)
		err = http.ListenAndServe(port, handler)
	} else {
		if web {
			listeningInfo = fmt.Sprintf(listeningInfo, ":%s (serving static content from '/').")
		}
		log.Printf(listeningInfo, port)
		err = http.ListenAndServe(":"+port, handler)
	}

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Processes requests to the /validate/ url
func validationHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var strRes string
	config := map[string]bool{}
	// Set response type to application/json.
	// See: https://www.owasp.org/index.php/XSS_(Cross_Site_Scripting)_Prevention_Cheat_Sheet#RULE_.233.1_-_HTML_escape_JSON_values_in_an_HTML_context_and_read_the_data_with_JSON.parse
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// extract iban parameter
	iban := ps.ByName("iban")

	// check for additional request parameters
	validateBankCodeQueryParam := r.FormValue("validateBankCode")
	config["validateBankCode"] = toBoolean(validateBankCodeQueryParam)

	// check for additional request parameters
	getBicQueryParam := r.FormValue("getBIC")
	config["getBIC"] = toBoolean(getBicQueryParam)

	// hit the cache
	value, found := hitCache(iban + strconv.FormatBool(config["getBIC"]) + strconv.FormatBool(config["validateBankCode"]))
	if found {
		go logFromCacheEntry("", value)
		fmt.Fprintf(w, value)
		return
	}

	// no value for request parameter
	// return HTTP 400
	if len(iban) == 0 {
		res, _ := json.MarshalIndent(goiban.NewValidationResult(false, "Empty request.", iban), "", "  ")
		strRes = string(res)
		w.Header().Add("Content-Length", strconv.Itoa(len(strRes)))
		// put to cache and render
		// c.Set(iban, strRes, 0)
		http.Error(w, strRes, http.StatusBadRequest)
		return
	}

	// IBAN is not parseable
	// return HTTP 200
	parserResult := goiban.IsParseable(iban)

	if !parserResult.Valid {
		res, _ := json.MarshalIndent(goiban.NewValidationResult(false, "Cannot parse as IBAN: "+parserResult.Message, iban), "", "  ")
		strRes = string(res)
		w.Header().Add("Content-Length", strconv.Itoa(len(strRes)))

		// put to cache and render
		c.Set(iban+strconv.FormatBool(config["getBIC"])+strconv.FormatBool(config["validateBankCode"]), strRes, 0)
		fmt.Fprintf(w, strRes)
		return
	}

	// Try to validate
	parsedIban := goiban.ParseToIban(iban)
	result := parsedIban.Validate()

	// intermediate result
	if len(config) > 0 {
		result = additionalData(parsedIban, result, config)
	}

	res, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	strRes = string(res)
	w.Header().Add("Content-Length", strconv.Itoa(len(strRes)))
	// put to cache and render

	go logFromIbanResult("", parsedIban)

	key := iban + strconv.FormatBool(config["getBIC"]) + strconv.FormatBool(config["validateBankCode"])

	c.Set(key, strRes, 0)
	fmt.Fprintf(w, strRes)
	return
}

func toBoolean(value string) bool {
	switch value {
	case "1":
		return true
	case "true":
		return true
	default:
		return false
	}
}

func additionalData(iban *goiban.Iban, intermediateResult *goiban.ValidationResult, config map[string]bool) *goiban.ValidationResult {
	validateBankCode, ok := config["validateBankCode"]
	if ok && validateBankCode {
		intermediateResult = goiban.ValidateBankCode(iban, intermediateResult, repo)
	}

	getBic, ok := config["getBIC"]
	if ok && getBic {
		intermediateResult = goiban.GetBic(iban, intermediateResult, repo)
	}
	return intermediateResult
}

func hitCache(iban string) (string, bool) {
	val, ok := c.Get(iban)
	if ok {
		return val.(string), ok
	}

	return "", false

}

// Only logs when metrics is defined
func logFromCacheEntry(ENV string, value string) {
	if metrics != nil {
		metrics.LogRequestFromValidationResult(ENV, value)
	} else {
		var result *goiban.ValidationResult
		json.Unmarshal([]byte(value), &result)

		inmemMetrics.Register(m.ValidationResultToEvent(result))
	}
}

// Only logs when metrics is defined
func logFromIbanResult(ENV string, value *goiban.Iban) {
	if metrics != nil {
		metrics.WriteLogRequest(ENV, value)
	} else {
		inmemMetrics.Register(m.IbanToEvent(value))
	}
}
