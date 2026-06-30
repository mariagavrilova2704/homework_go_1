# Homework 1 — условия задач

Нужно реализовать функции в файле:

```text
internal/homework/homework.go
```

В этой домашке тренируем только темы первого урока:

- `package`;
- `import`;
- `func`;
- `return`;
- строки;
- библиотеку `fmt`;
- запуск тестов;
- отправку Pull Request.

Здесь специально нет `int`, `bool`, `if`, циклов и сложной логики. Они появятся в следующих уроках.

---

## Задача 1. BuildGreeting

### Сигнатура

```go
func BuildGreeting(name string) string
```

### Что должна делать функция

Функция принимает имя ученика и возвращает приветствие.

### Формат результата

```text
"Привет, {name}! Добро пожаловать в Go."
```

### Примеры

| Вход | Выход |
|---|---|
| `"Мария"` | `"Привет, Мария! Добро пожаловать в Go."` |
| `"Игорь"` | `"Привет, Игорь! Добро пожаловать в Go."` |

---

## Задача 2. BuildCourseWelcome

### Сигнатура

```go
func BuildCourseWelcome(courseName string) string
```

### Что должна делать функция

Функция принимает название курса и возвращает строку с названием курса.

### Формат результата

```text
"Курс: {courseName}"
```

### Примеры

| Вход | Выход |
|---|---|
| `"Go backend"` | `"Курс: Go backend"` |
| `"Основы Go"` | `"Курс: Основы Go"` |

---

## Задача 3. BuildLessonTitle

### Сигнатура

```go
func BuildLessonTitle(lessonName string) string
```

### Что должна делать функция

Функция принимает название темы и возвращает название урока.

### Формат результата

```text
"Урок 1: {lessonName}"
```

### Примеры

| Вход | Выход |
|---|---|
| `"терминал, Git и первый Go-проект"` | `"Урок 1: терминал, Git и первый Go-проект"` |
| `"package, import и fmt"` | `"Урок 1: package, import и fmt"` |

---

## Задача 4. BuildRepositoryPath

### Сигнатура

```go
func BuildRepositoryPath(owner string, repo string) string
```

### Что должна делать функция

Функция принимает владельца репозитория и название репозитория, затем возвращает путь до репозитория на GitHub.

### Формат результата

```text
"github.com/{owner}/{repo}"
```

### Примеры

| Вход | Выход |
|---|---|
| `owner = "rinat-course"`, `repo = "homework1"` | `"github.com/rinat-course/homework1"` |
| `owner = "student"`, `repo = "go-homework"` | `"github.com/student/go-homework"` |

---

## Задача 5. BuildRunCommand

### Сигнатура

```go
func BuildRunCommand(packagePath string) string
```

### Что должна делать функция

Функция принимает путь до Go-пакета и возвращает команду запуска программы.

### Формат результата

```text
"go run {packagePath}"
```

### Примеры

| Вход | Выход |
|---|---|
| `"./cmd/demo"` | `"go run ./cmd/demo"` |
| `"./cmd/app"` | `"go run ./cmd/app"` |

---

## Как понять, что всё сделано правильно

Запустите:

```bash
make test-unit
make test-integration
make check
```

Домашка готова к отправке, если все команды завершились без ошибки.

## Что можно менять

Можно менять:

```text
internal/homework/homework.go
```

Не нужно менять:

```text
internal/homework/homework_test.go
test/integration/demo_test.go
.github/workflows/ci.yml
```

## Критерии проверки

Задание считается выполненным, если:

1. Все функции возвращают строки в точном ожидаемом формате.
2. Все тесты проходят.
3. `make check` не падает.
4. Pull Request открыт.
5. CI в Pull Request зелёный.
6. Ученик может объяснить каждую функцию своими словами.
