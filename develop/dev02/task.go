package main

import (
	"fmt"
	"unicode"
)

func unpackString(str string) string {
	var result []rune
	var copyRune rune
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		runeValue := runes[i]
		if unicode.IsLetter(runeValue) {
			copyRune = runeValue
			result = append(result, copyRune)
		} else if unicode.IsDigit(runeValue) {
			if copyRune == 0 {
				return "(некорректная строка)"
			}
			count := int(runeValue - '0')
			for j := 1; j < count; j++ { 
				result = append(result, copyRune)
			}
		} else {
			return "(некорректная строка)"
		}
	}

	return string(result)
}

func main() {
	fmt.Println(unpackString("a4bc2d5e"))
	fmt.Println(unpackString("abcd"))   
	fmt.Println(unpackString("45"))    
	fmt.Println(unpackString(""))        
}

