package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func get() {
	values := url.Values{
		"query": {"hello world"},
	}
	resp, err := http.Get("http://localhost:18888" + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	log.Println("Headers:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}

func head() {
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
}
