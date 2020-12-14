package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type instructions []string

type position struct {
	y       int
	x       int
	heading string
	compass int
}

const (
	east  = "E"
	west  = "W"
	north = "N"
	south = "S"
)

var compassValues = map[int]string{
	90:  east,
	180: south,
	270: west,
	360: north,
}

func main() {
	start := time.Now()

	fmt.Printf("Result is %v \n", run())

	log.Printf("Code took %s", time.Since(start))
}

func run() int {

	ins := instructions{}

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		ins = append(ins, scanner.Text())
	}

	px := position{heading: east, compass: 90}

	px.navigate(ins)

	return px.calculateManhattanDistance()
}

func (p *position) moveShip(direction string, amount int) {
	switch direction {
	case east:
		p.x += amount
	case west:
		p.x -= amount
	case north:
		p.y += amount
	case south:
		p.y -= amount
	}
}

func (p *position) navigate(ins instructions) {
	regS, err := regexp.Compile("[^a-zA-Z]+")
	regN, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range ins {
		action := regS.ReplaceAllString(v, "")
		amount, _ := strconv.Atoi(regN.ReplaceAllString(v, ""))
		switch action {
		case "F":
			p.moveShip(p.heading, amount)
		case "R":
			p.compass += amount
			if p.compass > 360 {
				p.compass -= 360
			}
			p.heading = compassValues[p.compass]
		case "L":
			p.compass -= amount
			if p.compass <= 0 {
				p.compass += 360
			}
			p.heading = compassValues[p.compass]
		default:
			p.moveShip(action, amount)
		}
	}
}

func (p *position) calculateManhattanDistance() int {
	return Abs(p.x) + Abs(p.y)
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
