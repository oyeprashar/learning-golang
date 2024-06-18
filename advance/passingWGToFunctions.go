package main

import (
	"fmt"
	"sync"
)

/*
func main() {
	go a()
	go b()
}

*/

func a(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("a")
}
func b(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("b")
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go a(&wg)
	go b(&wg)

	wg.Wait()
}
