package storage

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// RequestURL ...
func RequestURL(path string) (string, error) {

	//
	res, err := http.DefaultClient.Get(path)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	//
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	preq := make(map[string]string)

	if err := json.Unmarshal(b, &preq); err != nil {
		return "", err
	}

	return preq["url"], nil
}

// Download ...
func Download(path string) (*http.Response, error) {
	return Request("GET", path, nil)
}

// Upload ...
func Upload(path string, body io.Reader) error {
	_, err := Request("PUT", path, body)
	return err
}

// Request ...
func Request(method, path string, body io.Reader) (*http.Response, error) {

	//
	s3req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	//
	s3res, err := http.DefaultClient.Do(s3req)
	if err != nil {
		return nil, err
	}

	return s3res, nil
}
