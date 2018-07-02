package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

type problem struct {
	question string
	answer   string
}

func main() {
	filename := flag.String("csv", "problems.csv", "csv file with questions and answers")
	flag.Parse()

	file, err := os.Open(*filename)
	assert(err)
	reader := csv.NewReader(bufio.NewReader(file))
	correct := 0
	lines := 0
	for line, err := reader.Read(); err != io.EOF; line, err = reader.Read() {
		q := problem{line[0], strings.TrimSpace(line[1])}
		if evalProblem(q) {
			correct++
		}
		lines++
	}
	fmt.Printf("your score is %d out of %d\n", correct, lines)
}

func evalProblem(p problem) bool {
	fmt.Printf("the quiestion is : %s?\n", p.question)
	var ans string
	fmt.Scan(&ans)
	if ans == p.answer {
		return true
	}
	return false
}
