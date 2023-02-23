package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func main() {
	c := http.Client{}
	json := bytes.NewBuffer([]byte(`{"name": "Gabriel"}`))
	res, err := c.Post("http://www.google.com", "application/json", json)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
