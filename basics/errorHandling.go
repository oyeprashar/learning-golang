package main

import (
	"log"
)

type response struct {
}

func someAPICall() (response, error) {

	/*
		This will make an API call
	*/

	return response{}, nil
}

func main() {

	resp, err := someAPICall()

	if err != nil {
		// TODO: we need to stop the function execution here and process it
		return
	}

	log.Println(resp)
}
