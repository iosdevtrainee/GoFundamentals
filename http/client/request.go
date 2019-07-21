package client

import (
	"io/ioutil"
	"net/http"
)

func sendCustomRequest(url string) (*string, error) {
	customHTTPMethod := "GET"
	req, err := http.NewRequest(customHTTPMethod, url, nil)
	//req.AddCookie()
	//req.BasicAuth
	//req.Cookie
	//req.TLS.Version etc.
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	str := string(body[:])
	return &str, nil
}
