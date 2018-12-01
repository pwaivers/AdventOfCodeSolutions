package main

import (
	"bufio"
    "fmt"
    "log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(foo("test1.txt") == 3);
	fmt.Println(foo("test2.txt") == 0);
	fmt.Println(foo("test3.txt") == -6);

	fmt.Println(foo("input.txt"))
}

func foo(filename string) int {
	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	
	total := 0
    for scanner.Scan() {
        line := scanner.Text()
		val, _ := strconv.Atoi(line)
		total += val
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	return total;
}
