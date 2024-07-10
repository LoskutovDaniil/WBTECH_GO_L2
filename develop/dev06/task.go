package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// Функция для извлечения полей
func extractFields(reader io.Reader, sep string, cols []int, onlySep bool) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()

		if onlySep && !strings.Contains(text, sep) {
			continue
		}

		segments := strings.Split(text, sep)
		var result []string
		for _, col := range cols {
			if col > 0 && col <= len(segments) {
				result = append(result, segments[col-1])
			}
		}
		fmt.Println(strings.Join(result, "\t"))
	}
}

func main() {
	var (
		columns      string
		separator    string
		onlySeparated bool
	)
	flag.StringVar(&columns, "f", "1", "Выбрать столбцы")
	flag.StringVar(&separator, "d", "\t", "Использовать другой разделитель")
	flag.BoolVar(&onlySeparated, "s", false, "Только строки с разделителем")
	flag.Parse()

	columnIndices := parseColumns(columns)

	extractFields(os.Stdin, separator, columnIndices, onlySeparated)
}

// Функция для парсинга столбцов
func parseColumns(cols string) []int {
	colStr := strings.Split(cols, ",")
	var colIndices []int
	for _, col := range colStr {
		index := 0
		fmt.Sscanf(col, "%d", &index)
		colIndices = append(colIndices, index)
	}
	return colIndices
}

// Для запуска вам нужно ввести: cat text.txt | go run main.go -f 2 -d ","