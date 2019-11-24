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

response: 
{
    "items": [
        {
            "id": "67829293-36d9-4113-8b6b-ce42c7c5bf45",
            "image": "<img src=\"images/16.jpg\" />",
            "name": "MacBook Pro2",
            "price": 2110,
            "description": "Test2",
            "in_stock": 10
        },
        {
            "id": "7131ca3d-5984-4f91-9c9c-dbe4998ba09f",
            "image": "<img src=\"images/16.jpg\" />",
            "name": "MacBook Pro2",
            "price": 2110,
            "description": "Test2",
            "in_stock": 10
        },
        {
            "id": "769d09ea-f062-4680-8535-45b3e6b23013",
            "image": "<img src=\"images/16.jpg\" />",
            "name": "MacBook Pro2",
            "price": 2110,
            "description": "Test2",
            "in_stock": 10
        },
        {
            "id": "791aec84-3eb3-4e5a-863d-81db8cdd29b5",
            "image": "<img src=\"images/16.jpg\" />",
            "name": "MacBook Pro2",
            "price": 2110,
            "description": "Test2",
            "in_stock": 10
        },
        {
            "id": "81260099-2608-4495-98bb-b032689872ee",
            "image": "<img src=\"images/16.jpg\" />",
            "name": "MacBook Pro2",
            "price": 2110,
            "description": "Test2",
            "in_stock": 10
        },
        {
            "id": "89eea3a4-4bc0-4fb3-a054-16eab1bd50a5",
            "image": "<img src=\"images/16.jpg\" />",
            "name": "MacBook Pro2",
            "price": 2110,
            "description": "Test2",
            "in_stock": 10
        },
        {
            "id": "8bf6b3c9-8377-4cc3-b620-f8020005880f",
            "image": "<img src=\"images/16.jpg\" />",
            "name": "MacBook Pro2",
            "price": 2110,
            "description": "Test2",
            "in_stock": 10
        },
        {
            "id": "a8ca6d1e-1692-4b3a-81a3-a82c16322ede",
            "image": "<img src=\"images/16.jpg\" />",
            "name": "MacBook Pro2",
            "price": 2110,
            "description": "Test2",
            "in_stock": 10
        },
        {
            "id": "b528a7c8-da81-4ebd-90ea-6c5da9bce637",
            "image": "<img src=\"images/16.jpg\" />",
            "name": "MacBook Pro2",
            "price": 2110,
            "description": "Test2",
            "in_stock": 10
        },
        {
            "id": "f37b3f6f-596d-413e-9822-6c572dde4ede",
            "image": "<img src=\"images/16.jpg\" />",
            "name": "MacBook Pro2",
            "price": 2110,
            "description": "Test2",
            "in_stock": 10
        },
        {
            "id": "f3f5d342-2b0d-4521-9638-292ed5d25993",
            "image": "<img src=\"images/16.jpg\" />",
            "name": "MacBook Pro2",
            "price": 2110,
            "description": "Test2",
            "in_stock": 10
        }
    ],
    "total_count": 11,
    "count_of_page": 11,
    "current_page": 1
}
```

####Get items with by pagination request
```go
[GET] localhost:8080/items?page=2&size=5 
-> page it's next page, size it's size items on request page

response:
{
    "items": [
        {
            "id": "2c5ef07c-713e-4654-8bbb-a2cac91fe11a",
            "image": "<img src=\"images/16.jpg\" />",
            "name": "MacBook Pro2",
            "price": 2110,
            "description": "Test2",
            "in_stock": 10
        }
    ],
    "total_count": 5,
    "count_of_page": 5,
    "current_page": 2
}
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
