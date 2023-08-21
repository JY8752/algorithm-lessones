/*
N個の正の整数a0,a1,...,aN－1と正の整数Wが与えられます．
N個の整数それぞれについてm0,m1,...,mN－1回まで足してもよいとしたときに，
総和をWに一致させられるかをO(NW)で判定するアルゴリズムを設計してください．
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
	n  int
	w  int
	a  []int
	m  []int
)

func init() {
	isDebug := flag.Bool("d", false, "debug flag")
	flag.Parse()

	if *isDebug {
		f, err := os.Open("./data/5_6/data.txt")
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
	m = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt(sc)
		m[i] = scanInt(sc)
	}
}

func main() {
	// dp
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, w+1)
		for j := 0; j < w+1; j++ {
			dp[i][j] = math.MaxInt16
		}
	}

	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ {
			// i個で既にjが作れる場合
			if dp[i][j] < math.MaxInt16 {
				chmin(&dp[i+1][j], 0)
			}

			// i+1個でm[i]個未満でj-a[i]が作れるなら、jも作れる
			if j >= a[i] && dp[i+1][j-a[i]] < m[i] {
				chmin(&dp[i+1][j], dp[i+1][j-a[i]]+1)
			}
		}
	}

	if dp[n][w] < math.MaxInt16 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	fmt.Println(dp[n][w])
}

func chmax(a *int, b int) {
	if *a < b {
		*a = b
	}
}

func chmin(a *int, b int) {
	if *a > b {
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
