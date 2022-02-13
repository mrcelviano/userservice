# userservice
SocialTech

Сервис отвечает за работу с пользователями

Сервис слушает порт 8080

### HTTP

##### POST /api/
Создать пользователя

Request body
```
{
    "name": "testName",
    "email": "test@gmail.com"
}
```

Response body
```
{
    "id": 1,
    "email": "test@gmail.com",
    "name": "testName"
}
```

##### DELETE /api/:id/
Удалить пользователя

Response: NoContent Status OK

##### GET /api/1/
Получить пользователя по id

Response body
```
{
    "id": 1,
    "email": "test@gmail.com",
    "name": "testName"
}
```

##### GET /api/
Получить список пользователей

Query parameters:
```
    limit
    offset
    sortKey
    sortOrder
```

Response body
```
{
    "total": 1,
    "result": [
        {
            "id": 10,
            "email": "test@gmail.com",
            "name": "testUpdate"
        }
    ]
}
```

##### PUT /api/1/
Обновить пользователя

Request body
```
{
    "name": "testUpdate",
    "email": "test@gmail.com"
}
```

Response body
```
{
    "id": 1,
    "email": "test@gmail.com",
    "name": "testUpdate"
}
```
