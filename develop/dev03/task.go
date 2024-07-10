package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type inputData struct {
	filePath string
	option   string
	column   int
}

// читает строки из файла и помещает их в срез
func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var lines []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err == io.EOF {
			line += "\n"
			lines = append(lines, line)
			break
		}
		lines = append(lines, line)
	}

	return lines, nil
}

// возвращает срез с уникальными элементами
func getUniqueStrings(lines []string) []string {
	uniqueMap := make(map[string]string)
	var result []string

	for _, line := range lines {
		_, exists := uniqueMap[line]
		if !exists {
			uniqueMap[line] = line
		}
	}

	for _, line := range uniqueMap {
		result = append(result, line)
	}

	return result
}

func displayLines(lines []string, option string) {
	if option == "-r" {
		for i := len(lines) - 1; i >= 0; i-- {
			fmt.Print(lines[i])
		}
	} else {
		for _, line := range lines {
			fmt.Print(line)
		}
	}
}

func sortNumerically(lines []string) {
	sort.Slice(lines, func(i, j int) bool {
		iTrimmed := strings.Replace(lines[i], " ", "", -1)
		iTrimmed = strings.Replace(iTrimmed, "\n", "", -1)
		iNum, iErr := strconv.Atoi(iTrimmed)

		jTrimmed := strings.Replace(lines[j], " ", "", -1)
		jTrimmed = strings.Replace(jTrimmed, "\n", "", -1)
		jNum, jErr := strconv.Atoi(jTrimmed)

		if iErr != nil || jErr != nil {
			return lines[i] < lines[j]
		}

		return iNum < jNum
	})
}

// сортирует по определенной колонке, указанной индексом строки
func sortByColumn(lines []string, index int) {
	sort.Slice(lines, func(i, j int) bool {
		words1 := strings.Split(lines[i], " ")
		words2 := strings.Split(lines[j], " ")

		if index > len(words1) || index > len(words2) {
			return false
		}

		key1 := words1[index-1]
		key2 := words2[index-1]

		return key1 < key2
	})
}

func parseArgs(args []string) (*inputData, error) {
	var data inputData
	if len(args) < 2 || len(args) > 4 || (len(args) == 4 && args[1] != "-k") {
		err := errors.New("Invalid input")
		return nil, err
	}

	if len(args) == 4 {
		data.option = args[1]

		idx, err := strconv.Atoi(args[2])
		if err != nil {
			err := errors.New("Invalid input")
			return nil, err
		}

		data.column = idx
		data.filePath = args[3]

	} else if len(args) == 3 {
		data.option = args[1]
		data.filePath = args[2]

	} else {
		data.filePath = args[1]
	}

	return &data, nil
}

// поддержка обработки только одного файла и одного флага/без флага одновременно
func processFile() error {
	input, err := parseArgs(os.Args)
	if err != nil {
		return err
	}

	lines, err := readFile(input.filePath)
	if err != nil {
		return err
	}

	if input.option == "-u" {
		lines = getUniqueStrings(lines)
	}

	if input.option == "-n" {
		sortNumerically(lines)

	} else if input.option == "-k" {
		sortByColumn(lines, input.column)

	} else {
		sort.Strings(lines)
	}

	displayLines(lines, input.option)

	return nil
}

func main() {
	err := processFile()
	if err != nil {
		fmt.Println(err)
	}
}
