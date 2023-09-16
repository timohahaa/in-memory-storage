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

#### TODO
- добавить рейт-лимитер по ip (пока что есть глобальный лимитер по запросам к серверу)
- добавить поддержку большего типа данных для хранящихся значений (реализация ims поддерживает любые типы данных, но по api можно пока передавать только строки)
