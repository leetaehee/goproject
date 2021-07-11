package main

import "fmt"

func main() {
	var a int
	var b int
	
	n, err := fmt.Scanln(&a, &b) // 값을 입력받습니다.
	if err != nil { // 에러발생 시
		fmt.Println(n, err) // 에러를 출력합니다.
	} else {
		fmt.Println(n, a, b)
	}
}