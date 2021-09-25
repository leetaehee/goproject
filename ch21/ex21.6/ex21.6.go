package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3) // 함수 리터럴 3개를 가진 슬라이스
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i) // 캡쳐된 i값 출력
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v) // 캡쳐된 v값 출력
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}
