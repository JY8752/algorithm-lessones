package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var s = bufio.NewScanner(os.Stdin)

func init() {
	s.Split(bufio.ScanWords)
}

func scanInt(s *bufio.Scanner) (int, error) {
	s.Scan()
	i, err := strconv.Atoi(s.Text())
	if err != nil {
		return 0, err
	}
	return i, nil
}

// N = 4
// A = {16, 32, 64, 128}
// ans = 4

// N = 2
// A = {1024, 2048}
// ans = 10

// N = 5
// A = {16, 32, 64, 58}
// ans = 0
func main() {
	// 入力値
	n, err := scanInt(s)
	if err != nil {
		log.Fatal(err)
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		ii, err := scanInt(s)
		if err != nil {
			log.Fatal(err)
		}
		a[i] = ii
	}

	var min int
	isFirst := true
	for _, i := range a {
		c := divideByTwoCount(i)
		if isFirst {
			isFirst = false
			min = c
			continue
		}
		if c < min {
			min = c
		}
	}

	fmt.Println(min)
}

func divideByTwoCount(n int) int {
	count := 0
	for n > 1 {
		if n%2 != 0 {
			// 2で割り切れない
			return count
		}
		n /= 2
		count++
	}
	return count
}
