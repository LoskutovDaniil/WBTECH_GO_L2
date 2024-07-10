package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// Чтение содержимого файла и возврат его в виде строки
func getFileContent(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", errors.New("grepUtil: " + path + ": No such file or directory")
	}
	defer file.Close()

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buffer.WriteString(scanner.Text() + "\n")
	}
	return buffer.String(), nil
}

// Вывод строк, содержащих искомый шаблон
func filterAndPrint(output io.Writer, lines []string, search string, ignoreCase bool) {
	search = strings.ToLower(search)
	for _, line := range lines {
		lineToCheck := line
		if ignoreCase {
			lineToCheck = strings.ToLower(line)
		}
		if strings.Contains(lineToCheck, search) {
			highlightedLine := strings.ReplaceAll(line, search, "\033[31m"+search+"\033[0m")
			fmt.Fprintln(output, highlightedLine)
		}
	}
}

// Основная функция grep
func grep() {
	ignoreCase := flag.Bool("i", false, "ignore-case")
	flag.Parse()

	var searchTerm, filePath string

	if !*ignoreCase {
		searchTerm = os.Args[1]
		filePath = os.Args[2]
	} else {
		searchTerm = os.Args[2]
		filePath = os.Args[3]
	}

	content, err := getFileContent(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(content, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	filterAndPrint(os.Stdout, lines, searchTerm, *ignoreCase)
}

func main() {
	grep()
}
