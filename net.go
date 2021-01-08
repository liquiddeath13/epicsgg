package main

import (
	"io"
	"net/http"
)

func setRequestHeaders(req *http.Request, jwt string) {
	req.Header.Set("x-user-jwt", jwt)
	req.Header.Set("User-Agent", "okhttp/3.12.1")
	req.Header.Set("Content-Type", "application/json")
}

func doRequest(url string, method string, jwt string, body io.Reader) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, body)
	setRequestHeaders(req, jwt)
	res, _ := client.Do(req)
	return res
}
