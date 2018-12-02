package main

import (
	"bufio"
    "fmt"
    "log"
	"os"
)

func main() {
	fmt.Println(foo("test_input.txt") == "fgij");

	fmt.Println(foo("input.txt"))
}

func foo(filename string) string {
	lines := read_file(filename)
	for i := range lines {
		for j := i + 1; j < len(lines); j += 1 {
			equal, result := compare(lines[i], lines[j])
			if equal {
				return result
			}
		}
	}

	return ""
}

func compare(line1 string, line2 string) (bool, string) {
	if len(line1) != len(line2){
		return false, ""
	}

	errors := 0
	word := ""
	for i := range line1 {
		if line1[i] == line2[i] {
			word = word + string(line1[i])
		} else {
			errors += 1
			if errors > 1 {
				return false, ""
			}
		}
	}

	return true, word
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
