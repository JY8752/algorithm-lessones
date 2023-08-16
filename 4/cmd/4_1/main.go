/*
4.1 4.2
T0 = 0
T1 = 0
T2 = 1
TN = TN-1 + TN-2 + TN - 3

フィボナッチ数列より一つ項が多いトリボナッチ数列を考える
計算量はO(N)
*/
package main

import (
	"fmt"
)

var memo []int

func toribo(n int) int {
	fmt.Printf("toribo(%d)を呼びました。\n", n)

	// ベースケース
	switch n {
	case 0, 1:
		return 0
	case 2:
		return 1
	}

	// memoをチェック
	if memo[n] != 0 {
		return memo[n]
	}

	result := toribo(n-1) + toribo(n-2) + toribo(n-3)
	fmt.Printf("%d項目 = %d\n", n, result)
	memo[n] = result
	return result
}

func main() {
	n := 5
	memo = make([]int, n+1)
	fmt.Println(toribo(5))
}
