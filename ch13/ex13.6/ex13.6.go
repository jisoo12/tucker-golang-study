package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	age int32
	Score float64
}

func main() {
	user := User{ 23, 77.2}
	fmt.Println(unsafe.Sizeof(user))
}