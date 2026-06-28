# Homework 1 — первый Go-код, функции, типы и GitHub Pull Request

Это первое домашнее задание по курсу Go/backend.

Главная цель — не написать сложный алгоритм, а закрепить полный рабочий процесс:

1. Сделать fork репозитория.
2. Склонировать репозиторий к себе на компьютер.
3. Создать отдельную ветку для домашки.
4. Реализовать функции в Go.
5. Запустить unit-тесты, integration-тесты и линтеры локально.
6. Сделать commit и push.
7. Открыть Pull Request на проверку преподавателю.
8. Исправить замечания и добиться зелёного CI.

## Где писать код

Все решения нужно писать в файле:

```text
internal/homework/homework.go
```

Unit-тесты находятся в файле:

```text
internal/homework/homework_test.go
```

Integration-тесты находятся в файле:

```text
test/integration/demo_test.go
```

Integration-тесты собирают настоящий бинарник из `cmd/demo`, запускают его и проверяют вывод.
Так ученики видят разницу между проверкой отдельных функций и проверкой программы целиком.

## Как запустить приложение

```bash
make run
```

## Как запустить unit-тесты

```bash
make test-unit
```

Unit-тесты проверяют отдельные функции из `internal/homework/homework.go`.

## Как запустить integration-тесты

```bash
make test-integration
```

Integration-тесты делают следующее:

1. Собирают бинарник из `cmd/demo`.
2. Запускают собранный бинарник.
3. Проверяют, что программа напечатала ожидаемые строки.

## Как запустить все тесты

```bash
make run
```

Запуск с внешними входными параметрами:

```bash
make run ARGS="Алексей 17 250 4 8 16 true false"
```

После сборки бинарника можно запустить его напрямую:

```bash
make build
./bin/main Алексей 17 250 4 8 16 true false
```

Порядок параметров:

```text
name age price count completedLessons totalLessons goCoreCompleted homeworkDone
```

Если параметры не переданы, программа использует значения из исходного примера:

```text
Мария 18 100 3 3 12 true true
```

## Как запустить все локальные проверки

```bash
make check
```

`make check` проверяет:

- форматирование Go-кода;
- `go vet`;
- style-линтеры;
- static-analysis линтеры;
- security-линтеры;
- unit-тесты;
- integration-тесты.

## Как запустить полный CI локально

```bash
make ci
```

`make ci` максимально близок к тому, что запускается в GitHub Actions.

## Какие проверки будут в Pull Request

В GitHub Actions проверки разнесены на отдельные параллельные job-ы:

```text
check:dependencies
linter:format
linter:style
linter:static-analysis
security:go
go:unit-tests
go:integration-tests
go:race-tests
tests:coverage
build:binary
cd:package-artifact
```

Для Pull Request `cd:package-artifact` не запускается. Он нужен для push в `main`.

## Как отправить домашку

```bash
git checkout -b homework-1
# пишем код
make check
git status
git add .
git commit -m "complete homework 1"
git push origin homework-1
```

После этого откройте Pull Request на GitHub и добавьте преподавателя на review.

## Задания

Подробные условия находятся в [docs/task.md](docs/task.md).

## Ограничения

- Нельзя менять тесты без согласования с преподавателем.
- Нельзя удалять функции из `homework.go`.
- Можно добавлять вспомогательные функции, если они делают код понятнее.
- Код должен быть отформатирован через `gofmt`.
- Pull Request должен проходить unit-тесты, integration-тесты и линтеры.
- Pull Request должен проходить unit-тесты, integration-тесты и линтеры.
