package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	numbers []int
	op      string
}

func readInput1() []Problem {
	var problems []Problem

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		line = strings.Join(strings.Fields(line), " ")
		if err != nil {
			break
		}
		if len(line) != 0 {
			values := strings.Split(line, " ")

			if len(problems) == 0 {
				for range values {
					problems = append(problems, Problem{})
				}
			}

			if values[0] == "*" || values[0] == "+" {
				for i, value := range values {
					problems[i].op = value
				}
			} else {
				for i, number_string := range values {
					number, _ := strconv.Atoi(number_string)
					problems[i].numbers = append(problems[i].numbers, number)
				}
			}
		}
	}

	return problems
}

func readInput2() []string {
	var lines []string

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')

		if err != nil {
			break
		}
		if len(line) != 0 {
			lines = append(lines, line)
		}
	}

	return lines
}

func answer1(problems []Problem) {
	answer := 0

	for _, problem := range problems {
		op := problem.op

		if op == "+" {
			temp := 0
			for _, value := range problem.numbers {
				temp += value
			}
			answer += temp
		} else {
			temp := 1

			for _, value := range problem.numbers {
				temp *= value
			}
			answer += temp
		}
	}

	fmt.Println(answer)
}

func compute(numbers [4][4]int, op byte) int {
	var ns []int

	for j := range 4 {
		i := 3
		base := 3
		for i = 3; i >= 0; i-- {
			if numbers[i][j] == -1 {
				base--
			} else {
				break
			}
		}
		num := 0
		for i = 0; i < 4; i++ {
			if numbers[i][j] != -1 {
				num += int(math.Pow(10, float64(base)) * float64(numbers[i][j]))
			}
			base--
		}
		if num == 0 {
			continue
		}
		ns = append(ns, num)
	}

	if op == '+' {
		temp := 0

		for _, num := range ns {
			temp += num
		}

		return temp
	} else {
		temp := 1

		for _, num := range ns {
			temp *= num
		}
		return temp
	}
}

func blankColumn(lines []string, i int) bool {
	for _, line := range lines {
		if line[i] != ' ' && line[i] != '\n' {
			return false
		}
	}

	return true
}

func answer2(lines []string) {
	answer := 0
	length := len(lines[0]) - 1 //skipping newline
	numbers := [4][4]int{
		{-1, -1, -1, -1},
		{-1, -1, -1, -1},
		{-1, -1, -1, -1},
		{-1, -1, -1, -1},
	}
	pos := 3

	for i := length; i >= 0; i-- {
		if blankColumn(lines, i) {
			continue
		}

		for a := range 4 {
			if lines[a][i] == ' ' {
				continue
			}
			dig, _ := strconv.Atoi(string(lines[a][i]))
			numbers[a][pos] = dig
		}
		if lines[4][i] == '+' || lines[4][i] == '*' {
			answer += compute(numbers, lines[4][i])
			numbers = [4][4]int{
				{-1, -1, -1, -1},
				{-1, -1, -1, -1},
				{-1, -1, -1, -1},
				{-1, -1, -1, -1},
			}
			pos = 3
			continue
		}

		pos--
	}

	fmt.Println(answer)
}

func main() {
	problems1 := readInput1()
	answer1(problems1)
	problems2 := readInput2()
	answer2(problems2)
}
