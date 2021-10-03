package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare1(t *testing.T) {
	/*
		rst := square(9)
		if rst != 81 {
			t.Errorf("square(9) should be 81 but squre(9) returns %d", rst)
		}
	*/
	assert := assert.New(t)                               // 테스트 객체 생성
	assert.Equal(81, square(9), "sqaure(9) should be 81") // 테스트 함수 호출
}

func TestSquare2(t *testing.T) {
	/*
		rst := square(3)
		if rst != 9 {
			t.Errorf("square(3) should be 9 but square(3) returns %d", rst)
		}
	*/
	assert := assert.New(t)
	assert.Equal(9, square(3), "sqaure(3) should be 9")
}
