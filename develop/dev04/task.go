package main

import (
	"fmt"
	"sort"
	"strings"
)

// Функция для сортировки букв в строке
func sortString(word string) string {
	wordsRuneSlice := strings.Split(word, "")
	sort.Strings(wordsRuneSlice)
	return strings.Join(wordsRuneSlice, "")
}

// Функция для создания мапы анаграмм
func anagram(anagrams []string) (map[string][]string, error) {
	mapAnagram := make(map[string][]string)

	for _, word := range anagrams {
		word = strings.ToLower(word)
		// Сортируем буквы в строке, чтобы сформировать ключ
		sortedKey := sortString(word)
		found := false
		for _, w := range mapAnagram[sortedKey] {
			if w == word {
				found = true
				break
			}
		}
		if !found {
			mapAnagram[sortedKey] = append(mapAnagram[sortedKey], word)
		}
	}

	return mapAnagram, nil
}

func main() {
	anagrams := []string{"пятка", "пятак", "тяпка", "пятак", "листок", "сЛИток", "столик", "листок"}
	data, err := anagram(anagrams)
	if err != nil {
		fmt.Println(err)
		return
	}

	var keys []string
	for key := range data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		sort.Strings(data[key])
		fmt.Printf("Key: %s, Anagrams: %v\n", key, data[key])
	}
}
