package main

import (
	"bufio"
    "fmt"
    "log"
	"os"
)

func main() {
	fmt.Println(foo("test_input.txt") == 12);

	fmt.Println(foo("input.txt"))
}

func foo(filename string) int {
	lines := read_file(filename)
	twos, threes := 0, 0
	for i := range lines {
		a, b := get_counts(lines[i])
		twos += a
		threes += b
	}
	return twos*threes
}

func get_counts(line string) (int, int) {
	counts := make(map[byte]int, 0)
	for i := 0; i < len(line); i++ {
        _, ok := counts[line[i]]
		if (ok) {
			counts[line[i]] += 1
		} else {
			counts[line[i]] = 1
		}
	}
	
	twos, threes := 0, 0
	for _, v := range counts { 
		if v == 2 {
			twos = 1
		}
		if v == 3 {
			threes = 1
		}
	}

	return twos, threes
}

func read_file(filename string) []string {
	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	
	lines := make([]string, 0)
    for scanner.Scan() {
        line := scanner.Text()
		lines = append(lines, line)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	return lines;
}
