# 💸 FlowMoney

FlowMoney — Telegram бот для отслеживания личных финансов. Полноценный микросервис на Go с REST API и Telegram ботом.

## Возможности

- Регистрация и авторизация
- Управление категориями (добавить, просмотреть все, найти по ID)
- Управление транзакциями (доходы и расходы)
- Управление балансом
- Бюджет по категориям и месяцам

## Стек

- **Go** — язык разработки
- **PostgreSQL** — база данных
- **sqlx** — работа с БД
- **net/http** — REST API
- **telebot v3** — Telegram бот
- **JWT** — авторизация
- **Docker / Docker Compose** — деплой

## Структура

```
flowmoney/
├── api/        # REST API
├── bot/        # Telegram бот
├── migrations/ # SQL миграции
├── docker-compose.yml
└── config.yml
```

## Запуск

**1. Создай `.env` файл в корне проекта:**

```env
BOT_TOKEN=your_telegram_bot_token
JWT_SECRET=your_jwt_secret
DB_PASS=your_db_password
DB_USER=your_db_user
DB_NAME=moneyflow
HTTPS_PROXY=your_proxy  # опционально
```

**2. Запусти через Docker Compose:**

```bash
docker compose up --build
```

API будет доступен на `http://localhost:8080`

## API эндпоинты

| Метод | Путь | Описание |
|-------|------|----------|
| POST | /auth/register | Регистрация |
| POST | /auth/login | Вход |
| GET | /user/:id | Профиль пользователя |
| PUT | /user/:id | Обновить баланс |
| GET | /category/:id | Категория по ID |
| GET | /category/user/:id | Все категории пользователя |
| POST | /category | Создать категорию |
| GET | /transaction/:id | Транзакция по ID |
| POST | /transaction | Создать транзакцию |
| GET | /budget/:id | Бюджет по ID |
| POST | /budget | Создать бюджет |