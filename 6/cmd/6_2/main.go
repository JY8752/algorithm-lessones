/*
*
N要素からなる３つの整数列a0,...,aN－1，b0,...,bN－1，c0,...,cN－1が与えられます．
ai＜bj＜ckを満たすようなi,j,kの組が何個あるかをO(NlogN)で求めるアルゴリズムを設計してください．
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
	dataPath                = "./data/6_2/data.txt"
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
		b []int
		c []int
	)

	n = scanInt(sc)
	a = make([]int, n)
	b = make([]int, n)
	c = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt(sc)
	}
	for i := 0; i < n; i++ {
		b[i] = scanInt(sc)
	}
	for i := 0; i < n; i++ {
		c[i] = scanInt(sc)
	}

	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)

	// bを固定して考え
	var result int
	for i := 0; i < n; i++ {
		// bi未満のaの要素数をakとする
		ak := sort.Search(len(a), func(ii int) bool { return a[ii] >= b[i] })

		// biより上の要素数をckとする
		ck := n - sort.Search(len(c), func(ii int) bool { return c[ii] > b[i] })

		// ak * ckがbiの時の考えられる通り
		result += ak * ck
	}

	fmt.Println(result)
}
