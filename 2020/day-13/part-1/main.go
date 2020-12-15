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

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	time, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	buses := formatBuses(scanner.Text())

	return calcNextBus(time, buses)
}

func calcNextBus(st int, buses map[int]interface{}) int {
	et := st
	for {
		for busID := range buses {
			if et%busID == 0 {
				diff := et - st
				return busID * diff
			}
		}
		et++
	}
}

func formatBuses(buses string) map[int]interface{} {
	sB := strings.Split(buses, ",")
	mapBuses := make(map[int]interface{})
	for _, v := range sB {
		val, e := strconv.Atoi(v)
		if e != nil {
			continue
		}
		mapBuses[val] = nil
	}

	return mapBuses
}
