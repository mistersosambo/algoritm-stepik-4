package main

import "fmt"

func main() {
	var n, z int
	var m []int
	fmt.Scan(&n)
	for j := 0; j < n; j++ {
		fmt.Scan(&z)
		m = append(m, z)
	}
	for d := range m {
		if d%2 == 0 {
			fmt.Print(m[d], " ")
		} else {
			continue
		}
	}
}
