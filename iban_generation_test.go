package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/fourcube/goiban"
)

func TestGenerateIBANIgnoreCountryCase(t *testing.T) {
	respA, err := http.Get(server.URL + "/calculate/BE/539/007547034")

	if err != nil {
		t.Errorf("failed to generate iban %v", err)
		t.FailNow()
	}

	respB, err := http.Get(server.URL + "/calculate/be/539/007547034")

	if err != nil {
		t.Errorf("failed to generate iban %v", err)
		t.FailNow()
	}

	var resA goiban.ParserResult
	data, _ := ioutil.ReadAll(respA.Body)
	json.Unmarshal(data, &resA)

	var resB goiban.ParserResult
	data, _ = ioutil.ReadAll(respB.Body)
	json.Unmarshal(data, &resB)

	if !resA.Valid || !resB.Valid {
		t.Errorf("expected request to succeed")
	}

	if resA.Data != resB.Data {
		t.Errorf("expected case insensitivity")
	}
}

func TestGenerateIBANInvalidCountryCode(t *testing.T) {
	resp, err := http.Get(server.URL + "/calculate/12/539/007547034")

	if err != nil {
		t.Errorf("failed to generate iban %v", err)
		t.FailNow()
	}

	var res CalculateError
	data, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(data, &res)

	if res.Valid {
		t.Errorf("expected request to fail")
	}
}

func TestGenerateIBANTooMuchData(t *testing.T) {
	resp, err := http.Get(server.URL + "/calculate/BE/539/007547111111111111111111")

	if err != nil {
		t.Errorf("failed to generate iban %v", err)
		t.FailNow()
	}

	var res goiban.ParserResult
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &res)

	if res.Valid {
		t.Errorf("expected request to fail")
	}
}

func TestGenerateIBANV2WithValidation(t *testing.T) {
	resp, err := http.Get(server.URL + "/v2/calculate/DE/1/9?getBIC=true&validateBankCode=true")

	if err != nil {
		t.Errorf("failed to generate iban %v", err)
		t.FailNow()
	}

	var res goiban.ValidationResult
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &res)

	if res.Valid {
		t.Errorf("expected request to fail")
	}
}

func TestGenerateIBANV2NoValidation(t *testing.T) {
	resp, err := http.Get(server.URL + "/v2/calculate/DE/1/9")

	if err != nil {
		t.Errorf("failed to generate iban %v", err)
		t.FailNow()
	}

	var res goiban.ValidationResult
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &res)

	if res.Iban != "DE0819" {
		t.Errorf("expected returned iban to equal DE0819")
	}

	if res.Valid {
		t.Errorf("expected request to fail")
	}
}

func TestGenerateIBANV2GetBIC(t *testing.T) {
	resp, err := http.Get(server.URL + "/v2/calculate/DE/37040044/0532013000?getBIC=true")

	if err != nil {
		t.Errorf("failed to generate iban %v", err)
		t.FailNow()
	}

	var res goiban.ValidationResult
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &res)

	if res.Iban != "DE89370400440532013000" {
		t.Errorf("expected returned iban to equal DE89370400440532013000, was " + res.Iban)
	}

	if res.BankData.Bic != "COBADEFFXXX" {
		t.Errorf("expected returned bic to equal COBADEFFXXX, was " + res.BankData.Bic)
	}

	if !res.Valid {
		t.Errorf("expected request to succeed")
	}
}
