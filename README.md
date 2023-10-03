### In-memory хранилище с rest-api

(dockerfile в корне проекта, нужно запустить с параметром -p 8080:8080)

Примеры:
```
$ curl -i --header "Content-Type: application/json" \
--request POST \
--data '{"key":"amogus","value":"sus","seconds":1000}' \
http://localhost:8080/set

HTTP/1.1 200 OK
Date: Sat, 16 Sep 2023 09:45:25 GMT
Content-Length: 0
```
Для метода /set указывается ключ, значение и ttl в секундах

```
$ curl -i --header "Content-Type: application/json" \
--request GET \
--data '{"key":"amogus"}' \
http://localhost:8080/get

HTTP/1.1 200 OK
Date: Sat, 16 Sep 2023 09:51:34 GMT
Content-Length: 44
Content-Type: text/plain; charset=utf-8

{"key":"amogus","value":"sus","exists":true}
```
Метод /get возвращает ключ и его значение, а так же поле exists, говоряще о том, был ли такой ключ найден (т.к. можно хранить пустые ключи и значения)

```
$ curl -i --header "Content-Type: application/json" \
--request DELETE \
--data '{"key":"amogus"}' \
http://localhost:8080/delete

HTTP/1.1 200 OK
Date: Sat, 16 Sep 2023 09:51:42 GMT
Content-Length: 0 
```
Метод /delete удаляет ключ
После удаления ключ будет недоступен:
```
$ curl -i --header "Content-Type: application/json" \
--request GET \
--data '{"key":"amogus"}' \
http://localhost:8080/get

HTTP/1.1 200 OK
Date: Sat, 16 Sep 2023 09:51:46 GMT
Content-Length: 36
Content-Type: text/plain; charset=utf-8

{"key":"","value":"","exists":false}
```

### Обновление на 03.10.23
Я получил обратную связь по этому проекту, исправил некоторые ошибки и добавил новый функицонал
Ошибки:
- Неправильное использование интерфейсов - несоблюдение правила *"Accepting interfaces, returning structs"*:
    - интерфейс `Store` перенесен в пакет-консьюмер `controllers/http/v1`
    - `GetStore()` теперь возвращает структуру типа `*myStore`, а не интерфейс `Store`
- Неуместное использование синглтона для `Store`
    - теперь в пакете `controllers/http/v1` описана структура `myHttpHandler`, которая и имеет зависимость от объекта, соотвествующего интерфейсу `Store`

Нововведения:
- rate-limiter по ip адресам запросов с поддержкой асинхронности (использован `sync.RWMutex`) и соотвествующий middleware (см. `pkg/limiter`)

#### TODO
- добавить поддержку большего типа данных для хранящихся значений (реализация ims поддерживает любые типы данных, но по api можно пока передавать только строки)
