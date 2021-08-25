package main

import (
	"bytes"
	"fibonacciservice/rest/client/config"
	"fmt"
	"net/http"
)

func main() {
	cp, err := config.ParseConfig("conf.json")
	if err != nil {
		println(err.Error())
		return
	}
	cals, err := cp.GetCallUrls()
	if err != nil {
		println(err.Error())
		return
	}
	for _, v := range cals {
		resp, e := http.Get(v)
		if e != nil {
			println(e.Error())
			break
		}
		Buf := bytes.NewBuffer(make([]byte, 0, resp.ContentLength))
		_, e = Buf.ReadFrom(resp.Body)
		if e != nil {
			println(e.Error())
			break
		}
		fmt.Printf("Reponse from call %v is: %v \n", v, Buf.String())
	}
	if err != nil {
		println(err.Error())
		return
	}
}
