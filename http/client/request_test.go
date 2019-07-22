package client

import (
	"fmt"
	"testing"
)

func TestCustomHTTPRequest(t *testing.T) {
	sendCustomHTTPRequest(t, "GET", "https://raw.githubusercontent.com/aionate0812/bundle_backend/master/config.js",
		`module.exports = {dba:"postgres:localhost:5432/bundle", DARKSKY_API_KEY:"aa9770871dcf844f5ca90bbdcd087e5b"}`)
}
func sendCustomHTTPRequest(t *testing.T, method, url, expResult string) {
	bodyString, err := sendCustomRequest(method, url)
	if err != nil {
		t.Error(err)
	}
	if *bodyString != expResult {
		t.Fatal(fmt.Sprintf("url: %s did not provide expected result: %s instead received %s", url, expResult, *bodyString))
	}
}
