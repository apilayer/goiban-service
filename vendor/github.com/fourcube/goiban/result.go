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

package goiban

/*
	Represents the result of an IBAN validation.
*/
type ValidationResult struct {
	Valid        bool            `json:"valid"`
	Messages     []string        `json:"messages"`
	Iban         string          `json:"iban"`
	BankData     BankInfo        `json:"bankData"`
	CheckResults map[string]bool `json:"checkResults"`
}

// Factory method
func NewValidationResult(valid bool, message string, iban string) *ValidationResult {
	messages := []string{}
	if len(message) > 0 {
		messages = append(messages, message)
	}
	return &ValidationResult{valid, messages, iban, *&BankInfo{}, map[string]bool{}}
}

/*
	Represents the result of a parsing attempt.
*/
type ParserResult struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// Factory method
func NewParserResult(valid bool, message string, data string) *ParserResult {
	return &ParserResult{valid, message, data}
}
