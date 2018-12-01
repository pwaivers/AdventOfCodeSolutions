package main

import (
	"bufio"
    "fmt"
    "log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(foo("test1.txt") == 0);
	fmt.Println(foo("test2.txt") == 10);
	fmt.Println(foo("test3.txt") == 5);
	fmt.Println(foo("test4.txt")== 14);

	fmt.Println(foo("input.txt"))
}

func foo(filename string) int {
	vals := read_file(filename)

	total := 0

	freqs := make(map[int]bool)
	freqs[total] = true
	for true { 
		for _, v := range vals {
			total += v
			_, ok := freqs[total]
			if (ok) {
				return total
			}

			freqs[total] = true
		}
	}

	return 0
}

func read_file(filename string) []int {
	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	
	vals := make([]int, 0)
    for scanner.Scan() {
        line := scanner.Text()
		val, _ := strconv.Atoi(line)
		vals = append(vals, val)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	return vals;
}
