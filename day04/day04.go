package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type p struct {
	x int
	y int
}

func readInput() (map[p]string, int, int) {
	grid := make(map[p]string)
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil, -1, -1
	}

	defer f.Close()

	r := bufio.NewReader(f)
	i := 0
	width := 0
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			break
		}
		width = len(line)
		for j := 0; j < len(line); j++ {
			grid[p{i, j}] = string(line[j])
		}
		i++
	}

	return grid, i, width
}

func answer1(grid map[p]string, height int, width int) {
	answer := 0

	for i := range height {
		for j := range width {
			rolls := 0
			if grid[p{i, j}] == "." {
				continue
			}

			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}
					value, ok := grid[p{i + x, j + y}]
					if ok && value == "@" {
						rolls++
					}
				}
			}
			if rolls < 4 {
				answer++
			}
		}
	}

	fmt.Println(answer)
}

func answer2(grid map[p]string, height int, width int) {
	answer := 0

	for {
		var to_be_removed []p
		for i := range height {
			for j := range width {
				rolls := 0
				if grid[p{i, j}] == "." {
					continue
				}

				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if x == 0 && y == 0 {
							continue
						}
						value, ok := grid[p{i + x, j + y}]
						if ok && value == "@" {
							rolls++
						}
					}
				}
				if rolls < 4 {
					answer++
					to_be_removed = append(to_be_removed, p{i, j})
				}
			}
		}

		if len(to_be_removed) == 0 {
			break
		}
		for _, value := range to_be_removed {
			grid[value] = "."
		}
	}

	fmt.Println(answer)
}

func main() {
	grid, height, width := readInput()
	answer1(grid, height, width)
	answer2(grid, height, width)
}
