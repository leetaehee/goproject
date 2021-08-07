package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 { // slice 길이가 0인지 확인
		fmt.Println("slice in empty", slice)
	}

	slice[1] = 10 // 패닉 발생
	fmt.Println(slice)
}
