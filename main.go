package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MetaCommandResult int
type PrepareResult int
type StatementType int

const (
	META_COMMAND_SUCCESS      MetaCommandResult = 0
	META_COMMAND_UNRECOGNIZED MetaCommandResult = 1
)

const (
	PREPARE_SUCCESS      PrepareResult = 0
	PREPARE_UNRECOGNIZED PrepareResult = 1
)

const (
	STATEMENT_INSERT StatementType = 0
	STATEMENT_SELECT StatementType = 1
)

type Statement struct {
	stype StatementType
}

func print_prompt() {
	fmt.Print("db>")
}

func doMetaCommand(input string) MetaCommandResult {
	if strings.Compare(input, ".exit") == 0 {
		os.Exit(0)
	}
	return META_COMMAND_UNRECOGNIZED
}

func prepareStatment(input string, statement *Statement) PrepareResult {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	switch scanner.Text() {
	case "insert":
		statement.stype = STATEMENT_INSERT
		return PREPARE_SUCCESS
	case "select":
		statement.stype = STATEMENT_SELECT
		return PREPARE_SUCCESS
	default:
		return PREPARE_UNRECOGNIZED
	}
}

func executeStatment(statement Statement) {
	switch statement.stype {
	case STATEMENT_INSERT:
		fmt.Println("this is insert")
	case STATEMENT_SELECT:
		fmt.Println("this is select")
	}
}

func main() {
	for {
		var input string
		print_prompt()
		fmt.Scan()
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Printf("Error reading input,err:%s\n", err)
			os.Exit(-1)
		}

		input = input[:len(input)-1]
		if len(input) == 0 {
			continue
		}
		if input[0] == '.' {
			switch doMetaCommand(input) {
			case META_COMMAND_SUCCESS:
				continue
			case META_COMMAND_UNRECOGNIZED:
				fmt.Println("unrecognized command")
				continue
			}
		}

		var statement Statement
		switch prepareStatment(input, &statement) {
		case PREPARE_SUCCESS:
			break
		case PREPARE_UNRECOGNIZED:
			fmt.Printf("unrecognized keyword at start of %s\n", input)
			continue
		}

		executeStatment(statement)
		fmt.Println("Executed")
	}
}
