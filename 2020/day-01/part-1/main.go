package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	fmt.Printf("Result is %v \n", run())

	log.Printf("Code took %s", time.Since(start))
}

func run() int64 {
	vals := make(map[int64]interface{})

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		v, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		expected := 2020 - v
		if _, ok := vals[expected]; ok {
			return v * expected
		}

		vals[v] = nil
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return 0
}
