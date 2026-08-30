# TOKENCHECKER

Учебное веб-приложение на Go с регистрацией, авторизацией через сессии и сохранением персонального сообщения пользователя.

## Возможности

Приложение поддерживает:

- регистрацию пользователя;
- безопасное хранение пароля через bcrypt;
- вход по логину и паролю;
- создание серверной сессии;
- хранение session token в HttpOnly cookie;
- защиту личного кабинета;
- сохранение персонального сообщения в PostgreSQL;
- загрузку сообщения после повторного входа;
- выход с удалением активной сессии.

## Технологии

Backend:

- Go
- Gin
- GORM

Database:

- PostgreSQL

Frontend:

- HTML
- CSS
- JavaScript
- Fetch API

Безопасность:

- bcrypt для хеширования паролей;
- криптографически случайные session tokens;
- SHA-256 для хранения хеша токена сессии;
- HttpOnly cookie;
- SameSite=Lax;
- проверка авторизации через middleware.

## Архитектура

Проект разделён на несколько слоёв:

```text
Browser
   ↓
Handler
   ↓
Service
   ↓
Repository
   ↓
GORM
   ↓
PostgreSQL
```

### Handler

`internal/handlers`

Отвечает за HTTP:

- принимает запросы;
- читает JSON;
- возвращает HTTP-ответы;
- устанавливает cookie;
- выполняет middleware авторизации.

### Service

`internal/service`

Содержит бизнес-логику приложения:

- проверку данных;
- регистрацию;
- проверку пароля;
- создание сессий;
- проверку токенов;
- сохранение сообщений.

### Repository

`internal/repository`

Работает непосредственно с PostgreSQL через GORM:

- поиск пользователей;
- создание пользователей;
- создание и удаление сессий;
- сохранение и загрузка сообщений.

### Model

`internal/model`

Содержит структуры данных приложения:

- User;
- Session;
- Message;
- структуры HTTP-запросов и ответов.

## Структура проекта

```text
TOKENCHECKER/
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   ├── config/
│   │   └── config.go
│   │
│   ├── handlers/
│   │   ├── api_handler.go
│   │   ├── auth_middleware.go
│   │   ├── message_handler.go
│   │   └── page_handler.go
│   │
│   ├── model/
│   │   ├── auth.go
│   │   ├── message.go
│   │   ├── session.go
│   │   └── user.go
│   │
│   ├── repository/
│   │   ├── database.go
│   │   ├── message_repository.go
│   │   ├── migration.go
│   │   ├── session_repository.go
│   │   └── user_repository.go
│   │
│   └── service/
│       ├── auth_service.go
│       └── message_service.go
│
├── web/
│   ├── static/
│   │   └── app.js
│   │
│   └── templates/
│       ├── cabinet.html
│       ├── index.html
│       ├── login.html
│       └── register.html
│
├── .env
├── .env.example
├── .gitignore
├── README.md
├── go.mod
└── go.sum
```

## База данных

Приложение использует PostgreSQL.

Основные таблицы:

```text
users
sessions
messages
```

### users

Хранит пользователей.

Пароль в открытом виде не сохраняется.

Вместо него хранится bcrypt-хеш:

```text
password
↓
bcrypt
↓
password_hash
```

### sessions

Хранит активные пользовательские сессии.

Сам session token в базе не хранится.

```text
session token
↓
SHA-256
↓
token_hash
```

Настоящий token находится только в cookie браузера.

### messages

Хранит сообщение пользователя.

Сообщение связано с пользователем через:

```text
user_id
```

## Настройка окружения

Создайте `.env` в корне проекта.

Пример:

```env
APP_PORT=8080

DB_HOST=127.0.0.1
DB_PORT=5433
DB_USER=go_auth_app
DB_PASSWORD=your_password
DB_NAME=go_auth_message_app
DB_SSLMODE=disable
```

Файл `.env` содержит секреты и не должен попадать в Git.

Для примера настроек используется:

```text
.env.example
```

## Запуск

Установите зависимости:

```bash
go mod tidy
```

Запустите PostgreSQL.

Затем из корня проекта:

```bash
go run ./cmd/app
```

После запуска приложение доступно по адресу:

```text
http://localhost:8080
```

## Основные маршруты

Страницы:

```text
GET /
GET /register
GET /login
GET /cabinet
```

API:

```text
POST /api/register
POST /api/login
POST /api/logout

GET  /api/message
POST /api/message
```

## Сценарий работы

Регистрация:

```text
HTML form
↓
JavaScript fetch
↓
POST /api/register
↓
AuthService
↓
bcrypt
↓
UserRepository
↓
PostgreSQL
```

Вход:

```text
POST /api/login
↓
поиск пользователя
↓
bcrypt.CompareHashAndPassword
↓
генерация session token
↓
сохранение token hash в PostgreSQL
↓
HttpOnly cookie
```

Доступ к кабинету:

```text
GET /cabinet
↓
Auth middleware
↓
cookie
↓
SHA-256
↓
поиск session
↓
проверка срока действия
↓
UserID
↓
cabinet
```

Сообщение:

```text
UserID
↓
POST /api/message
↓
MessageService
↓
MessageRepository
↓
PostgreSQL
```

После повторного входа:

```text
GET /api/message
↓
поиск сообщения по UserID
↓
текст возвращается в кабинет
```

Выход:

```text
POST /api/logout
↓
удаление session из PostgreSQL
↓
удаление cookie
↓
пользователь больше не имеет доступа к /cabinet
```

## Разработка

Перед проверкой изменений рекомендуется выполнять:

```bash
go fmt ./...
go mod tidy
go run ./cmd/app
```

## Важно

Текущая конфигурация предназначена прежде всего для локальной разработки.

Для production необходимо дополнительно использовать HTTPS и устанавливать для session cookie параметр:

```text
Secure=true
```