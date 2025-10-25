## Структура проекта

├── cmd/server/ # Точка входа приложения
│ └── main.go # Инициализация: конфиг, БД, сервисы, запуск HTTP-сервера
│
├── internal/
│ ├── config/ # Загрузка и валидация конфигурации из .env
│ │
│ ├── domain/ # Ядро бизнес-логики
│ │ ├── models.go # Чистые структуры: User, Asset, Order и т.д.
│ │ └── repositories/ # Интерфейсы для работы с данными (без реализации!)
│ │
│ ├── storage/postgres/ # Реализация репозиториев через PostgreSQL
│ │ ├── db.go # Инициализация подключения к БД
│ │ ├── user.go # Реализация UserRepository
│ │ ├── asset.go # Реализация AssetRepository
│ │ └── order.go # Реализация OrderRepository
│ │
│ ├── services/ # Бизнес-логика: создание ассета, покупка, модерация и т.д.
│ │
│ └── delivery/http/ # HTTP-слой
│ ├── server.go # Запуск HTTP-сервера
│ ├── middleware/ # CORS, логгирование, аутентификация
│ ├── handlers/ # Обработчики запросов (вызывают сервисы)
│ └── routes.go # Регистрация маршрутов через chi.Router
│
├── migrations/ # SQL-миграции для goose
│ ├── 000001_create_users_table.sql
│ └── ...
│
├── pkg/ # Общие утилиты (если появятся)
│ └── uploader/ # Например, клиент для MinIO/S3
│
├── .env.example # Пример переменных окружения
├── go.mod
├── go.sum
└── README.md
