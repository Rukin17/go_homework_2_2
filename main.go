package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const vowels string = "аеёиоуыэюя"

func main() {
	reader := bufio.NewReader(os.Stdin)
	var input string
	fmt.Print("Введите строку: ")
	input, _ = reader.ReadString('\n')
	countVowel := 0
	for _, char := range input {
		lowerChar := unicode.ToLower(char)
		if strings.ContainsRune(vowels, lowerChar) {
			countVowel++
		}
	}
	fmt.Println(countVowel)
}
