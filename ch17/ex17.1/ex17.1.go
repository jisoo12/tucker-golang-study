package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(5)
	fmt.Println(n+1)
}