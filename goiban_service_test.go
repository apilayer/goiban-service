/*
The MIT License (MIT)

Copyright (c) 2014 Chris Grieger

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fourcube/goiban-data"
	"github.com/fourcube/goiban-data-loader/loader"

	"encoding/json"

	"strings"

	"github.com/fourcube/goiban"
	"github.com/julienschmidt/httprouter"
)

var server *httptest.Server
var repoSQL data.BankDataRepository

func TestMain(m *testing.M) {
	router := httprouter.New()
	router.GET("/v2/calculate/:countryCode/:bankCode/:accountNumber", calculateAndValidateIBAN)
	router.GET("/calculate/:countryCode/:bankCode/:accountNumber", calculateIBAN)
	router.GET("/validate/:iban", validationHandler)
	router.GET("/countries", countryCodeHandler)
	server = httptest.NewServer(router)

	// repoSQL = data.NewSQLStore("mysql", "root:root@/goiban?charset=utf8")
	repo = data.NewInMemoryStore()
	loader.LoadBundesbankData(loader.DefaultBundesbankPath(), repo)

	retCode := m.Run()
	server.Close()

	os.Exit(retCode)
}

func BenchmarkValidation(b *testing.B) {
	list, _ := readLines("test/iban_test.txt")
	b.ResetTimer()
	var result goiban.ValidationResult

	for _, iban := range list {

		resp, _ := http.Get(server.URL + "/validate/" + iban)
		decoder := json.NewDecoder(resp.Body)

		err = decoder.Decode(&result)
		if err != nil {
			b.Errorf("Expected success %v", err)
		}

		if !result.Valid {
			b.Errorf("Validation failed %v", result)
		}
		resp.Body.Close()
	}
}

func TestSafeContentTypeOnSuccess(t *testing.T) {
	resp, _ := http.Get(server.URL + "/validate/DE89370400440532013000")
	contentType := resp.Header.Get("Content-Type")

	expectedContentType := "application/json; charset=utf-8"

	if contentType != expectedContentType {
		t.Errorf("Content type was %v instead of %v", contentType, expectedContentType)
	}
}

func TestIbanTooShort(t *testing.T) {
	resp, _ := http.Get(server.URL + "/validate/IT96370400440532013000")
	decoder := json.NewDecoder(resp.Body)
	var result goiban.ValidationResult

	err = decoder.Decode(&result)
	if err != nil {
		t.Errorf("Expected success %v", err)
	}

	if !strings.Contains(result.Messages[0], "Expected length") {
		t.Errorf("Expected length error %v", result)
	}

}

// code taken from
// http://stackoverflow.com/a/18479916/1408463
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println("Read", len(lines), "lines of sample data.")
	file.Close()
	return lines, scanner.Err()
}
