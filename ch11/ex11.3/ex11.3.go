package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for { // 무한 루프
		fmt.Print("입력하세요.")
		var number int
		_, err := fmt.Scanln(&number) // 한 줄 입력을 받습니다.
		if err != nil {               // 숫자가 아닌 경우
			fmt.Println("숫자를 입력하세요.")

			// 키보드 버퍼를 비웁니다.
			stdin.ReadString('\n') // 키보드 버퍼를 지워줍니다.
			continue               // 1로 돌아갑니다.
		}
		fmt.Printf("입력하신 숫자는 %d입니다.\n", number)
		if number%2 == 0 { // 짝수 검사를 합니다.
			break // for문을 종료합니다.
		}
	}
	fmt.Println("for문이 종료됐습니다.")

}
