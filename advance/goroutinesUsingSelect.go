package main

import (
	"log"
)

func main() {

	nameCh := make(chan string)

	// This can be an API call that gets the data and pushes it to the chanel
	go func(s string) {
		nameCh <- s
	}("world")

	// This can be an API call that gets the data and pushes it to the chanel
	go func(s string) {
		nameCh <- s
	}("hello")

	var name1 string
	var name2 string

	select {
	case name1 = <-nameCh:
	}

	select {
	case name2 = <-nameCh:
	}

	log.Println("first value retrieved from the channel is :", name1)
	log.Println("second value retrieved from the channel is :", name2)

}
