package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("제곱근은 양수여야 합니다. f:%g", f) // f가 음수이면 에러 반환
		/*
		 * 사용시 errors 패키지 import 할 것..
		 */
		//return 0, errors.New("제곱근은 양수여야 합니다.")
	}
	return math.Sqrt(f), nil
}

func main() {
	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("Error: %v\n", err) // 에러 출력
		return
	}
	fmt.Printf("Sqrt(-2) = %v\n", sqrt)
}
