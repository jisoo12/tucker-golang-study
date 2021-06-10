package main

import "fmt"

type PasswordError struct {
	len					int
	RequireLen	int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{ len(password), 8 }
	}
	
	return nil
}

func main() {
	err := RegisterAccount("myID", "myPw")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequireLen:%d\n", errInfo, errInfo.len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입이 되었습니다.")
	}
}