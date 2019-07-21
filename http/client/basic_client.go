package clients

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// needs to return pointer instead of type
// otherwise cannot return nil will expect zero value

func fetchResource(url string) (*string, error) {
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	str := bytes.NewBuffer(body).String()
	return &str, nil
}
