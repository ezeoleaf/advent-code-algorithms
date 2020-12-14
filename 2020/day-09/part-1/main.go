package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type numbers []int

var preamble = 25

func main() {
	start := time.Now()

	fmt.Printf("Result is %v \n", run())

	log.Printf("Code took %s", time.Since(start))
}

func run() int {

	nums := numbers{}

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, val)
	}

	return getBreak(nums)
}

func getBreak(nums numbers) int {
	pos := preamble
	for {
		found := false
		total := nums[pos]

		rangeVals := nums[pos-preamble : pos]
		mapVals := make(map[int]interface{})
		for _, v := range rangeVals {
			mapVals[v] = nil
		}

		for k := range mapVals {
			if k > total {
				continue
			}

			result := total - k

			if result == k {
				continue
			}

			if _, ok := mapVals[result]; ok {
				found = true
				break
			}
		}

		if !found {
			return total
		}

		if pos+1 >= len(nums) {
			break
		}

		pos++
	}

	return 0
}
