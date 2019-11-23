##Simple Rest Api server on Go.
1. Usage in memory db "BoltDb"
2. gorilla/mux for Handling endpoints functions

###For start server
```shell script
Run appleShop file or in command line execute go run main.go
```

###Rest Endpoints
####Get all items
```go
[GET] http://localhost:8080/items
```
####Create new item
```go
[POST] http://localhost:8080/items

body:
{
    "image": "<img src=\"images/16.jpg\" />",
    "name": "MacBook Pro2",
    "price": 2110,
    "description": "Test2",
    "in_stock": 10
}
```

####Update item
```go
[PUT] http://localhost:8080/items/{id}
body:
{
    "image": "<img src=\"images/16.jpg\" />",
    "name": "MacBook Pro2",
    "price": 2110,
    "description": "Test2",
    "in_stock": 10
}
```

####Delete item
```go
[DELETE] http://localhost:8080/items/{id}
```

###Generate unique id for any entity
```go
func (i *Item) GenerateUniqueId() {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	i.ID = string(out)
}
```
