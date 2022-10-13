# trancur

### Использование

GET /courses/:source

Получение курсов из указанного в source источника. Пока только два: `rus` и `th`. Если источник не указан, то используется по умолчанию, указанный в env-е.

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

POST /transit/:source

Перевод из одной валюты в другую. 
```
{
	"input": "foo",  --- входящая валюта в общепринятом сокращении
	"output": "bar", --- исходящая валюта в общепринятом сокращении
	"value": 99.99   --- количество входящей валюты
}
```

Пример запроса: 
`http://localhost:8080/transit/rus`

Пример ответа:
```
{
    "data": {
        "input": "RUB",
        "output": "USD",
        "result": 1.5684822894822283,
        "value": 100
    },
    "error": null
}
```


### Настройка

В корне проекта лежит файл `.env`. В нем следующие настройки:
- `SELF_HTTP_PORT` --- порт http-сервера;
- `DEFAULT_COURSE_SOURCE` --- источник курсов обмена;
- `HEARTBEAT_DURATION` --- период обновления курсов.

### Текущие нерешенные проблемы

1. Никакого логирование наружу, только чистый высер в консоль.
2. Источники обновляются раз в сутки (так как ЦБ вроде выставляют суточныфе курсы). Но точное время изменения у всех предположительно разные. Отсюда предположение, что просто всеми обновляться раз в 24 часа -- малость некорректно.
3. Формат ввода и вывода (сколько цифр после запятой?). Обычно такое решается на фронте.
4. Локализация текстовок об ошибках. И в логах.
5. Валидация входящих данных, особенно в переводе. Отсутствует.

