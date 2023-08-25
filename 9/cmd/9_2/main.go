/*
逆ポーランド記法の実装
3 4 + 1 2 - *
逆ポーランド記法の入力を(3+4)*(1-2)として計算する
*/
package main

import (
	"fmt"
	"log"
)

var (
	operations = map[rune]bool{
		'+': true,
		'-': true,
		'*': true,
	}
)

type stackTypes interface {
	~int16 | ~int32 | ~int64 | ~int
}

func pop[T stackTypes](stack []T) ([]T, T) {
	last := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return stack, last
}

func calc(s string) (int, error) {
	var (
		stack  []int
		n1, n2 int
	)

	for _, r := range s {
		if r >= '0' && r <= '9' {
			// 文字列リテラルが数値
			stack = append(stack, int(r-'0'))
			continue
		}

		// 文字列リテラルが演算子
		if _, ok := operations[r]; ok {
			if len(stack) < 2 {
				return 0, fmt.Errorf("two numbers on the stack do not exist to be calculatedinput: %v", stack)
			}

			stack, n1 = pop(stack)
			stack, n2 = pop(stack)

			switch r {
			case '+':
				stack = append(stack, n2+n1)
			case '-':
				stack = append(stack, n2-n1)
			case '*':
				stack = append(stack, n2*n1)
			}
			continue
		}

		return 0, fmt.Errorf("invalid input value... input: %s", s)
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("there are still uncomputed values on the stack... stack: %v", stack)
	}

	return stack[0], nil
}

func main() {
	result, err := calc("34+12-*")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
