# API service for sending notifications

### Запуск через Docker-compose
1. Склонировать репозиторий.
2. Изменить переменные в `configs/config.toml` файле
3. По желанию можно изменить токен аутентификации для отправки сообщений на сервис в ```sweater/__init__.py```
4. В консоли перейти в папку с репозиторием и выполнить команду ```docker-compose --env-file configs/config.toml up``` (предварительно запустив docker).
5. После завершения сборки образов и запуска контейнеров адрес http://localhost:8080/ (по умолчанию) начнет принимать запросы и тг бот начнет работу.

### Адреса запросов и примерами структуры запроса к ним на python
* http://localhost:8080/api/getJWT
    * Отвечает информацией с jwt для последующей аутентификации.
    ```Python
    import requests
    requests.get("http://localhost:8080/api/getJWT")
    ```

* http://localhost:8080/api/contact - взаимодействует с таблицей contact
    * /get - Отвечает информацией со всеми клиентами из базы данных.
    ```Python
    import requests
    requests.get("http://localhost:8080/api/contact/get", headers={"Token": "some token"})
    ```

    * /post - Отвечает информацией с id созданной записи. Принимает json.
    ```Python
    import requests
    requests.post("http://localhost:8080/api/contact/post", headers={"Token": "some token"}, json={"number":"+7917235678", "operator_code":"917", "tag":"tag1", "time_zone":"+2"})
    ```

    * /put - Отвечает измененной записью клиента. Принимает json с id записи которую нужно изменить и названиями столбов с новой информацией. (Можно отправлять только столбцы с новой информацией, столбцы с данными которые изменяться не должны допускается не отправлять)
    ```Python
    import requests
    requests.put("http://localhost:8080/api/contact/put", headers={"Token": "some token"}, json={"change": "2", "tag": "tag1", "time_zone": "+6", "number":"+791723534"})
    ```

    * /delete - Отвечает статусом. Принимает id записи которую нужно удалить.
    ```Python
    import requests
    requests.delete("http://localhost:8080/api/contact/delete", headers={"Token": "some token"}, json={"id": 2})
    ```


### Для следющих адрессов запросы будут аналогичны примерам выше
* http://localhost:8080/api/mailing/get
* http://localhost:8080/api/mailing/post
* http://localhost:8080/api/mailing/put
* http://localhost:8080/api/mailing/delete

* http://localhost:8080/api/message/get
* http://localhost:8080/api/message/post
* http://localhost:8080/api/message/put
* http://localhost:8080/api/message/delete

#### http://localhost:8080/api/stat/get/{id} - Адресс принимает только GET запрос. Отвечает детальной информацией о рассылке с отправленными по ней сообщениями. Принимает id рассылки в своем адресе.
* ```Python
  import requests
  requests.get("http://localhost:8080/api/mailing/{id of mailing}", headers={"Token": "some token"})
  ```


### Языки и технологии
![Golang](https://img.shields.io/badge/-Golang-090909?style=for-the-badge&logo=go)
![Python](https://img.shields.io/badge/-Python-090909?style=for-the-badge&logo=python)
![Framework](https://img.shields.io/badge/-Gin_Framework-090909?style=for-the-badge)
![GORM](https://img.shields.io/badge/-GORM-090909?style=for-the-badge&)
![JWT](https://img.shields.io/badge/-jwt_authontification-090909?style=for-the-badge)
![Docker](https://img.shields.io/badge/-Docker-090909?style=for-the-badge&logo=Docker)
![Postgres](https://img.shields.io/badge/-Postgres-090909?style=for-the-badge&logo=Postgresql)
![Redis](https://img.shields.io/badge/-Redis-090909?style=for-the-badge&logo=Redis)
![RabbitMQ](https://img.shields.io/badge/-RabbitMQ-090909?style=for-the-badge&logo=RabbitMQ)