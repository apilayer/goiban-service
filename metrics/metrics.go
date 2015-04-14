package metrics

import (
	"encoding/json"
	"net/http"
	"time"

	gm "github.com/armon/go-metrics"
	"github.com/fourcube/goiban"
)

type Event struct {
	Country string
}

//
type MetricsRegister interface {
	Register(Event)
	Data() []*gm.IntervalMetrics
}

type InmemMetricsRegister struct {
	*gm.InmemSink
}

func NewInmemMetricsRegister() *InmemMetricsRegister {
	return &InmemMetricsRegister{
		gm.NewInmemSink(5*time.Minute, 24*7*time.Hour),
	}
}

func (imr *InmemMetricsRegister) Register(e Event) {
	imr.IncrCounter([]string{e.Country}, 1.0)
}

func IbanToEvent(iban *goiban.Iban) Event {
	return Event{
		Country: iban.GetCountryCode(),
	}
}

func ValidationResultToEvent(validationResult *goiban.ValidationResult) Event {
	return Event{
		Country: goiban.ExtractCountryCode(validationResult.Iban),
	}
}

func (imr *InmemMetricsRegister) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	// Allow CORS
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder.Encode(imr.Data())
}
