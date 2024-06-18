package main

import "log"

/*
	Control Structures (if, for, switch)v
*/

func main() {

	// a typical for loop
	for i := 0; i < 5; i++ {
		log.Println("Statement printed inside a for loop")
	}

	// conditional loop
	i := 0
	for i <= 3 {
		log.Println("inside a loop :", i)
		i++
	}

	// range, value for loop
	arr := []int{1, 2, 3, 4}
	for index, value := range arr {
		log.Println("-> index =", index, "value =", value)
	}

	// infinite for loop
	for {
		log.Println("Statement printed inside an infinite for loop")
	}

}
