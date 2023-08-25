/*
連結リストを用いてスタックとキューを実現してください
*/
package main

import (
	"container/list"
	"fmt"
)

func pop[T any](l *list.List) T {
	e := l.Back().Value.(T)
	l.Remove(l.Back())
	return e
}

func deque[T any](l *list.List) T {
	e := l.Front().Value.(T)
	l.Remove(l.Front())
	return e
}

func main() {
	names := [6]string{
		"yamamoto",
		"watanabe",
		"ito",
		"takahashi",
		"suzuki",
		"sato",
	}

	// スタック
	l := list.New()

	for _, name := range names {
		l.PushBack(name)
	}

	n1 := pop[string](l)
	println(n1)
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	// キュー
	l2 := list.New()

	for _, name := range names {
		l2.PushBack(name)
	}

	l2.PushBack("yamanaka")
	n2 := deque[string](l2)

	println(n2)
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
