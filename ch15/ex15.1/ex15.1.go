package main

import "fmt"

func main() {
	// 큰 따옴표로 묶으면 특수 문자가 동작합니다.
	str1 := "Hello\t'world'\n"

	// 백쿼트로 묶으면 특수 문자가 동작하지 않습니다.
	str2 := `Go is "awesome!\nGo is simple and\t'powerful'`

	fmt.Println(str1)
	fmt.Println(str2)
}
