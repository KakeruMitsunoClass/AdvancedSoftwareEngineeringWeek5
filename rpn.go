package main

import (
	"fmt"
	"strconv"
	"strings"
)

func rpn(s string) string {
	var stack []string
	var num1 int
	var num2 int

	// 標準入力を受け付ける
	if s == "-1" {
		// 標準入力処理
		// ...
		// input = scanner.Text()

		s = "1 3 / 1 6 / +"
	}

	// sを１文字ずつ処理する

	/*
			for _, value := range slice {
		    	// fmt.Println( _ )
		    	fmt.Println( value )
		    }
	*/

	splitteds := strings.Split(s, " ")
	for _, v := range splitteds {
		// fmt.Println(v)

		switch v {
		case "+":
			stack, snum1 := pop(stack)
			num1, _ = strconv.Atoi(snum1)
			stack, snum2 := pop(stack)
			num2, _ = strconv.Atoi(snum2)
			fmt.Println(num1)
			fmt.Println(num2)
			fmt.Println("pls")
		case "-":
			stack, snum1 := pop(stack)
			num1, _ = strconv.Atoi(snum1)
			stack, snum2 := pop(stack)
			num2, _ = strconv.Atoi(snum2)
			fmt.Println("mns")
		case "*":
			stack, snum1 := pop(stack)
			num1, _ = strconv.Atoi(snum1)
			stack, snum2 := pop(stack)
			num2, _ = strconv.Atoi(snum2)
			fmt.Println("mul")
		case "/":
			stack, snum1 := pop(stack)
			num1, _ = strconv.Atoi(snum1)
			stack, snum2 := pop(stack)
			num2, _ = strconv.Atoi(snum2)
			fmt.Println("div")
		default: //数字
			stack = push(stack, v)
			fmt.Println(v)
		}
	}

	stack = append(stack, "42")
	return stack[len(stack)-1]
}

func push(stack []string, s string) []string {
	stack = append(stack, s)
	return stack
}

func pop(stack []string) ([]string, string) {
	var s string
	s = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return stack, s
}

// func print(s string) {
// 	fmt.Println(s)
// }

func main() {
	// テスト用
	input := "1 2 + 1 - 2 * 1 /"
	// 運用時
	// input := "-1"
	result := rpn(input)
	fmt.Println(result)

}
