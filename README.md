# REST API Avito balance system

## Разобранны следующие операции:
- get:    /api/users/{id}?currency=                       - получение баланса пользователя(currency - опционально; например: USD, EUR, CNY)
- get:    /api/history/{id}                               - получение истории операций пользователя
- post:   /api/accrual?id=&amount=                        - начисление денег пользователю на счет
- post:   /api/write-downs?id=&amount=                    - списание денег с пользовательского счета
- post:   /api/transfer?sender_id=&receiver_id=&amount=   - списание денег между пользовательскими счетами

### Для запуска приложения:

```
make build && make run
```

Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
make migrate
```

Если требуется прекратить работу сервера или бд:
```
make app_down
make db_down
```