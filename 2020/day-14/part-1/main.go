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

type memory map[int]string

const maskLength = 36

func main() {
	start := time.Now()

	fmt.Printf("Result is %v \n", run())

	log.Printf("Code took %s", time.Since(start))
}

func run() int64 {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	mem := make(memory)
	var mask string

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "mask") {
			mask = scanner.Text()
			continue
		}

		mem.setInstruction(scanner.Text(), mask)
	}

	return mem.calcResult()
}

func (m memory) setInstruction(value, mask string) {
	regN, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	sv := strings.Split(value, " = ")

	memoryAddress, _ := strconv.Atoi(regN.ReplaceAllString(sv[0], ""))

	n, _ := strconv.ParseInt(sv[1], 10, 64)
	memVal := fmt.Sprintf("%036s\n", strconv.FormatInt(n, 2))
	m[memoryAddress] = applyMask(memVal, mask)
}

func applyMask(val, mask string) string {
	smask := strings.Split(mask, " = ")
	ns := ""
	for i, s := range smask[1] {

		if string(s) == "X" || string(s) == val[i:i+1] {
			ns += val[i : i+1]
		} else {
			if val[i:i+1] == "1" {
				ns += "0"
			} else {
				ns += "1"
			}
		}
	}
	return ns
}

func (m memory) calcResult() int64 {
	var amount int64
	for _, v := range m {
		val, _ := strconv.ParseInt(v, 2, 64)
		amount += val
	}

	return amount
}
