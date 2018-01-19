package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
}

func (httpClient HttpClient) Post(url string, dataReader io.Reader) []byte {
	var resBodyBytes []byte
	resp, err := http.Post(url,
		"application/json",
		dataReader)
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		} else {
			resBodyBytes = data
		}
	}
	return resBodyBytes
}
