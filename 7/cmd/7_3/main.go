/*
AtCoder王国の王立問題工房でABC管理官の座に就いたキザハシ君は、浮かれるあまり仕事を引き受けすぎてしまいました。

現在の時刻は
0 です。キザハシ君は
1 から
N までの番号が振られた
N 件の仕事を持っています。

キザハシ君が仕事
i を終えるには
Ai単位時間かかります。
また、仕事 i の〆切は時刻Bi
であり、これまでに仕事を終わらせる必要があります。
時刻Biちょうどに仕事iを終わらせてもかまいません。

キザハシ君は
2 件以上の仕事を同時にすることはできませんが、ある仕事を終わらせた直後に別の仕事を始めることはできます。

キザハシ君はすべての仕事を〆切までに終わらせることができるでしょうか。可能ならば Yes、不可能ならば No を出力してください。
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
	dataPath                = "./data/7_3/data.txt"
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

	var (
		n int
		v [][2]int
	)

	n = scanInt(sc)
	v = make([][2]int, n)
	for i := 0; i < n; i++ {
		v[i][0] = scanInt(sc)
		v[i][1] = scanInt(sc)
	}

	// 一旦締切時間でソート
	sort.Slice(v, func(i, j int) bool { return v[i][1] < v[j][1] })

	var count int
	result := true
	for i := 0; i < n; i++ {
		count += v[i][0]
		if count > v[i][1] {
			result = false
			break
		}
	}

	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
