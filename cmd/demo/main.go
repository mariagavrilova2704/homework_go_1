package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/rinat-course/homework1/internal/homework"
)

// demoInput хранит все входные данные для запуска демонстрационной программы.
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

const maxArgs = 8

// main — точка входа в программу.
//
// Программа берёт аргументы из командной строки.
// Если аргументы не переданы, используются значения по умолчанию.
func main() {
	input, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		printUsage(os.Stderr)
		os.Exit(1)
	}

	printResult(os.Stdout, input)
}

// parseInput преобразует аргументы командной строки в структуру demoInput.
//
// Например:
// ./main Алексей 17 250 4 8 16 true false
//
// Если часть аргументов не передана, недостающие значения остаются дефолтными.
func parseInput(args []string) (demoInput, error) {
	input := defaultInput()

	if len(args) > maxArgs {
		return demoInput{}, fmt.Errorf("too many arguments: got %d, max %d", len(args), maxArgs)
	}

	for index, value := range args {
		if err := applyArgument(&input, index, value); err != nil {
			return demoInput{}, fmt.Errorf("argument %d: %w", index+1, err)
		}
	}

	return input, nil
}

// defaultInput возвращает входные данные по умолчанию.
//
// Эти значения используются, если пользователь запустил программу без аргументов.
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

// applyArgument записывает один аргумент командной строки в нужное поле input.
//
// Номер аргумента определяет, какое именно поле нужно заполнить.
func applyArgument(input *demoInput, index int, value string) error {
	switch index {
	case 0:
		input.name = value
		return nil
	case 1:
		return setInt(value, &input.age, "age")
	case 2:
		return setInt(value, &input.price, "price")
	case 3:
		return setInt(value, &input.count, "count")
	case 4:
		return setInt(value, &input.completedLessons, "completedLessons")
	case 5:
		return setInt(value, &input.totalLessons, "totalLessons")
	case 6:
		return setBool(value, &input.goCoreCompleted, "goCoreCompleted")
	case 7:
		return setBool(value, &input.homeworkCompleted, "homeworkDone")
	default:
		return fmt.Errorf("unknown argument index: %d", index)
	}
}

// setInt преобразует строку в int и записывает результат в нужное поле.
func setInt(value string, target *int, name string) error {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("%s must be integer", name)
	}

	*target = parsed
	return nil
}

// setBool преобразует строку в bool и записывает результат в нужное поле.
//
// Подходящие значения: true, false, 1, 0, t, f.
func setBool(value string, target *bool, name string) error {
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fmt.Errorf("%s must be boolean", name)
	}

	*target = parsed
	return nil
}

// printResult вызывает функции из домашнего задания и печатает результат.
func printResult(writer io.Writer, input demoInput) {
	fmt.Fprintln(writer, homework.BuildGreeting(input.name))
	fmt.Fprintln(writer, homework.IsAdult(input.age))
	fmt.Fprintln(writer, homework.CalculateTotal(input.price, input.count))
	fmt.Fprintln(writer, homework.FormatCourseProgress(input.completedLessons, input.totalLessons))
	fmt.Fprintln(writer, homework.CanStartBackendBlock(input.goCoreCompleted, input.homeworkCompleted))
}

// printUsage печатает подсказку, как правильно запускать программу.
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
