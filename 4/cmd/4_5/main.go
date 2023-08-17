package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	initFunc := func() func(n, cur, use int) int {
		var counter int
		var recursiveFunc func(n, cur, use int) int

		recursiveFunc = func(n, cur, use int) int {
			// ベースコード
			if cur > n {
				return counter
			}

			// 0b111ということは7, 5, 3を全て使っているということ
			if use == 0b111 {
				// fmt.Println(cur)
				counter++
			}

			// 7を付け加える
			recursiveFunc(n, 10*cur+7, use|0b001)

			// 5を付け加える
			recursiveFunc(n, 10*cur+5, use|0b010)

			// 3を付け加える
			recursiveFunc(n, 10*cur+3, use|0b100)

			return counter
		}
		return recursiveFunc
	}

	recursiveFunc := initFunc()
	result := recursiveFunc(n, 0, 0)

	fmt.Println(result)
}
