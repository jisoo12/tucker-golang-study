package main

import (
	"unsafe"
	"reflect"
	"fmt"
)

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceheader := (*reflect.StringHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str:\t%x\n", stringheader.Data)
	fmt.Printf("slice:\t%x\n", sliceheader.Data)
}