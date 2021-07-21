package main

import "fmt"

func main() {
	a := 1
	b := 1
	found := false
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true // 찾았음을 표시하고 break
				break
			}
		}

		if found {
			break // 바깥쪽 for문에서 찾았는지 검사해서 break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
