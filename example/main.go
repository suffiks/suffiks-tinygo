package main

import "fmt"

func main() {
	ptrAndSize := uint32(3225419836)
	// Inverse of uint32(ptr)<<16 | uint32(len(b))

	ptr := ptrAndSize >> 16
	size := ptrAndSize & 0xffff

	fmt.Println("Size", size)
	fmt.Println("Ptr", ptr)

	if size != 60 {
		panic("wrong size")
	}

	if ptr != 180288 {
		panic("wrong ptr")
	}
}
