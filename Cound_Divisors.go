package main

import "fmt"

func main() {
	var l, r, k int
	var c int
	c = 0
	fmt.Scan(&l, &r, &k)
	if l <= 1000 && r <= 1000 && k <= 1000 {
		for i := l; i <= r; i++ {
			if i%k == 0 {
				c = c + 1
			}
		}
		fmt.Print(c)
	}
}
