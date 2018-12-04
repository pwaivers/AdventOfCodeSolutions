package main

import (
	"bufio"
    "fmt"
    "log"
	"os"
	"strconv"
	"regexp"
)

func main() {
	fmt.Println(foo("test_input.txt") == 4);

	fmt.Println(foo("input.txt"))
}

func foo(filename string) int {
	cuts := read_file(filename)

	// initialize array
	var sheet [1000][1000]int

	for _, cut := range cuts {
		for i := cut.X; i < cut.X+cut.Width; i++ {
			for j := cut.Y; j < cut.Y + cut.Height; j++ {
				sheet[i][j]++
			}
		}
	}
	
	total := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if sheet[i][j] > 1 {
				total++
			}
		}
	}

	return total
}

func read_file(filename string) []Cut {
	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	
	cuts := make([]Cut, 0)
	r := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)  // Ex: #1 @ 1,3: 4x4
    for scanner.Scan() {
		line := scanner.Text()
		groups := r.FindStringSubmatch(line)
		id, _ := strconv.Atoi(groups[1])
		x, _ := strconv.Atoi(groups[2])
		y, _ := strconv.Atoi(groups[3])
		width, _ := strconv.Atoi(groups[4])
		height, _ := strconv.Atoi(groups[5])
		cut := Cut{id, x, y, width, height}

		cuts = append(cuts, cut)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	return cuts;
}

type Cut struct {
	Id int
	X int
	Y int
	Width int
	Height int
}
