package io

import (
	"io"
	"net/http"

	"github.com/exact/elle/random"
)

func Get(target string, headers map[string]string, defaults bool) (string, error) {
	req, _ := http.NewRequest("GET", target, nil)

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
