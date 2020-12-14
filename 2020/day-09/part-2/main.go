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

	breakValue, ix := getBreak(nums)

	return calculateWeakness(nums, breakValue, ix)
}

func getBreak(nums numbers) (int, int) {
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
			return total, pos
		}

		if pos+1 >= len(nums) {
			break
		}

		pos++
	}

	return 0, 0
}

func calculateWeakness(nums numbers, breakVal, breakIndex int) int {
	sp := -1

	for {
		sumNumbers := []int{}
		total := 0
		found := false
		if sp >= breakIndex {
			break
		}
		np := sp + 1
		for {
			total += nums[np]
			sumNumbers = append(sumNumbers, nums[np])
			np++
			if len(sumNumbers) < 2 {
				continue
			}

			if total > breakVal {
				break
			}

			if total == breakVal {
				found = true
				break
			}
		}

		if found {
			min, max := getMinAndMax(sumNumbers)

			return min + max
		}

		sp++
	}

	return 0
}

func getMinAndMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]

	for _, n := range nums {
		if min == -1 {
			min = n
		}

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	return min, max
}
