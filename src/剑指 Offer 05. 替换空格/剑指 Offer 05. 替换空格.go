package main

import "fmt"

func replaceSpace(s string) string {
	tmpStr := ""
	for _, y := range s {
		if string(y) == " " {
			tmpStr += "%20"
		} else {
			tmpStr += string(y)
		}
	}

	return tmpStr
}

func main() {
	fmt.Printf(" == %s", replaceSpace("ni hao 111 hhhh"))
}
