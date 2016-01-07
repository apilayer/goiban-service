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

	"github.com/julienschmidt/httprouter"
)

var server *httptest.Server

func TestMain(m *testing.M) {
	router := httprouter.New()
	router.GET("/calculate/:countryCode/:bankCode/:accountNumber", calculateIBAN)
	router.GET("/validate/:iban", validationHandler)
	router.GET("/countries", countryCodeHandler)
	server = httptest.NewServer(router)

	retCode := m.Run()
	server.Close()

	os.Exit(retCode)
}

func BenchmarkValidation(b *testing.B) {
	list, _ := readLines("test/iban_test.txt")
	b.ResetTimer()
	for _, iban := range list {

		resp, _ := http.Get(server.URL + "/validate/" + iban)
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
