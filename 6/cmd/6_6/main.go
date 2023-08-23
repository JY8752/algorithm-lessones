/*
高橋君は野球が得意です。高橋君は、高橋君ボール
1 号という変化球を投げることが出来ます。

このボールは、投げてから
t 秒後のボールの位置を
f(t) とすると、
f(t)=A×t+B×sin(C×t×π) と表すことが出来ます。

あなたは、
t≧0 かつ
f(t)=100 となるタイミングで、ボールを打たなければいけません。この時の
t を求めたいです。
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc       *bufio.Scanner = bufio.NewScanner(os.Stdin)
	dataPath                = "./data/6_6/data.txt"
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
		a int
		b int
		c int
	)

	a = scanInt(sc)
	b = scanInt(sc)
	c = scanInt(sc)

	// f(t) = 100の解を求めるのは難しそう.
	// なので以下のような二分探索を考える
	// α < x < β
	// f(x) = 100
	// となるようなα、βについて考えると二分探索が使えそう
	// 区間(α, β)の中にf(x) < 100からf(x) > 100に変わる瞬間がある
	// γ = (α + β)/2 としたときに
	// f(γ) <= 100 なら (γ, β)
	// f(γ) >= 100 なら　　(α, γ)
	var (
		alpha float64 = 0
		beta  float64 = 200
	)

	calc := func(t float64) float64 {
		return float64(a)*t + float64(b)*math.Sin(float64(c)*math.Pi*t)
	}

	// ループ回数を100にしているのは100回(α, β)の範囲を半分にする2^-100ということで
	// 求められている精度を出すにはこのくらいで十分という判断
	for i := 0; i < 100; i++ {
		gamma := (alpha + beta) / 2
		if calc(gamma) <= 100 {
			alpha = gamma
		} else {
			beta = gamma
		}
	}

	fmt.Println(alpha, beta)
}
