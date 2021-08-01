package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 시간 값을 랜덤 시드로 생성

	n := rand.Intn(100)
	fmt.Println(n)
}
