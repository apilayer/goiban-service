package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestReturnsCountryCodeMap(t *testing.T) {
	resp, err := http.Get(server.URL + "/countries")

	if err != nil {
		t.Errorf("failed to get countries %v", err)
		t.FailNow()
	}

	var res map[string]string
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &res)

	if len(res) < 1 {
		t.Errorf("Received no country codes")
	}
}
