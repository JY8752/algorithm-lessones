/*
N個の整数列Aの中に和がWとなる組み合わせが存在するかどうか
再帰関数を使用した解法
計算量はO(2^n)
これをメモ化を使用してO(NW)とする
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	s       = bufio.NewScanner(os.Stdin)
	memo    [][]int
	counter int
)

func init() {
	s.Split(bufio.ScanWords)
}

func scanInt(s *bufio.Scanner) int {
	s.Scan()
	i, err := strconv.Atoi(s.Text())
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func resolve(i, w int, a []int) int {
	if i == 0 {
		if w == 0 {
			return 1
		}
		return 0
	}

	// 既にmemo化された値があれば返す
	if memo[i][w] != -1 {
		return memo[i][w]
	}

	counter++

	// a[i - 1]を選ばない場合
	if resolve(i-1, w, a) == 1 {
		memo[i][w] = 1
		return 1
	}

	// a[i - 1]を選ぶ場合
	if resolve(i-1, w-a[i-1], a) == 1 {
		memo[i][w] = 1
		return 1
	}

	memo[i][w] = 0
	return 0
}

func main() {
	n := scanInt(s)
	w := scanInt(s)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = scanInt(s)
	}

	// メモの初期化
	memo = make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, w+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	if resolve(n, w, a) == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	fmt.Println(counter)
}
