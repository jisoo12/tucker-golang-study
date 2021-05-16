package main

import (
	"bufio"
	"os"
	"math/rand"
	"time"
	"fmt"
)

const (
	Balance = 1000
	EarnPoint = 500
	LosePoint = 100
	VictoryPoint = 5000
	GameoverPoint = 0
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
	balance := Balance

	for {
		fmt.Print("1~5 사이의 값을 입력해주세요: ")
		n, err := InputIntValue()

		if err != nil {
			fmt.Println("숫자만 입력해주세요.")
		} else if n < 1 || n > 5 {
			fmt.Println("1~5 사이의 값이 아니에요.")
		} else {
			r := rand.Intn(5) + 1
			if n == r {
				balance += EarnPoint
				fmt.Printf("축하합니다. 당첨되었습니다! 잔액: %d원\n", balance)
				
				if balance >= VictoryPoint {
					fmt.Println("게임 승리")
					break
				}
			} else {
				balance -= LosePoint
				fmt.Printf("아쉬워요.. 꽝이에요... 잔액: %d원\n", balance)

				if balance <= GameoverPoint {
					fmt.Println("게임 오버")
					break
				}
			}
		}
	}
}