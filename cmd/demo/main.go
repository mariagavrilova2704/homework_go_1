package main

import (
	"fmt"
	"io"
	"os"

	"github.com/rinat-course/homework1/internal/homework"
)

// demoInput хранит входные строки для демонстрационной программы.
type demoInput struct {
	name        string
	courseName  string
	lessonName  string
	owner       string
	repo        string
	packagePath string
}

const maxArgs = 6

// main — точка входа в программу.
//
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

// parseInput перекладывает аргументы командной строки в demoInput.
//
// Все аргументы — строки. Никаких int и bool в первой домашке нет.
func parseInput(args []string) (demoInput, error) {
	input := defaultInput()

	if len(args) > maxArgs {
		return demoInput{}, fmt.Errorf("too many arguments: got %d, max %d", len(args), maxArgs)
	}

	for index, value := range args {
		applyArgument(&input, index, value)
	}

	return input, nil
}

// defaultInput возвращает значения, которые используются без CLI-аргументов.
func defaultInput() demoInput {
	return demoInput{
		name:        "Мария",
		courseName:  "Go backend",
		lessonName:  "терминал, Git и первый Go-проект",
		owner:       "rinat-course",
		repo:        "homework1",
		packagePath: "./cmd/demo",
	}
}

// applyArgument записывает один CLI-аргумент в нужное поле input.
func applyArgument(input *demoInput, index int, value string) {
	switch index {
	case 0:
		input.name = value
	case 1:
		input.courseName = value
	case 2:
		input.lessonName = value
	case 3:
		input.owner = value
	case 4:
		input.repo = value
	case 5:
		input.packagePath = value
	}
}

// printResult вызывает функции из домашнего задания и печатает результат.
func printResult(writer io.Writer, input demoInput) {
	fmt.Fprintln(writer, homework.BuildGreeting(input.name))
	fmt.Fprintln(writer, homework.BuildCourseWelcome(input.courseName))
	fmt.Fprintln(writer, homework.BuildLessonTitle(input.lessonName))
	fmt.Fprintln(writer, homework.BuildRepositoryPath(input.owner, input.repo))
	fmt.Fprintln(writer, homework.BuildRunCommand(input.packagePath))
}

// printUsage печатает подсказку, как правильно запускать программу.
func printUsage(writer io.Writer) {
	fmt.Fprintln(writer, "Usage:")
	fmt.Fprintln(writer, "  main [name courseName lessonName owner repo packagePath]")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Examples:")
	fmt.Fprintln(writer, "  ./main")
	fmt.Fprintln(writer, "  ./main Алексей 'Основы Go' 'package и fmt' student homework1 ./cmd/demo")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "If an argument is omitted, the default value is used.")
}
