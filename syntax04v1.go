package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 난수 추출된 수의 소수 판정 프로그램 v0.1
// 소수 : 1과 자기자신외에는 나누어 떨어지지 않는 수 (0과 1은 제외)
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0
	number := rand.Intn(150) + 2 // 0과 1제외, 2 ~ 151 사이의 수
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}

	if count == 2 {
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다~")
	}
}

// func main() {
// 	// seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	for i := 1; i < 6; i++ {
// 		dice := rand.Intn(6) + 1
// 		fmt.Println(dice)
// 	}
// }
