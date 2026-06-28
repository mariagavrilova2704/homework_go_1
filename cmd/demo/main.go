package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/rinat-course/homework1/internal/homework"
)

type demoInput struct {
	name              string
	age               int
	price             int
	count             int
	completedLessons  int
	totalLessons      int
	goCoreCompleted   bool
	homeworkCompleted bool
}

func main() {
	input, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		printUsage(os.Stderr)
		os.Exit(1)
	}

	printResult(os.Stdout, input)
}

func parseInput(args []string) (demoInput, error) {
	input := defaultInput()
	parsers := inputParsers(&input)

	if len(args) > len(parsers) {
		return demoInput{}, fmt.Errorf("too many arguments: got %d, max %d", len(args), len(parsers))
	}

	for index, arg := range args {
		if err := parsers[index](arg); err != nil {
			return demoInput{}, fmt.Errorf("argument %d: %w", index+1, err)
		}
	}

	return input, nil
}

func defaultInput() demoInput {
	return demoInput{
		name:              "Мария",
		age:               18,
		price:             100,
		count:             3,
		completedLessons:  3,
		totalLessons:      12,
		goCoreCompleted:   true,
		homeworkCompleted: true,
	}
}

func inputParsers(input *demoInput) []func(string) error {
	return []func(string) error{
		func(value string) error {
			input.name = value
			return nil
		},
		func(value string) error { return parseInt(value, &input.age, "age") },
		func(value string) error { return parseInt(value, &input.price, "price") },
		func(value string) error { return parseInt(value, &input.count, "count") },
		func(value string) error {
			return parseInt(value, &input.completedLessons, "completedLessons")
		},
		func(value string) error { return parseInt(value, &input.totalLessons, "totalLessons") },
		func(value string) error {
			return parseBool(value, &input.goCoreCompleted, "goCoreCompleted")
		},
		func(value string) error {
			return parseBool(value, &input.homeworkCompleted, "homeworkDone")
		},
	}
}

func parseInt(value string, target *int, name string) error {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("%s must be integer", name)
	}

	*target = parsed
	return nil
}

func parseBool(value string, target *bool, name string) error {
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fmt.Errorf("%s must be boolean", name)
	}

	*target = parsed
	return nil
}

func printResult(writer io.Writer, input demoInput) {
	fmt.Fprintln(writer, homework.BuildGreeting(input.name))
	fmt.Fprintln(writer, homework.IsAdult(input.age))
	fmt.Fprintln(writer, homework.CalculateTotal(input.price, input.count))
	fmt.Fprintln(writer, homework.FormatCourseProgress(input.completedLessons, input.totalLessons))
	fmt.Fprintln(writer, homework.CanStartBackendBlock(input.goCoreCompleted, input.homeworkCompleted))
}

func printUsage(writer io.Writer) {
	fmt.Fprintln(writer, "Usage:")
	fmt.Fprintln(writer, "  main [name age price count completedLessons totalLessons goCoreCompleted homeworkDone]")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Examples:")
	fmt.Fprintln(writer, "  ./main")
	fmt.Fprintln(writer, "  ./main Алексей 17 250 4 8 16 true false")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "If an argument is omitted, the default value is used.")
}
