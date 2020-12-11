package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	k, v := sp[0], sp[1]

	min, max, char := getPolicyData(k)
	search := regexp.MustCompile(char)

	amountOfChars := len(search.FindAllStringIndex(v, -1))

	return (amountOfChars >= min && amountOfChars <= max)
}

func getPolicyData(s string) (int, int, string) {
	var min, max int
	var c string
	sp := strings.Split(s, " ")

	vs := strings.Split(sp[0], "-")

	min, _ = strconv.Atoi(vs[0])
	max, _ = strconv.Atoi(vs[1])

	c = sp[1]

	return min, max, c
}
