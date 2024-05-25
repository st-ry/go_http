package main

import (
	"io"
	"log"
	"net/http"
)

func main(){
	resp, err := http.Get("http://localhost")
	if err != nil {
		panic(err)
	}
	
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	
	if err != nil {
		panic(err)
	}
	
	log.Println(string(body))
}