package main

import (
	"fmt"
	"regexp"
	"strings"
)

//Input a name of cipher algorithm
func AskAlgorithm() string {
	var buf string
	for {
		fmt.Print("Enter Algorithm: ")
		fmt.Scanln(&buf)

		buf = strings.ToLower(buf)

		switch buf {
		case "caesar":
			return "caesar"
		}
	}
}

//Input a word
func AskWord() string {
	rexp := regexp.MustCompile("[[:lower:]]+")
	var buf string
	for {
		fmt.Print("Enter Word (written in lowercase): ")
		fmt.Scanln(&buf)
		if rexp.MatchString(buf) {
			return buf
		}
	}
}

//Input a letter
func AskChar() byte {
	var buf byte
	fmt.Print("Enter Letter (written in lowercase): ")
	fmt.Scanln(&buf)
	return buf
}

//Check the given number is a prime number
func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

//Input a prime number less than 50
func AskPrime() int {
	var num int
	for {
		fmt.Print("Enter Prime number less than 50: ")
		fmt.Scanln(&num)
		if num < 50 && IsPrime(num) {
			return num
		}
	}
}

//Input a number between a and b
func AskNum(a, b int) int {
	var num int
	if a > b {
		a, b = b, a
	}
	for {
		fmt.Printf("Enter number between %d and %d: ", a, b)
		fmt.Scanln(&num)

		if a <= num && num <= b {
			return num
		}
	}
}

func main() {
	switch AskAlgorithm() {
	case "caesar":
		Caesar()
	}
}
