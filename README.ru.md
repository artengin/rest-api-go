# REST API на Go для управления моделью Person

## Описание

Этот проект — пример реализации RESTful CRUD сервиса на языке Go с использованием архитектуры, вдохновлённой принципами Clean Architecture.  
Сервис позволяет создавать, получать, обновлять и удалять записи о людях (Person) в базе данных PostgreSQL.

## Стек технологий

- Язык: Go
- HTTP-фреймворк: [Echo](https://echo.labstack.com/)
- Работа с БД: [gocraft/dbr](https://github.com/gocraft/dbr)
- Логирование: [logrus](https://github.com/sirupsen/logrus)
- Валидация: [go-playground/validator](https://github.com/go-playground/validator)
- База данных: PostgreSQL

## Архитектура

Проект разделён на 4 слоя:

- **internal/app** — описание модели Person
- **internal/http** — HTTP-обработчики (handlers)
- **internal/logic** — бизнес-логика (use cases)
- **internal/repository/postgres** — слой доступа к данным (PostgreSQL)

## Модель Person

```go
type Person struct {
    ID        int64     `json:"id"`
    Email     string    `json:"email"`
    Phone     string    `json:"phone"`
    FirstName string    `json:"firstName"`
    LastName  string    `json:"lastName"`
    UpdatedAt time.Time `json:"updatedAt"`
    CreatedAt time.Time `json:"createdAt"`
}
```

## Запуск проекта

1. **Клонируйте репозиторий:**

2. **Создайте файл `.env`** в корне проекта

3. **Запустите PostgreSQL и создайте базу данных.**

4. **Запустите приложение:**
   ```bash
   go run .
   ```

   При первом запуске автоматически применится миграция для таблицы `person`.

## API

### Получить список Person

```
GET /person
```

или

```
GET /person?limit=10&offset=0&search=Иван
```

### Получить одного Person

```
GET /person/{id}
```

### Создать Person

```
POST /person
Content-Type: application/json

{
  "email": "test@example.com",
  "phone": "+79991234567",
  "firstName": "Иван",
  "lastName": "Иванов"
}
```

### Обновить Person

```
PUT /person/{id}
Content-Type: application/json

{
  "email": "test2@example.com",
  "phone": "+79991234568",
  "firstName": "Пётр",
  "lastName": "Петров"
}
```

### Удалить Person

```
DELETE /person/{id}
```