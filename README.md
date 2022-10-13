# trancur

GET /courses/:source

Получение курсов из указанного в source источника. Пока только два: `rus` и `th`

Пример запроса:
`http://localhost:8080/courses/rus`

Пример ответа:
```
{
    "data": {
        "exchanges": [
            {
                "First": "RUB",
                "Second": "AUD",
                "Rate": 39.9813
            },
            {
                "First": "RUB",
                "Second": "AZN",
                "Rate": 37.5035
            },
        ],
        "updated": "2022-10-13T00:00:00Z"
    },
    "error": null
}
```
