package main

import "fmt"

func main() {
	dan := 2
	b := 1
	for {
		for {
			fmt.Printf("%d * %d = %d\n", dan, b, dan*b)
			b++
			if b == 10 {
				break // 안쪽 for문을 종료헙나다,
			}
		}
		b = 1
		dan++
		if dan == 10 {
			break // 바깥쪽 for문을 종료합니다.
		}
	}
	fmt.Println("for문을 종료합니다.")
}
