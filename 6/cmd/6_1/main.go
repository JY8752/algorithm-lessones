/*
どの２要素も互いに相異なるN要素からなる整数列a0,a1,...,aN－1が与えられます．
i＝0,1,...,N－1に対して，aiが全体の中で何番目に小さい値であるかを
O(NlogN)で求めるアルゴリズムを設計してください．
たとえばa＝12,43,7,15,9のとき，答えは(2,4,0,3,1)となります．
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc       *bufio.Scanner = bufio.NewScanner(os.Stdin)
	dataPath                = "./data/6_1/data.txt"
	close    func() error
)

func init() {
	isDebug := flag.Bool("d", false, "debug flag")
	flag.Parse()

	if *isDebug {
		debug(dataPath)
	}

	sc.Split(bufio.ScanWords)
}

func debug(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	sc = bufio.NewScanner(f)
	close = func() error { return f.Close() }
}

func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	if close != nil {
		defer close()
	}

	// 入力値
	var (
		n int
		a []int
	)

	n = scanInt(sc)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt(sc)
	}

	b := make([]int, n)
	copy(b, a)
	sort.Ints(b)

	// 二分探索がlogNでN個のループなので計算量はNlogN
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = sort.SearchInts(b, a[i])
	}

	fmt.Println(result)
}
