package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing URL")
	}

	resp, err := http.Get(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
