package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background()) // 컨테스트 생성
	go PrintEverySencond(ctx)
	time.Sleep(5 * time.Second)
	cancel() // 취소

	wg.Wait()
}

func PrintEverySencond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done(): // 취소 확인
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
