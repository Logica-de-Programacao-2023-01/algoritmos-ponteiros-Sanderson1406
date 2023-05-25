package main

import "fmt"

func primo(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primo2(slice *[]int, n int) {
	soma := 0
	num := 2
	for soma < n {
		if primo(num) {
			*slice = append(*slice, num)
			soma++
		}
		num++
	}
}

func main() {
	var primes []int
	size := 10
	primo2(&primes, size)
	fmt.Println(primes)
}
