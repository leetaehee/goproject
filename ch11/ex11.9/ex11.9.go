package main

import "fmt"

func find45(a int) (int, bool) { // 곱해서 45가 되는 값을 찾음
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}
	return 0, false
}

func main() {
	a := 1
	b := 0

	for ; a <= 9; a++ {
		var found bool
		if b, found = find45(a); found { // 함수
			break
		}
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
