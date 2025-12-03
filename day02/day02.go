package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func readInput() []Range {
	var ranges []Range
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

		aux := strings.Split(line, ",")

		for _, value := range aux {
			r := strings.Split(value, "-")

			s, _ := strconv.Atoi(strings.TrimSpace(r[0]))
			e, _ := strconv.Atoi(strings.TrimSpace(r[1]))
			temp := Range{start: s, end: e}

			ranges = append(ranges, temp)
		}
	}

	return ranges
}

func isInvalid1(n int) bool {
	string_num := strconv.Itoa(n)
	if len(string_num)%2 == 1 {
		return false
	}
	l := len(string_num) / 2
	return string_num[:l] == string_num[l:]
}

func answer1(ranges []Range) {
	answer := 0

	for _, value := range ranges {
		for i := value.start; i <= value.end; i++ {
			if isInvalid1(i) {
				answer += i
			}
		}
	}

	fmt.Println(answer)
}

func isInvalid2(n int) bool {
	if isInvalid1(n) {
		return true
	}

	string_num := strconv.Itoa(n)
	for l := len(string_num) / 2; l > 0; l-- {
		seq := string_num[:l]
		var count float64 = 1

		start := l
		end := start + l

		expected := float64(len(string_num)) / float64(len(seq))

		for end <= len(string_num) {
			if seq == string_num[start:end] {
				count += 1
			}
			start = end
			end = start + l
		}

		if count == expected {
			return true
		}
	}
	return false
}

func answer2(ranges []Range) {
	answer := 0

	for _, value := range ranges {
		for i := value.start; i <= value.end; i++ {
			if isInvalid2(i) {
				answer += i
			}
		}
	}

	fmt.Println(answer)
}

func main() {
	input := readInput()
	answer1(input)
	answer2(input)
}
