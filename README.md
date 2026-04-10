# Что изменено

Создание графика задач реализовано как отдельные API запросы.

Можно создавать несколько расписаний на задачу. Так же есть возможность гибко настроить одно расписание.

Минусы:
- Если нужно сразу создать задачу с расписанием, то придется делать два запроса - к задаче и к графику

# Task Service

Сервис для управления задачами с HTTP API на Go.

## Требования

- Go `1.23+`
- Docker и Docker Compose

## Быстрый запуск через Docker Compose

```bash
docker compose up --build
```

После запуска сервис будет доступен по адресу http://localhost:8080

Swagger UI для тестирования http://localhost:8080/swagger

Если `postgres` уже запускался ранее со старой схемой, пересоздай volume:

```bash
docker compose down -v
docker compose up --build
```

Причина в том, что SQL-файлы из `migrations` монтируются в `docker-entrypoint-initdb.d` и применяются только при инициализации пустого data volume.

## API

Базовый префикс API:

```text
/api/v1
```

Основные маршруты:

- `GET /api/v1/tasks`
- `POST /api/v1/tasks`
- `GET /api/v1/tasks/{id}`
- `PUT /api/v1/tasks/{id}`
- `DELETE /api/v1/tasks/{id}`

Вторичные маршруты:

- `GET /api/v1/tasks/{task_id}/schedule`
- `POST /api/v1/tasks/{task_id}/schedule`
- `GET /api/v1/tasks/{task_id}/schedule/{schedule_id}`
- `PUT /api/v1/tasks/{task_id}/schedule/{schedule_id}`
- `DELETE /api/v1/tasks/{task_id}/schedule/{schedule_id}`
