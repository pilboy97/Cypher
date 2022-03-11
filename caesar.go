package main

import "fmt"

//Explain Caesar Algorithm
func Caesar() {
	fmt.Println("Caesar cipher")
	fmt.Println("This algorithm replaces each letter with the letter after nth in the alphabet")

	//input a word
	str := AskWord()
	//input a key number
	key := AskNum(1, 25)

	fmt.Println("How to shift alphabet with the given key")
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c => %c\n", i, rune((int(i)-'a'+key)%26+'a'))
	}

	result := []byte(str)
	for i := range result {
		//letter is the ith letter
		letter := result[i]
		//replace with the letter after the given key in the alphabet
		result[i] = byte((int(letter-'a')+key)%26 + 'a')
	}

	fmt.Printf("%s => %s\n", str, result)
}
