package main

import "fmt"

func main() {
	var len, offset int
	var input string
	fmt.Scanf("%d\n%s\n%d\n", &len, &input, &offset)

	for _, c := range input {
		if c >= 'a' && c <= 'z' {
			fmt.Print(convert(c, offset, 'a'))
		} else if c >= 'A' && c <= 'Z' {
			fmt.Print(convert(c, offset, 'A'))
		} else {
			fmt.Print(string(c))
		}
	}
	fmt.Println()
}

func convert(c rune, offset int, base rune) string {
	a := int(c - base)
	a = (((a + offset) % 26) + 26) % 26
	return string(int(base) + a)
}
