package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/micro/go-micro/v2/logger"
)

// request get method
func Get(rawurl string, params map[string]string) (data map[string]interface{}, err error) {
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
		logger.Error("http Get reqURL error:", err)
		return
	}
	if err = handleResult(resp, &data); err != nil {
		logger.Error("Get handleResult error:", err)
		return
	}
	return
}

// request post method
func Post(url string, params interface{}) (data map[string]interface{}, err error) {
	body, err := json.Marshal(params)
	if err != nil {
		logger.Error("http Post json.Marshal error:", err)
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		logger.Error("http Post url error:", err)
		return
	}

	if err = handleResult(resp, &data); err != nil {
		logger.Error("Post handleResult error:", err)
		return
	}
	return
}

// handle http result
func handleResult(r *http.Response, data *map[string]interface{}) error {
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
