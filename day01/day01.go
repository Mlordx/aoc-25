package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	direction string
	steps     int
}

func readInput() []Operation {
	var ops []Operation
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

		var temp Operation
		temp.direction = string(line[0])
		num, _ := strconv.Atoi(strings.TrimSpace(line[1:]))
		temp.steps = num

		ops = append(ops, temp)
	}

	return ops
}

func answer1(input []Operation) {
	password := 0
	position := 50

	for _, value := range input {
		var dir int
		if value.direction == "R" {
			dir = 1
		} else {
			dir = -1
		}
		steps := value.steps

		position += dir * steps
		position = (position%100 + 100) % 100
		if position == 0 {
			password += 1
		}
	}

	fmt.Println(password)
}

func answer2(input []Operation) {
	password := 0
	position := 50

	for _, value := range input {
		var dir int
		if value.direction == "R" {
			dir = 1
		} else {
			dir = -1
		}
		steps := value.steps
		clicks := steps / 100
		steps = steps % 100
		if position != 0 {
			if position+dir*steps < 0 || position+dir*steps > 100 {
				clicks += 1
			}
		}
		position += dir * steps
		position = (position%100 + 100) % 100
		if position == 0 {
			password += 1
		}

		password += clicks
	}

	fmt.Println(password)
}

func main() {
	input := readInput()
	answer1(input)
	answer2(input)
}
