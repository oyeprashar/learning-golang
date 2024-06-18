package main

import "log"

func main() {

	ch1 := make(chan string, 4)

	go func() {
		ch1 <- "Shubham"
		ch1 <- "Shikha"
		ch1 <- "Sumi"
		ch1 <- "Ashank"
		close(ch1) // if we don't close the channel, then the for loop will throw a deadlock error
	}()

	for value := range ch1 {
		log.Println("-> value =", value)
	}

}
