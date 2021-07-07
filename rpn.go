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
			stack, num1 = pop(stack)
			stack, num2 = pop(stack)
			stack = push(stack, num2+num1)
			fmt.Println(stack)
			fmt.Println("pls")
		case "-":
			stack, num1 = pop(stack)
			stack, num2 = pop(stack)
			stack = push(stack, num2-num1)
			fmt.Println("mns")
		case "*":
			stack, num1 = pop(stack)
			stack, num2 = pop(stack)
			stack = push(stack, num2*num1)
			fmt.Println("mul")
		case "/":
			stack, num1 = pop(stack)
			stack, num2 = pop(stack)
			stack = push(stack, num2/num1)
			fmt.Println("div")
		default: //数字
			stack = append(stack, v)
			//fmt.Println(stack)
			fmt.Println(v)
		}
	}

	//stack = append(stack, "42")
	return stack[len(stack)-1]
}

func push(stack []string, i int) []string {
	s := strconv.Itoa(i)
	stack = append(stack, s)
	return stack
}

func pop(stack []string) ([]string, int) {
	var s string
	var num int
	s = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	num, _ = strconv.Atoi(s)
	//fmt.Println(num)

	return stack, num
}

// func print(s string) {
// 	fmt.Println(s)
// }

func main() {
	// テスト用
	input := "2 3 + 3 - 4 * 2 /"
	// 運用時
	// input := "-1"
	result := rpn(input)
	fmt.Println(result)

}
