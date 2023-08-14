package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 入力1 正の整数N
// 入力2 正の整数W
// 入力3 正の整数列a
//
// {a0, a1, ..., aN-1} 整数列a
// 整数列aの中の部分和でWと等しくなるものがあるかどうか

var s = bufio.NewScanner(os.Stdin)

func init() {
	s.Split(bufio.ScanWords)
}

func scanInt() int {
	s.Scan()
	i, err := strconv.Atoi(s.Text())
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func main() {
	// 入力値
	n := scanInt()
	w := scanInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}

	var exist bool
	for bit := 0; bit < (1 << n); bit++ {
		var sum int
		var arr []int
		for i := 0; i < n; i++ {
			if bit&(1<<i) == (1 << i) {
				sum += a[i]
				arr = append(arr, a[i])
			}
		}
		if sum == w {
			exist = true
			fmt.Println(arr)
			break
		}
	}

	if exist {
		fmt.Println("Yes")
		os.Exit(0)
	} else {
		fmt.Println("No")
		os.Exit(1)
	}
}
