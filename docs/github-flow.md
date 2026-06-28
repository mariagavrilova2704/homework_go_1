# Как сдавать домашнюю работу через GitHub

## Один раз

1. Открыть репозиторий домашки.
2. Нажать Fork.
3. Склонировать свой fork на компьютер.

```bash
git clone git@github.com:YOUR_USERNAME/homework1.git
cd homework1
```

## Для каждой домашки

```bash
git checkout -b homework-1
make test-unit
make test-integration
# пишем код
make check
git status
git add .
git commit -m "complete homework 1"
git push origin homework-1
```

Потом открыть Pull Request на GitHub.

## Что будет проверяться в Pull Request

После push GitHub Actions запустит несколько независимых проверок:

- зависимости и `go.mod`;
- форматирование;
- style-линтеры;
- static-analysis линтеры;
- security-линтеры;
- unit-тесты;
- integration-тесты;
- race-тесты;
- coverage gate;
- сборка бинарника.

Если хотя бы одна проверка красная, Pull Request нужно исправить и запушить новый commit.

## Что будет проверяться в Pull Request

После push GitHub Actions запустит несколько независимых проверок:

- зависимости и `go.mod`;
- форматирование;
- style-линтеры;
- static-analysis линтеры;
- security-линтеры;
- unit-тесты;
- integration-тесты;
- race-тесты;
- coverage gate;
- сборка бинарника.

Если хотя бы одна проверка красная, Pull Request нужно исправить и запушить новый commit.
