/*
(()(())())) のような`(`と`)`からなる長さ2Nの文字列が与えられる。
かっこがちゃんと整合しているか判別する。
また、何文字目と何文字目のかっこが対応しているかをN組求める
計算量はO(N)
*/
package main

import "fmt"

type Stack[T any] []T

func (s Stack[T]) pop() (Stack[T], T, error) {
	if len(s) < 1 {
		var zero T
		return nil, zero, fmt.Errorf("stack is empty")
	}
	return s[:len(s)-1], s[len(s)-1], nil
}

func main() {
	str := "(())"

	type element struct {
		Index int
		Value rune
	}

	var stack Stack[element]
	var result [][]int
	for i, s := range str {
		// stackがカラなので要素を追加して次
		if len(stack) < 1 {
			stack = append(stack, element{Index: i, Value: s})
			continue
		}

		// stack内のかっこと対ではないので追加して次
		if stack[len(stack)-1].Value == s {
			stack = append(stack, element{Index: i, Value: s})
			continue
		}

		// stack内のかっこと対になっているはずなので
		var (
			elm element
			err error
		)

		// stackから取り出してペア情報をresultに入れる
		if stack, elm, err = stack.pop(); err == nil {
			result = append(result, []int{elm.Index, i})
		}
	}

	// stackが空なら整合してる
	fmt.Printf("整合しているかどうか: %v\n", len(stack) == 0)
	// かっこが対応してるN組
	fmt.Printf("%v\n", result)
}
