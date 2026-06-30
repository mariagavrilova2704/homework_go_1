# Homework 1 — первый Go-код, `fmt` и GitHub Pull Request

Это первая домашняя работа по курсу Go/backend.

Главная цель — спокойно пройти полный рабочий маршрут разработчика:

1. Сделать fork репозитория.
2. Склонировать свой fork на компьютер.
3. Создать отдельную ветку для домашки.
4. Реализовать простые функции через `fmt.Sprintf`.
5. Запустить тесты локально.
6. Сделать commit и push.
7. Открыть Pull Request.
8. Дождаться зелёного CI и review преподавателя.

В первой домашке **нет `int`, `bool`, условий и сложной логики**. Всё задание построено только на том, что уже было на первом уроке: `package`, `import`, `func`, `return`, строки и библиотека `fmt`.

## Где писать код

Все решения нужно писать только в файле:

```text
internal/homework/homework.go
```

Тесты уже готовы. Их менять не нужно.

Unit-тесты лежат здесь:

```text
internal/homework/homework_test.go
```

Integration-тесты лежат здесь:

```text
test/integration/demo_test.go
```

## Что нужно сделать

Откройте файл:

```text
internal/homework/homework.go
```

И реализуйте функции:

```go
BuildGreeting(name string) string
BuildCourseWelcome(courseName string) string
BuildLessonTitle(lessonName string) string
BuildRepositoryPath(owner string, repo string) string
BuildRunCommand(packagePath string) string
```

Во всех функциях удобно использовать `fmt.Sprintf`.

Пример:

```go
func BuildGreeting(name string) string {
    return fmt.Sprintf("Привет, %s! Добро пожаловать в Go.", name)
}
```

Не забудьте добавить импорт `fmt`.

## Как запустить приложение

```bash
make run
```

Ожидаемый вывод после решения:

```text
Привет, Мария! Добро пожаловать в Go.
Курс: Go backend
Урок 1: терминал, Git и первый Go-проект
github.com/rinat-course/homework1
go run ./cmd/demo
```

Запуск с другими строками:

```bash
make run ARGS="Алексей Основы-Go package-и-fmt student homework1 ./cmd/demo"
```

## Как запустить тесты

Только unit-тесты:

```bash
make test-unit
```

Integration-тесты:

```bash
make test-integration
```

Все тесты:

```bash
make test
```

Все локальные проверки перед отправкой:

```bash
make check
```

## Как отправить домашку

```bash
git checkout -b homework-1
# пишем код в internal/homework/homework.go
make test-unit
make test-integration
make check
git status
git add .
git commit -m "complete homework 1"
git push origin homework-1
```

После push откройте Pull Request на GitHub и добавьте преподавателя на review.

## Подробная инструкция

- Условия задач: [docs/task.md](docs/task.md)
- GitHub workflow: [docs/github-flow.md](docs/github-flow.md)

## Ограничения

- Нельзя менять тесты без согласования с преподавателем.
- Нельзя удалять функции из `homework.go`.
- Нужно использовать `fmt.Sprintf` хотя бы в тех функциях, где собирается строка из частей.
- Код должен быть отформатирован через `gofmt`.
- Pull Request должен проходить тесты и CI.
