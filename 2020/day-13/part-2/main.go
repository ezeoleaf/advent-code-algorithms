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
	scanner.Scan()
	buses, maxBusID, maxIx := formatBuses(scanner.Text())

	return calcNextBus(buses, maxBusID, maxIx)
}

func calcNextBus(buses map[int]int, maxBusID, maxIx int) int {
	i := 0
	for {
		i++
		t := maxBusID * i

		found := true
		for bt, bID := range buses {
			nt := t + (bt - maxIx)
			if maxIx > bt {
				nt = t - (maxIx - bt)
			} else if bt == maxIx {
				nt = t
			}
			// fmt.Println(t)
			// fmt.Println(bt, maxIx)
			// fmt.Println(nt)
			if nt%bID != 0 {
				found = false
				break
			}
		}
		if found {
			return t - maxIx
		}
	}
}

func formatBuses(buses string) (map[int]int, int, int) {
	sB := strings.Split(buses, ",")
	maxBusID, maxIx := 0, 0
	mapBuses := make(map[int]int)
	for i, v := range sB {
		val, e := strconv.Atoi(v)
		if e != nil {
			continue
		}
		if val > maxBusID {
			maxBusID = val
			maxIx = i
		}
		mapBuses[i] = val
	}

	return mapBuses, maxBusID, maxIx
}
