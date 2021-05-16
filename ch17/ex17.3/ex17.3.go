package main

import (
	"bufio"
	"os"
	"math/rand"
	"time"
	"fmt"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(100)
	cnt := 0

	for {
		fmt.Print("숫자값을 입력하세요: ")
		n, err := InputIntValue()

		if err != nil {
			fmt.Println("숫자만 입력해주세요.")
		} else {
			cnt++
			
			if n > randNum {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if n < randNum {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Printf("축하합니다. 숫자를 맞추셨습니다. 시도 횟수: %d번", cnt)
				break
			}
		}
	}
}