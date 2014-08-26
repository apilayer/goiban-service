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
package metrics

import (
	"encoding/json"
	"log"

	goiban "github.com/fourcube/goiban"
	"github.com/franela/goreq"
)

//KeenMetrics
type KeenMetrics struct {
	ProjectID   string
	WriteAPIKey string
}

type logEntry struct {
	Country string
}

func (keen *KeenMetrics) getEndpoint() string {
	return "http://api.keen.io/3.0/projects/" + keen.ProjectID + "/events/"
}

func ibanToLogEntry(iban *goiban.Iban) *logEntry {
	return &logEntry{
		Country: iban.GetCountryCode(),
	}
}

func validationResultToLogEntry(validationResult *goiban.ValidationResult) *logEntry {
	return &logEntry{
		Country: goiban.ExtractCountryCode(validationResult.Iban),
	}
}

//WriteLogRequest logs to keen.io
//
// http://api.keen.io/3.0/projects/<project_id>/events/<event_collection>
func (keen *KeenMetrics) WriteLogRequest(collectionName string, iban *goiban.Iban) {
	var url = keen.getEndpoint() + collectionName

	req := goreq.Request{
		Method:      "POST",
		Uri:         url,
		ContentType: "application/json",
		Body:        ibanToLogEntry(iban),
	}

	req.AddHeader("Authorization", keen.WriteAPIKey)

	res, err := req.Do()

	if err != nil {
		log.Printf("Error while posting stats: %v", err)
		return
	}

	// Close the response body
	if res.Body != nil {
		defer res.Body.Close()
	}

	if collectionName == "Test" {
		log.Printf(url)
		text, _ := res.Body.ToString()
		log.Printf("Response (%v): %v", res.StatusCode, text)
	}

}

//LogRequestFromValidationResult unmarshalls the ValidationResult and logs to keen.io
//
//http://api.keen.io/3.0/projects/<project_id>/events/<event_collection>
func (keen *KeenMetrics) LogRequestFromValidationResult(collectionName string, validationResult string) {
	var url = keen.getEndpoint() + collectionName

	var result goiban.ValidationResult
	json.Unmarshal([]byte(validationResult), &result)

	req := goreq.Request{
		Method:      "POST",
		Uri:         url,
		ContentType: "application/json",
		Body:        validationResultToLogEntry(&result),
	}

	req.AddHeader("Authorization", keen.WriteAPIKey)

	res, err := req.Do()

	if err != nil {
		log.Printf("Error while posting stats: %v", err)
	}

	// Close the response body
	if res.Body != nil {
		defer res.Body.Close()
	}

	if collectionName == "Test" {
		log.Printf(url)
		text, _ := res.Body.ToString()
		log.Printf("Response (%v): %v", res.StatusCode, text)
	}

}
