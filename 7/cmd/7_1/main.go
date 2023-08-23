/*
N個の整数a0, a1,...,aN-1と、N個の整数b0,b1,...bN-1が与えられます。
a0,...,aN-1から何個かとb0,...,bN-1から何個か選んでペアを作ります。
ただし、各ペア(ai,bj)はai<bjを満たさなければならない。
最大で何ペア作れるか？
計算量はO(NlogN)で
*/
package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
)

var (
	sc       *bufio.Scanner = bufio.NewScanner(os.Stdin)
	dataPath                = "./data/7_1/data.txt"
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
