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

	operations := []string{}

	for scanner.Scan() {
		operations = append(operations, scanner.Text())
	}

	return calcTotal(operations)
}

func calcTotal(ops []string) int {
	total := 0

	for _, o := range ops {
		total += getTotalOp(o)
	}

	return total
}

func getTotalOp(o string) int {
	for strings.Contains(o, "(") {
		p := strings.LastIndex(o, "(")
		o = solveParenthesis(o, p)
	}

	so := strings.Split(o, " ")
	t := 0

	sign := "+"

	for _, val := range so {
		if val == "+" || val == "*" {
			sign = val
			continue
		}

		a, _ := strconv.Atoi(val)

		switch sign {
		case "+":
			t += a
		case "*":
			t *= a
		}

	}

	return t
}

func solveParenthesis(o string, p int) string {
	toSolve := ""
	var closeP int

	for i := p + 1; i < len(o); i++ {
		switch string(o[i]) {
		case ")":
			closeP = i
		default:
			toSolve += string(o[i])
		}

		if closeP != 0 {
			break
		}
	}

	toSolve = o[p+1 : closeP]

	total := getTotalOp(toSolve)

	toSolve = fmt.Sprintf("(%s)", toSolve)

	o = strings.ReplaceAll(o, toSolve, strconv.Itoa(total))

	return o
}
