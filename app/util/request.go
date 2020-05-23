package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// request get method
func Get(rawurl string, params map[string]string) (data interface{}, err error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		logger.Error("Get url parse error:", err)
		return
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	reqURL := u.String()
	resp, err := http.Get(reqURL)
	if err != nil {
		logger.Error("http GetGet reqURL error:", err)
		return
	}
	if err = handleResult(resp, &data); err != nil {
		logger.Error("Get handleResult error:", handleResult)
		return
	}
	return
}

// handle http result
func handleResult(r *http.Response, data *interface{}) error {
	Body := r.Body
	defer Body.Close()

	body, err := ioutil.ReadAll(Body)
	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("statusCode: %d ; body: %s", r.StatusCode, string(body))
	}

	return json.Unmarshal(body, &data)
}
