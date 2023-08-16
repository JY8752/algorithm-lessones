// ２つの正の整数K,Nが与えられます．
// 0<=X,Y,Z<=Kを満たす整数(X,Y,Z)の組であってX＋Y＋Z＝Nを満たすものが何通りあるかを求めるO(N^2)のアルゴリズムを設計してください．
// (出典:AtCoderBeginnerContest 051 B SumofThreeIntegers，
package main

import "fmt"

func main() {
	// 入力値
	var k, n int
	fmt.Scan(&k, &n)

	var count int
	for x := 0; x <= min(k, n); x++ {
		for y := 0; y <= min(k, n); y++ {
			z := n - x - y
			if z >= 0 && z <= k {
				count++
			}
		}
	}
	fmt.Println(count)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
