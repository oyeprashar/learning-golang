package main

import "log"

/*
	1.	Numeric Types:
		•	int: Signed integer type, its size depends on the platform (32 or 64 bits).
		•	uint: Unsigned integer type, its size also depends on the platform.
		•	int8, int16, int32, int64: Signed integers with explicit size in bits.
		•	uint8, uint16, uint32, uint64: Unsigned integers with explicit size in bits.
		•	uintptr: An unsigned integer large enough to store the uninterpreted bits of a pointer value.
		•	float32, float64: Floating-point numbers.
	2.	Boolean Type:
		•	bool: Represents true or false.
	3.	String Type:
		•	string: A sequence of characters, immutable once created.
	4.	Composite Types:
		•	array: A fixed-size collection of items of the same type.
		•	slice: A dynamically-sized, flexible view into the elements of an array.
		•	map: A collection of key-value pairs, also known as an associative array or dictionary.
		•	struct: A composite type that groups together variables under a single name.
	5.	Pointer Types:
		•	pointer: A variable that stores the memory address of another variable.
	6.	Function Types:
		•	func: Represents a function type.
	7.	Interface Types:
		•	interface: Specifies a set of methods that a concrete type must possess to be considered of that interface type.
	8.	Channel Types:
		•	chan: Provides a way for goroutines to communicate with each other.
	9.	Slice Descriptor:
		•	slice: A descriptor that describes the slice’s length, capacity, and a pointer to its underlying array.
	10.	Complex Types:
		•	complex64: Complex number with float32 real and imaginary parts.
		•	complex128: Complex number with float64 real and imaginary parts.
*/

func main() {

	// We will see how to make slices and maps

	// declaring empty slice and appending to it
	var numArr []int
	numArr = append(numArr, 1)
	numArr = append(numArr, 2)
	numArr = append(numArr, 3)
	numArr = append(numArr, 4)

	// declaring and initialising in the same line
	numArr2 := []int{1, 2, 3, 4}
	//log.Println(numArr2)
	numArr2 = append(numArr2, 5)

	// Note: a map cannot be declared like this if we want to dynamically add keys
	// 		The assignment won't give any error but accessing will throw NPE

	// declaring empty map
	//var mp map[string]string
	//mp["key1"] = "value1"
	//log.Println(mp)

	/*
		 Unlike arrays, which have a fixed size, slices, maps, and channels are dynamically sized data structures.
		When you use make(), you’re initializing and allocating memory for these data structures.
	*/

	mp2 := make(map[string]string)
	mp2["key1"] = "value1"
	log.Println(mp2)

}
