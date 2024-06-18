package main

import (
	"fmt"
	"log"
	"sync"
)

/*
	How wait groups work?
		1. Assign a delta to wg variable, lets say 2
		2. Using wg.Done() we can decrement the delta
		3. wg.Wait() keeps the program flow waiting at this line until the delta decrements to zero
*/

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	nameCh := make(chan string, 2)

	go func(s string) {
		defer wg.Done()
		nameCh <- s
	}("hello")

	go func(s string) {
		defer wg.Done()
		nameCh <- s
	}("world")

	var name1 string
	var name2 string

	name1 = <-nameCh
	name2 = <-nameCh

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All goroutines finished")
	log.Println("----> value of name1 = ", name1, "value of name2 =", name2)
}
