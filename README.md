##Simple Rest Api server on Go.
1. Usage in memory db "BoltDb"
2. gorilla/mux for Handling endpoints functions

###For start server
```shell script
1. Run appleShop file or in command line execute go run main.go
2. Build startup file for your system, for example windows:
sudo GOOS=<os> GOARCH=<architecture_type> go build -o <nameOfApp>.<extenstion> <projectEntryPointFile>.go
sudo GOOS=windows GOARCH=386 go build -o appleShop.exe appleShop/main.go 
```

###Rest Endpoints
####Get all items
```go
[GET] http://localhost:8080/items
```

####Get items with by pagination request
```go
[GET] localhost:8080/items?page=2&size=5 
-> page it's next page, size it's size items on request page
```

####Get item by id
```go
[GET] http://localhost:8080/{id}
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
	newID := uuid.New()
	item, _ := GetItem(newID.String())
	if item == nil {
		i.ID = newID.String()
	} else {
		i.GenerateUniqueId()
	}
}
```
