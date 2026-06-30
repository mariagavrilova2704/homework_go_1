# Как сдавать домашнюю работу через GitHub

Эта инструкция описывает маршрут от fork до Pull Request.

## Главная идея

Вы не отправляете домашку архивом в чат. Вы делаете изменения в своём fork на GitHub и открываете Pull Request.

Маршрут выглядит так:

```text
fork → clone → branch → code → test → commit → push → pull request → review
```

## Один раз перед первой домашкой

### 1. Создайте аккаунт GitHub

Перейдите на GitHub и зарегистрируйтесь.

### 2. Установите Git

Проверьте установку:

```bash
git --version
```

Если команда показывает версию Git — всё хорошо.

### 3. Настройте имя и email для commit-ов

```bash
git config --global user.name "Your Name"
git config --global user.email "your_email@example.com"
```

### 4. Настройте SSH-ключ

Создайте ключ:

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```

Можно нажимать Enter на вопросы терминала и оставить путь по умолчанию.

Показать публичный ключ:

```bash
cat ~/.ssh/id_ed25519.pub
```

Скопируйте всю строку, которая начинается с `ssh-ed25519`.

На GitHub откройте:

```text
Settings → SSH and GPG keys → New SSH key
```

Вставьте ключ в поле `Key` и нажмите `Add SSH key`.

Проверьте подключение:

```bash
ssh -T git@github.com
```

Нормальный результат:

```text
Hi YOUR_USERNAME! You've successfully authenticated, but GitHub does not provide shell access.
```

## Для каждой домашки

### 1. Fork

На странице репозитория домашки нажмите `Fork`.

Fork создаёт вашу копию репозитория в вашем GitHub-аккаунте.

### 2. Clone

На странице вашего fork нажмите:

```text
Code → SSH → copy
```

Затем в терминале выполните:

```bash
git clone git@github.com:YOUR_USERNAME/homework1.git
cd homework1
```

Важно: после `git clone` терминал может напечатать несколько строк про cloning, objects, resolving deltas. Это нормальный технический вывод. Ничего дополнительно вводить не нужно.

### 3. Branch

```bash
git checkout -b homework-1
```

Ветка нужна, чтобы делать домашку отдельно от `main`.

### 4. Code

Откройте файл:

```text
internal/homework/homework.go
```

Реализуйте функции.

### 5. Test

```bash
make test-unit
make test-integration
make check
```

Если команда красная, сначала исправьте ошибку локально.

### 6. Commit

```bash
git status
git add .
git commit -m "complete homework 1"
```

### 7. Push

```bash
git push origin homework-1
```

### 8. Pull Request

На GitHub появится кнопка `Compare & pull request`.

Проверьте:

```text
base: main
compare: homework-1
```

Добавьте понятное название:

```text
Homework 1
```

Нажмите `Create pull request`.

### 9. CI

После открытия Pull Request GitHub запустит проверки.

Зелёные галочки означают, что всё прошло.

Красный крестик означает, что нужно открыть Details, прочитать ошибку, исправить код и снова сделать:

```bash
git add .
git commit -m "fix homework 1"
git push origin homework-1
```

Новый push автоматически обновит Pull Request и перезапустит CI.

## Что не делать

- Не отправлять `.zip` вместо Pull Request.
- Не менять тесты без согласования с преподавателем.
- Не писать решение прямо в `main`.
- Не коммитить папку `bin`.
- Не вставлять в терминал подсказки с сайта GitHub после `git clone`, если вы уже скопировали SSH-ссылку и выполнили `git clone`.

## Мини-чеклист перед отправкой

```text
[ ] Я нахожусь в ветке homework-1
[ ] Код написан в internal/homework/homework.go
[ ] make test-unit прошёл
[ ] make test-integration прошёл
[ ] make check прошёл
[ ] commit создан
[ ] push сделан
[ ] Pull Request открыт
[ ] CI зелёный
[ ] Преподаватель добавлен на review
```
