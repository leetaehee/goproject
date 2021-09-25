package main

import "fmt"

func sum(nums ...int) int { // 가변 인수를 받는 함수
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)
	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5)) // 인수 5개를 사용합니다.
	fmt.Println(sum(10, 20))        // 인수 2개를 사용합니다.
	fmt.Println(sum())              // 인수 0개를 사용합니다.
}
