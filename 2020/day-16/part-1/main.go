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

type rules map[string][]validValues
type validValues struct {
	min int
	max int
}
type myTicket []string
type tickets [][]string

var validKeys = map[string]interface{}{
	"departure location": nil,
	"departure station":  nil,
	"departure platform": nil,
	"departure track":    nil,
	"departure date":     nil,
	"departure time":     nil,
	"arrival location":   nil,
	"arrival station":    nil,
	"arrival platform":   nil,
	"arrival track":      nil,
	"class":              nil,
	"duration":           nil,
	"price":              nil,
	"route":              nil,
	"row":                nil,
	"seat":               nil,
	"train":              nil,
	"type":               nil,
	"wagon":              nil,
	"zone":               nil,
}

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

	loadingMyTicket := false
	loadingNearbyTickets := false

	// myTkt := myTicket{}
	nearbyTkts := tickets{}
	rls := rules{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		if strings.Contains(scanner.Text(), "your ticket") {
			loadingMyTicket = true
			loadingNearbyTickets = false
			continue
		}

		if strings.Contains(scanner.Text(), "nearby tickets") {
			loadingNearbyTickets = true
			loadingMyTicket = false
			continue
		}

		if !loadingMyTicket && !loadingNearbyTickets {
			key, vals := getKeyAndVals(scanner.Text())
			rls[key] = vals
		} else if loadingMyTicket {
			_ = strings.Split(scanner.Text(), ",")
		} else if loadingNearbyTickets {
			nearbyTkts = append(nearbyTkts, strings.Split(scanner.Text(), ","))
		}
	}

	// fmt.Println(rls)
	// fmt.Println(myTkt)
	// fmt.Println(nearbyTkts)

	return getErrorRate(rls, nearbyTkts)
}

func getErrorRate(r rules, tks tickets) int {
	var result int
	founds := make(map[int]interface{})

	for _, t := range tks {
		for _, v := range t {
			found := false
			ticketVal, _ := strconv.Atoi(v)

			for _, rule := range r {
				for _, valRule := range rule {
					if ticketVal >= valRule.min && ticketVal <= valRule.max {
						found = true
						break
					}
				}
				if found {
					break
				}
			}

			if !found {
				founds[ticketVal] = nil
				result += ticketVal
			}
		}
	}

	return result
}

func getKeyAndVals(s string) (string, []validValues) {
	data := strings.Split(s, ":")
	k := data[0]

	values := strings.Split(data[1], " or ")

	vv := []validValues{}

	for _, v := range values {
		vrange := strings.Split(strings.TrimSpace(v), "-")
		min, _ := strconv.Atoi(vrange[0])
		max, _ := strconv.Atoi(vrange[1])

		validValue := validValues{min: min, max: max}

		vv = append(vv, validValue)
	}

	return k, vv
}
