# REST API Avito balance system

## Разобранны следующие операции:
- GET:    /api/users/{id}?currency=                       - получение баланса пользователя (currency - опционально; например: USD, EUR, CNY)
- GET:    /api/history/{id}?sort=&type=&limit=&offset=    - получение истории операций пользователя (Параметры опциональны. sort=sum/date, type=inc/dec, показываются элементы с offset до offset+limit. Default values = date, inc, все элементы без отступа)
- GET:    /api/history/count/{id}                         - получение колличества записей в таблице истори по id 
- POST:   /api/accrual?id=&amount=                        - начисление денег пользователю на счет
- POST:   /api/write-downs?id=&amount=                    - списание денег с пользовательского счета
- POST:   /api/transfer?sender_id=&receiver_id=&amount=   - списание денег между пользовательскими счетами

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

### To do:
1. Unit-tests
2. Индексирование таблицы истории по сумме, индексирование + партицирование таблицы истории по дате.
3. Улучшенное логирование. Можно писать не только ошибки но и успешные операции.