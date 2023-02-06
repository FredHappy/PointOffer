package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	if s == "" || n == 0 {
		return ""
	}
	return s[n:] + s[:n]
}

func main() {
	fmt.Printf("reverseLeftWords = %s", reverseLeftWords("", 0))
}
