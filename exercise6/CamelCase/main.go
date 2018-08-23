package main

import "fmt"

func main() {
	var input string
	fmt.Scanf("%s", &input)
	ans := 1
	for _, c := range input {
		if c >= 'A' && c <= 'Z' {
			ans++
		}
	}
	fmt.Printf("%d\n", ans)
}
