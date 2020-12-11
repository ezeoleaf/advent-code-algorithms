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
var validColors = map[string]interface{}{
	"amb": nil,
	"blu": nil,
	"brn": nil,
	"gry": nil,
	"grn": nil,
	"hzl": nil,
	"oth": nil,
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
		case HairColorType: //a # followed by exactly six characters 0-9 or a-f
			if fieldSplit[1][0:1] == "#" && len(fieldSplit[1]) == 7 {
				p.HairColor = fieldSplit[1]
			}

		case EyeColorType: //exactly one of: amb blu brn gry grn hzl oth.
			if _, ok := validColors[fieldSplit[1]]; ok {
				p.EyeColor = fieldSplit[1]
			}
		case BirthYearType: //four digits; at least 1920 and at most 2002.
			by, _ := strconv.Atoi(fieldSplit[1])
			if by >= 1920 && by <= 2002 {
				p.BirthYear = by
			}
		case IssueYearType: //four digits; at least 2010 and at most 2020.
			iy, _ := strconv.Atoi(fieldSplit[1])
			if iy >= 2010 && iy <= 2020 {
				p.IssueYear = iy
			}
		case ExpirationYearType: //four digits; at least 2020 and at most 2030.
			ey, _ := strconv.Atoi(fieldSplit[1])
			if ey >= 2020 && ey <= 2030 {
				p.ExpirationYear = ey
			}
		case HeighType:
			/*
				a number followed by either cm or in:
				If cm, the number must be at least 150 and at most 193.
				If in, the number must be at least 59 and at most 76.
			*/
			hgt := fieldSplit[1]
			if strings.Contains(hgt, "cm") {
				val, _ := strconv.Atoi(hgt[0 : len(hgt)-2])
				if val >= 150 && val <= 193 {
					p.Height = hgt
				}
			} else if strings.Contains(hgt, "in") {
				val, _ := strconv.Atoi(hgt[0 : len(hgt)-2])
				if val >= 59 && val <= 76 {
					p.Height = hgt
				}
			}
		case PassportIDType: //a nine-digit number, including leading zeroes.
			if len(fieldSplit[1]) == 9 {
				p.PassportID = fieldSplit[1]
			}
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

	if p.IssueYear == 0 {
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

	if p.PassportID == "" {
		return
	}

	passports = append(passports, *p)
}
