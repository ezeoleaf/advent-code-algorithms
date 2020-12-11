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

type Passport struct {
	CountryID      int
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	PassportID     string
	HairColor      string
	EyeColor       string
	Height         string
}

const (
	BirthYearType      = "byr"
	IssueYearType      = "iyr"
	ExpirationYearType = "eyr"
	HeighType          = "hgt"
	HairColorType      = "hcl"
	EyeColorType       = "ecl"
	PassportIDType     = "pid"
	CountryIDType      = "cid"
)

var passports []Passport

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
	var passport Passport

	for scanner.Scan() {
		if scanner.Text() == "" {
			passport.saveIfValid()
			passport = Passport{}
			continue
		}

		passport.setData(scanner.Text())
	}

	passport.saveIfValid()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return len(passports)
}

func (p *Passport) setData(data string) {
	sp := strings.Split(data, " ")
	for _, s := range sp {
		fieldSplit := strings.Split(s, ":")
		switch fieldSplit[0] {
		case HairColorType:
			p.HairColor = fieldSplit[1]
		case EyeColorType:
			p.EyeColor = fieldSplit[1]
		case BirthYearType:
			p.BirthYear, _ = strconv.Atoi(fieldSplit[1])
		case IssueYearType:
			p.IssueYear, _ = strconv.Atoi(fieldSplit[1])
		case ExpirationYearType:
			p.ExpirationYear, _ = strconv.Atoi(fieldSplit[1])
		case HeighType:
			p.Height = fieldSplit[1]
		case PassportIDType:
			p.PassportID = fieldSplit[1]
		}
	}
}

func (p *Passport) saveIfValid() {
	if p.BirthYear == 0 {
		return
	}

	if p.ExpirationYear == 0 {
		return
	}

	if p.EyeColor == "" {
		return
	}

	if p.HairColor == "" {
		return
	}

	if p.Height == "" {
		return
	}

	if p.IssueYear == 0 {
		return
	}

	if p.PassportID == "" {
		return
	}

	passports = append(passports, *p)
}
