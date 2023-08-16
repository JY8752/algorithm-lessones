package main

import "fmt"

func main() {
	var (
		s   string
		n   int
		sum int64
	)

	fmt.Scan(&s)
	n = len(s)

	for bit := 0; bit < (1 << (n - 1)); bit++ {
		var tmp int64
		for i := 0; i < n-1; i++ {
			tmp *= 10
			tmp += int64(s[i] - '0')

			if bit&(1<<i) != 0 {
				sum += tmp
				tmp = 0
			}
		}

		// 最後の部分を足す
		tmp *= 10
		tmp += int64(s[n-1] - '0')
		sum += tmp
	}

	fmt.Println(sum)
}
