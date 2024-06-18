/*
	The default value of a pointer is always nil.
*/

package main

import "log"

func main() {
	randomName := "shubham"
	var name *string
	name = &randomName
	log.Println(name)
}
