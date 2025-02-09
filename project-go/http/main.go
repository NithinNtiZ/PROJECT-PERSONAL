package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://ifconfig.me/all.json")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("%s", body)
}
