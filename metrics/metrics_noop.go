// +build no_metrics

package metrics

import (
	"net/http"

	"github.com/fourcube/goiban"
)

type Event struct {
	Country string
}

//
type MetricsRegister interface {
	Register(Event)
}

type InmemMetricsRegister struct {
}

func NewInmemMetricsRegister() *InmemMetricsRegister {
	return &InmemMetricsRegister{}
}

func (imr *InmemMetricsRegister) Register(e Event) {
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
}
