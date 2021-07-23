package main

import "fmt"

type User struct { // 일반 고객용 구조체
	Name string
	ID   string
	Age  int
}

type VIPUser struct { // VIP 고객용 구조체
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 40},
		3,
		250, // 여러 줄로 초기화 할 때는 제일 마지막 값 뒤에 꼭 쉼표를 달아주세요,
	} // User를 포함한 VIPUser 구조체 변수를 초기화 합니다.

	fmt.Printf("유저: %s, ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s, ID: %s, 나이: %d VIP 레벨: %d VIP 가격: %d 만 원\n",
		vip.UserInfo.Name, // UserInfo 안의 Name
		vip.UserInfo.ID,   // UserInfo 안의 ID
		vip.UserInfo.Age,  // UserInfo 안의 Age
		vip.VIPLevel,      // UserInfo 안의 VIPLevel
		vip.Price,         // 마지막에 쉼표
	)
}
