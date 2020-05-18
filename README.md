# README #
Simple REST API for Users without authorization on GoLang

## BUILD ##

Before run, please install next modules:

```
go get github.com/jinzhu/gorm
go get github.com/joho/godotenv
go get github.com/gorilla/mux
go get github.com/lib/pq
```

## CONTROLLERS ##
**userController**

*GET ALL USERS*
REQUEST GET:
http://localhost:8082/users
RESPONSE:

```JSON
[{"ID":6,"CreatedAt":"2020-05-18T14:51:18.219384+06:00","UpdatedAt":"2020-05-18T14:51:18.219384+06:00","DeletedAt":null,"name":"TEST_USER","email":"test@gmail.com","birthday":"1998-09-01"}]
```

*CREATE USER*
REQUEST POST:
http://localhost:8082/users
BODY:
```JSON
{"name":"TEST_USER", "email":"test@gmail.com","birthday":"1999-09-01"}
```

RESPONSE:
```JSON
{"id":1, "name":"TEST_USER", "email":"test@gmail.com","birthday":"1999-09-01",CreatedAt":"2020-05-18T17:40:59.5283223+06:00","UpdatedAt":"2020-05-18T17:40:59.5283223+06:00"}

```


*UPDATE USER*
REQUEST POST:
http://localhost:8082/users/{id}
BODY:
```JSON
{"id":1,"name":"TEST_USER_NAME", "email":"test@gmail.com","birthday":"1999-09-01"}
```

RESPONSE:
```JSON
{"id":1, "name":"TEST_USER_NAME", "email":"test@gmail.com","birthday":"1999-09-01",CreatedAt":"2020-05-18T17:40:59.5283223+06:00","UpdatedAt":"2020-05-19T17:40:59.5283223+06:00"}
```

*DELETE USER*
REQUEST DELETE:
http://localhost:8082/users/{id}

RESPONSE:
```
true
```
