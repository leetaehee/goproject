package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int { // op에 따른 함수 타입 변환
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else { // +나 *rk dkslaus nil 반환
		return nil
	}
}

func main() {
	// int 타입 인수 2개를 받아서 int 타입을 반환하는 함수 타입 변수
	var operator func(int, int) int = getOperator("*")

	var result = operator(3, 4) // 함수 타입 변수를 사용해서 함수 호출
	fmt.Println(result)
}
