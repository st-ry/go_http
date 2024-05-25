package main

import (
	"io"
	"log"
	"net/http"
)

func main(){
	resp, err := http.Get("http://localhost:8888")
	if err != nil {
		panic(err)
	}
	
	defer resp.Body.Close()
	printHttpResponse(resp)
}

func printHttpResponse (response *http.Response){
	log.Println("Status:", response.Status)
	log.Println("StatusCode:", response.StatusCode)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	log.Println("Body:",string(body))
	log.Println("Fields:", response.Header)
	log.Println("Content-Length:", response.Header.Get("Content-Length"))
}