package main

import (
	"fmt"
	"regexp"
	"strings"
)

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
func AskChar() byte {
	var buf byte
	fmt.Print("Enter Letter (written in lowercase): ")
	fmt.Scanln(&buf)
	return buf
}
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

func Caesar() {
	fmt.Println("Caesar cipher")
	fmt.Println("This algorithm replaces each letter with the letter after nth in the alphabet")

	str := AskWord()
	key := AskNum(1, 25)

	fmt.Println("How to shift alphabet with the given key")
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c => %c\n", i, rune((int(i)-'a'+key)%26+'a'))
	}

	result := []byte(str)
	for i := range result {
		letter := result[i]
		result[i] = byte((int(letter-'a')+key)%26 + 'a')
	}

	fmt.Printf("%s => %s\n", str, result)
}

func main() {
	switch AskAlgorithm() {
	case "caesar":
		Caesar()
	}
}
