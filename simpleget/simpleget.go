package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// 文字列で "200 OK"
	log.Println("Status:", resp.Status)
	// 数値で
	log.Println("StatusCode:", resp.StatusCode)

	log.Println(string(body))
}
