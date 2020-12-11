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
		vals[v] = nil
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return calculate(vals)
}

func calculate(vals map[int64]interface{}) int64 {
	for i := range vals {
		for j := range vals {
			if i != j {
				e := 2020 - i - j
				if _, ok := vals[e]; ok {
					return e * i * j
				}
			}
		}
	}

	return 0
}
