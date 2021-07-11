package main

import "fmt"

func main() {
	var a int; // 값을 저장할 변수
	var b int;

	n, err := fmt.Scan(&a, &b); // 입력 두 개 받기 
	if err != nil {
		fmt.Println(n, err) // 에러가 발생 하면 에러 코드 출력
	} else {
		fmt.Println(n, a, b) // 정상 입력되면 입력값 출력
	}
}