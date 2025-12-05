package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	l int
	r int
}

func readInput() ([]Range, []int) {
	var ranges []Range
	var ingredients []int
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			break
		}
		if len(line) != 0 {
			if strings.ContainsRune(line, '-') {
				temp := strings.Split(line, "-")
				l, _ := strconv.Atoi(temp[0])
				r, _ := strconv.Atoi(temp[1])

				ranges = append(ranges, Range{l, r})
			} else {
				num, _ := strconv.Atoi(line)
				ingredients = append(ingredients, num)
			}
		}
	}

	return ranges, ingredients
}

func answer1(ranges []Range, ingredients []int) {
	answer := 0

	for _, value := range ingredients {
		fresh := false
		for _, ran := range ranges {
			l := ran.l
			r := ran.r

			if l <= value && value <= r {
				fresh = true
				break
			}
		}
		if fresh {
			answer++
			continue
		}
	}

	fmt.Println(answer)
}

func answer2(ranges []Range) {
	answer := 0

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].l < ranges[j].l
	})
	var merged_ranges []Range
	current := ranges[0]

	for i := 1; i < len(ranges); i++ {
		l1 := current.l
		r1 := current.r
		next := ranges[i]
		l2 := next.l
		r2 := next.r

		if l2 > r1 {
			merged_ranges = append(merged_ranges, current)
			current = next
			continue
		}

		if r2 <= r1 {
			continue
		}

		if l2 <= r1 {
			current = Range{l1, r2}
		}
	}
	merged_ranges = append(merged_ranges, current)

	for _, value := range merged_ranges {
		answer += value.r - value.l + 1
	}

	fmt.Println(answer)
}

func main() {
	ranges, ingredients := readInput()
	answer1(ranges, ingredients)
	answer2(ranges)
}
