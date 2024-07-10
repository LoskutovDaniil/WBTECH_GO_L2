package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	goPs "github.com/mitchellh/go-ps"
)

func cd(args []string) {
	switch len(args) {
	case 1:
		fmt.Fprintln(os.Stderr, "Insufficient number of arguments")
	case 2:
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	default:
		fmt.Fprintln(os.Stderr, "too many arguments")
	}
}

func pwd(args []string) {
	if len(args) == 1 {
		path, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Println(path)
		}
	} else {
		fmt.Fprintln(os.Stderr, "too many arguments")
	}
}

func echo(args []string) {
	for i := 1; i < len(args); i++ {
		fmt.Printf("%s ", args[i])
	}
	fmt.Println()
}

func kill(args []string) {
	if len(args) == 1 {
		fmt.Fprintln(os.Stderr, "missing PID argument")
		return
	}
	if len(args) > 2 {
		fmt.Fprintln(os.Stderr, "too many arguments")
		return
	}
	pid, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	err = proc.Kill()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func ps(args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "too many arguments")
		return
	}
	processList, _ := goPs.Processes()

	for _, proc := range processList {
		fmt.Printf("Process name: %v process id: %v\n", proc.Executable(), proc.Pid())
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		args := strings.Split(input.Text(), " ")
		switch args[0] {
		case "cd":
			cd(args)
		case "pwd":
			pwd(args)
		case "echo":
			echo(args)
		case "kill":
			kill(args)
		case "ps":
			ps(args)
		}
	}
}
