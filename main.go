package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("day1\n answer1: %d\n answer2: %d", day1Task1(), day1Task2())
	fmt.Printf("\nday2\n answer1: %d\n answer2: %d", day2Task1(), day2Task2())
	fmt.Printf("\nday3\n answer1: %d\n answer2: %d", day3Task1(), day3Task2())
	fmt.Printf("\nday4\n answer1: %d\n answer2: %d", day4Task1(), day4Task2())
}
func day4Task2() int {
	input := readInput("assets/input4.txt")
	input = strings.Replace(input, "\n", " ", -1)
	lines := strings.Split(input, "  ")
	requiredAttrsStr := "byr,iyr,eyr,hgt,hcl,ecl,pid"

	validCount := 0
	for _, line := range lines {
		attrs := strings.Split(line, " ")
		attrsCount := 0
		for _, attr := range attrs {
			pair := strings.Split(attr, ":")
			key := pair[0]
			value := pair[1]
			if strings.Contains(requiredAttrsStr, key) && validPassortValue(key, value) {
				attrsCount++
			}
		}
		if attrsCount == 7 {
			validCount++
		}
	}
	return validCount
}

func validPassortValue(key, value string) bool {
	switch key {
	case "byr":
		if n, err := strconv.Atoi(value); err == nil {
			return 1920 <= n && n <= 2002
		}
	case "iyr":
		if n, err := strconv.Atoi(value); err == nil {
			return 2010 <= n && n <= 2020
		}
	case "eyr":
		if n, err := strconv.Atoi(value); err == nil {
			return 2020 <= n && n <= 2030
		}
	case "hgt":
		if strings.HasSuffix(value, "cm") {
			value = strings.Replace(value, "cm", "", -1)
			if n, err := strconv.Atoi(value); err == nil {
				return 150 <= n && n <= 193
			}
		} else if strings.HasSuffix(value, "in") {
			value = strings.Replace(value, "in", "", -1)
			if n, err := strconv.Atoi(value); err == nil {
				return 59 <= n && n <= 76
			}
		}
	case "hcl":
		reg := regexp.MustCompile(`^#[0-9a-f]{6}$`)
		return reg.MatchString(value)
	case "ecl":
		return strings.Contains("amb,blu,brn,gry,grn,hzl,oth", value)
	case "pid":
		reg := regexp.MustCompile(`^[0-9]{9}$`)
		return reg.MatchString(value)
	case "cid":
		return true
	default:
		return false
	}

	return false
}

func day4Task1() int {
	input := readInput("assets/input4.txt")
	input = strings.Replace(input, "\n", " ", -1)
	lines := strings.Split(input, "  ")

	requiredAttrsStr := "byr,iyr,eyr,hgt,hcl,ecl,pid"
	validCount := 0
	for _, line := range lines {
		attrs := strings.Split(line, " ")
		attrsCount := 0
		for _, attr := range attrs {
			pair := strings.Split(attr, ":")
			key := pair[0]
			if strings.Contains(requiredAttrsStr, key) {
				attrsCount++
			} else {
				break
			}
		}
		if attrsCount == 7 {
			validCount++
		}
	}
	return validCount
}

func day3Task2() int64 {
	input := readInput("assets/input3.txt")
	width := strings.Index(input, "\n")
	height := strings.Count(input, "\n")
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	trees := make([]int, 0)
	for _, slope := range slopes {
		right := slope[0]
		down := slope[1]
		tree := 0
		for down <= height {
			if string(input[(width+1)*down+right]) == "#" {
				tree++
			}
			right = (right + slope[0]) % width
			down += slope[1]
		}
		trees = append(trees, tree)
	}

	var pod int64 = 1
	for _, tree := range trees {
		pod *= int64(tree)
	}

	return pod
}

func day3Task1() int {
	input := readInput("assets/input3.txt")
	rightStep := 3
	downStep := 1
	right := rightStep
	down := downStep
	width := strings.Index(input, "\n")
	height := strings.Count(input, "\n")
	trees := 0
	for down <= height {
		if string(input[(width+1)*down+right]) == "#" {
			trees++
		}
		right = (right + rightStep) % width
		down += downStep
	}
	return trees
}

func day2Task1() int64 {
	input := readInput("assets/input2.txt")
	correctPassCount := 0
	for _, line := range strings.Split(input, "\n") {
		s := strings.SplitAfter(line, ": ")
		policy := strings.TrimRight(s[0], ": ")
		pass := s[1]
		tmp := strings.Split(policy, " ")
		char := tmp[1]
		tmp = strings.Split(tmp[0], "-")
		min, _ := strconv.Atoi(tmp[0])
		max, _ := strconv.Atoi(tmp[1])
		count := 0
		for _, ch := range pass {
			if string(ch) == char {
				count++
			}
		}
		if count >= min && count <= max {
			correctPassCount++
		}
	}

	return int64(correctPassCount)
}

func day2Task2() int64 {
	input := readInput("assets/input2.txt")
	correctPassCount := 0
	for _, line := range strings.Split(input, "\n") {
		s := strings.SplitAfter(line, ": ")
		policy := strings.TrimRight(s[0], ": ")
		pass := s[1]
		tmp := strings.Split(policy, " ")
		char := tmp[1]
		tmp = strings.Split(tmp[0], "-")
		low, _ := strconv.Atoi(tmp[0])
		up, _ := strconv.Atoi(tmp[1])
		lowChar := string(pass[low-1])
		upChar := string(pass[up-1])
		if lowChar != char && upChar == char || lowChar == char && upChar != char {
			correctPassCount++
		}
	}

	return int64(correctPassCount)
}

func day1Task1() int64 {
	input := readInput("assets/input1.txt")
	numbers := make([]int64, 0)
	for _, line := range strings.Split(input, "\n") {
		if n, err := strconv.ParseInt(line, 10, 64); err == nil {
			numbers = append(numbers, n)
		}
	}
	for i := 0; i < len(numbers)-1; i++ {
		for j := 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == 2020 {
				return numbers[i] * numbers[j]
			}
		}
	}
	return -1
}

func day1Task2() int64 {
	input := readInput("assets/input1.txt")
	numbers := make([]int64, 0)
	for _, line := range strings.Split(input, "\n") {
		if n, err := strconv.ParseInt(line, 10, 64); err == nil {
			numbers = append(numbers, n)
		}
	}
	for i := 0; i < len(numbers)-2; i++ {
		for j := i + 1; j < len(numbers)-1; j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					return numbers[i] * numbers[j] * numbers[k]
				}
			}
		}
	}
	return -1
}

func readInput(filepath string) string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(bytes.NewBuffer(file).String())
}
