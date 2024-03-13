package main

import "fmt"

func main() {
	var input int
	head := 1
	button := 100
	status := 1
	answer := 23
	fmt.Println("猜數字 請輸入一個0~100的數字")
	//fmt.Scan(&input)
	fmt.Println(&input)
	for status > 0 {
		fmt.Scan(&input)
		fmt.Println(fmt.Sprintf("你輸入的數字是 %d", input))
		if input == answer {
			fmt.Println(fmt.Sprintf("猜中了! 答案是 %d", answer))
			status = 0
		}
		if answer > input {
			head = input
		}
		if answer < input {
			button = input
		}

		fmt.Println(fmt.Sprintf("現在的範圍是 %d 到 %d 之間", head, button))

	}

}
