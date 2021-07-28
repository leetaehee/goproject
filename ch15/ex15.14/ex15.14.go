package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "Hello World!"
	str2 := str1 // str1 변수값을 str2에 복사

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1)) // Data 값 추출
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2)) // Data 값 추출

	fmt.Println(stringHeader1) // 각 필드 값을 출력합니다.
	fmt.Println(stringHeader2)
}
