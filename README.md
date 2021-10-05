# REST API Avito balance system

## Разобранны следующие операции:
- get/users/{id} - получение баланса пользователя
- get/history/{id}
- post/accrual/{id, sum}
- post/writedowns/{id, sum}
- post/transfer/{sender, reciver, sum} 

### Для запуска приложения:

```
make build && make run
```

Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
make migrate
```