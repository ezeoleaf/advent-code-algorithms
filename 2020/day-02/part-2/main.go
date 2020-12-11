package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	fmt.Printf("Result is %v \n", run())

	log.Printf("Code took %s", time.Since(start))
}

func run() int {
	valid := 0

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if calculate(scanner.Text()) {
			valid++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return valid
}

func calculate(val string) bool {
	sp := strings.Split(val, ":")

	v := sp[1]

	p1, p2, char := getPolicyData(sp[0])

	c1 := v[p1 : p1+1]
	c2 := v[p2 : p2+1]

	if c1 == char && c2 == char {
		return false
	} else if c1 == char {
		return true
	} else if c2 == char {
		return true
	}

	return false
}

func getPolicyData(s string) (int, int, string) {
	sp := strings.Split(s, " ")

	vs := strings.Split(sp[0], "-")

	p1, _ := strconv.Atoi(vs[0])
	p2, _ := strconv.Atoi(vs[1])

	c := sp[1]

	return p1, p2, c
}
