package main

import "fmt"

type PasswordError struct { // 에러 구조체 선언
	Len         int
	RequiredLen int
}

func (err PasswordError) Error() string { // Error() 메서드
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8} // error 반환
	}

	return nil
}

func main() {
	err := RegisterAccount("myId", "myPw") // ID, PW 입력
	if err != nil {                        // 에러 확인
		if errInfo, ok := err.(PasswordError); ok { // 인터페이스 변환
			fmt.Printf("%v Len:%d RequiredLen:%d\n", errInfo, errInfo.Len, errInfo.RequiredLen)
		}
	} else {
		fmt.Println("회원 가입됐습니다.")
	}
}
