/*
N日間の夏休みがあり，i日目に海で泳ぐ幸福度はai，虫捕りする幸福度はbi，宿題をする幸福度はciで与えられるとします．
それぞれの日について，これらの３つの行動のうちのいずれかを行います．
ただし２日連続で同じ行動はしないものとします．
N日間の幸福度の最大値をO(N)で求めるアルゴリズムを設計してください．
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
	a  [][3]int
)

func init() {
	isDebug := flag.Bool("d", false, "debug flag")
	flag.Parse()

	if *isDebug {
		f, err := os.Open("./data/5_1/data.txt")
		if err != nil {
			log.Fatalf("not found data.txt... err: %s\n", err.Error())
		}
		defer f.Close()
		sc = bufio.NewScanner(f)
	}

	sc.Split(bufio.ScanWords)

	// 入力値
	n = scanInt(sc)
	a = make([][3]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = scanInt(sc)
		}
	}
}

// N = 3
// a = {1, 2, 3}
// b = {3. 2, 1}
// c = {3, 3, 3}
func main() {

	// dpテーブル
	dp := make([][3]int, n+1)

	// i日目のときjを選択
	// i + 1日目ではkを選択するとする
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				// 同じイベントは連続で選択できない
				if j == k {
					continue
				}
				chmax(&dp[i+1][k], dp[i][j]+a[i][k])
			}
		}
	}

	var result int
	for j := 0; j < 3; j++ {
		chmax(&result, dp[n][j])
	}

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
