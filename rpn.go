package main

import (
	"fmt"
	// "strconv"
)

func rpn(s string) string {
	var stack []string

	// 標準入力を受け付ける
	if s == "-1" {
		// 標準入力処理
		// ...
		// input = scanner.Text()

		s = "1 3 / 1 6 / +"
	}

	// sを１文字ずつ処理する

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

func print(s string) {
	fmt.Println(s)
}

func main() {
	// テスト用
	input := "1 3 / 1 6 / +"
	// 運用時
	// input := "-1"
	result := rpn(input)
	fmt.Println(result)

}
