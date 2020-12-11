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

var plane = map[string]string{}
var availableIDS = map[int64]interface{}{}

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

	for scanner.Scan() {
		setSeats(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	seatID := findSeat()

	return seatID
}

func findSeat() int64 {
	for row := 0; row < 128; row++ {
		for col := 0; col < 8; col++ {
			key := fmt.Sprintf("%v-%v", row, col)
			if _, ok := plane[key]; !ok {
				id := int64(row*8 + col)

				_, idPrevOk := availableIDS[id-1]
				_, idPostOk := availableIDS[id+1]

				if idPrevOk && idPostOk {
					return id
				}
			}
		}
	}

	return 0
}

func setSeats(seatCode string) {

	for k, v := range mapConvert {
		seatCode = strings.ReplaceAll(seatCode, k, v)
	}

	row, _ := strconv.ParseInt(seatCode[0:7], 2, 64)
	col, _ := strconv.ParseInt(seatCode[7:], 2, 64)

	key := fmt.Sprintf("%v-%v", row, col)
	plane[key] = seatCode
	availableIDS[row*8+col] = nil
}
