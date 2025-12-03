package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func readInput() []string {
	var input []string
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

		input = append(input, strings.TrimSpace(line))
	}

	return input
}

func getNextDigit(bank string, start int, remaining int) (float64, int) {
	for l := 9; l > 0; l-- {
		for i := start; i <= len(bank)-remaining; i++ {
			digit := int(bank[i] - '0')
			if digit == l {
				return float64(l), i + 1
			}
		}
	}
	return -1, -1
}

func answer(banks []string, r int) {
	answer := float64(0)
	for _, bank := range banks {
		var num float64
		var start int
		start = 0
		remaining := r
		for remaining > 0 {
			num, start = getNextDigit(bank, start, remaining)

			answer += num * math.Pow(10, float64(remaining-1))
			remaining--
		}
	}
	fmt.Printf("%.0f\n", answer)
}

func main() {
	input := readInput()
	answer(input, 2)
	answer(input, 12)
}
