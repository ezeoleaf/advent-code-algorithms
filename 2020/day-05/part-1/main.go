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

var mapConvert = map[string]string{
	"F": "0",
	"B": "1",
	"R": "1",
	"L": "0",
}

func main() {
	start := time.Now()

	fmt.Printf("Result is %v \n", run())

	log.Printf("Code took %s", time.Since(start))
}

func run() int64 {
	var maxSeatID int64

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		seatID := getSeatID(scanner.Text())
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return maxSeatID
}

func getSeatID(seatCode string) int64 {

	for k, v := range mapConvert {
		seatCode = strings.ReplaceAll(seatCode, k, v)
	}

	row, _ := strconv.ParseInt(seatCode[0:7], 2, 64)
	col, _ := strconv.ParseInt(seatCode[7:], 2, 64)

	return row*8 + col
}
