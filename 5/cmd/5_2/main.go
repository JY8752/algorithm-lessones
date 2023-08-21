/*
N個の正の整数a0,a1,...,aN－1からいくつか選んで
総和を所望の整数Wに一致させることができるかどうかを判定する問題を
O(NW)で解くアルゴリズムを設計してください．
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
	n  int
	w  int
	a  []int
)

func init() {
	isDebug := flag.Bool("d", false, "debug flag")
	flag.Parse()

	if *isDebug {
		f, err := os.Open("./data/5_2/data.txt")
		if err != nil {
			log.Fatalf("not found data.txt... err: %s\n", err.Error())
		}
		defer f.Close()
		sc = bufio.NewScanner(f)
	}

	sc.Split(bufio.ScanWords)

	// 入力値
	n = scanInt(sc)
	w = scanInt(sc)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt(sc)
	}
}

// N = 4
// A = {2, 2, 3, 5}
// W = 7
func main() {
	// dp
	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, w+1)
	}

	dp[0][0] = true

	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ {
			// a[i]を選ばない場合
			if dp[i][j] {
				dp[i+1][j] = true
			}

			// a[i]を選ぶ場合
			if j >= a[i] && dp[i][j-a[i]] {
				dp[i+1][j] = true
			}
		}
	}

	result := dp[n][w]
	fmt.Println(result)
}

func chmax(a *int, b int) {
	if *a < b {
		*a = b
	}
}

func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}
