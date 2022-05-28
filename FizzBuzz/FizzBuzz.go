package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		str := ""

		// ここから記述
		if i%3 == 0 {
			str = "Fizz" //3の倍数のときFizzをstrに代入
		}
		if i%5 == 0 {
			str += "Buzz" //5の倍数のときBuzzをstrに連結する
		}
		if str == "" {
			str = strconv.Itoa(i) //strが空ならiをstring型に変換して代入
		}
		// ここまで記述

		fmt.Println(str)
	}

}
