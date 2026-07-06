package homework

import "fmt"

// BuildGreeting возвращает приветствие для ученика.
//
// Ожидаемый формат:
// "Привет, {name}! Добро пожаловать в Go."
func BuildGreeting(name string) string {
	return fmt.Sprintf("Привет, %s! Добро пожаловать в Go.", name)
}

// BuildCourseWelcome возвращает название курса.
//
// Ожидаемый формат:
// "Курс: {courseName}"
func BuildCourseWelcome(courseName string) string {
	return fmt.Sprintf("Курс: %s", courseName)
}

// BuildLessonTitle возвращает название первого урока.
//
// Ожидаемый формат:
// "Урок 1: {lessonName}"
func BuildLessonTitle(lessonName string) string {
	return fmt.Sprintf("Урок 1: %s", lessonName)
}

// BuildRepositoryPath возвращает путь до репозитория на GitHub.
//
// Ожидаемый формат:
// "github.com/{owner}/{repo}"
func BuildRepositoryPath(owner string, repo string) string {
	return fmt.Sprintf("github.com/%s/%s", owner, repo)
}

// BuildRunCommand возвращает команду запуска Go-программы.
//
// Ожидаемый формат:
// "go run {packagePath}"
func BuildRunCommand(packagePath string) string {
	return fmt.Sprintf("go run %s", packagePath)
}
