package main

import "fmt"

func main() {
	m := make(map[int]int) // 맵 생성
	m[1] = 0               // 값 추가
	m[2] = 2
	m[3] = 3

	delete(m, 3)      // 요소 삭제
	delete(m, 4)      // 없는 요소 삭제 시도
	fmt.Println(m[3]) // 삭제된 요소값 출력
	fmt.Println(m[1]) // 존재하는 요소값 출력

	_, m3ok := m[3]
	_, m4ok := m[4]

	fmt.Println(m3ok)
	fmt.Println(m4ok)
}
