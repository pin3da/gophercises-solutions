package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func assert(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	filename := "problems.csv"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	file, err := os.Open(filename)
	assert(err)
	reader := csv.NewReader(bufio.NewReader(file))
	total := 0
	lines := 0
	for line, err := reader.Read(); err != io.EOF; line, err = reader.Read() {
		fmt.Printf("the quiestion is : %s?\n", line[0])
		var ans string
		fmt.Scan(&ans)
		if ans == line[1] {
			fmt.Printf("correct :D\n")
			total++
		} else {
			fmt.Printf("incorrect ):\n")
		}
		lines++
	}
	fmt.Printf("your score is %v out of %v\n", total, lines)
}
