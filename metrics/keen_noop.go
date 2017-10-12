// +build no_metrics

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
	goiban "github.com/fourcube/goiban"
)

// KeenMetrics is deprecated
type KeenMetrics struct {
	ProjectID   string
	WriteAPIKey string
}

func (keen *KeenMetrics) getEndpoint() string {
	return ""
}

//WriteLogRequest logs to keen.io
//
// http://api.keen.io/3.0/projects/<project_id>/events/<event_collection>
func (keen *KeenMetrics) WriteLogRequest(collectionName string, iban *goiban.Iban) {

}

//LogRequestFromValidationResult unmarshalls the ValidationResult and logs to keen.io
//
//http://api.keen.io/3.0/projects/<project_id>/events/<event_collection>
func (keen *KeenMetrics) LogRequestFromValidationResult(collectionName string, validationResult string) {

}
