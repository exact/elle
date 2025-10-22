package io

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/exact/elle/random"
)

func Request(method, target string, headers, data map[string]string, defaults bool) (string, error) {
	var req *http.Request
	if data == nil {
		req, _ = http.NewRequest(method, target, nil)
	} else {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return "", err
		}

		req, _ = http.NewRequest(method, target, bytes.NewBuffer(jsonData))
	}

	if defaults {
		for k, v := range random.NewHeaders() {
			req.Header.Set(k, v)
		}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
