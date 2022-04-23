# Legoha - Go Simple and Minimal Todolist API

Legoha is a simple and minimal Todolist api using Golang with Gorm and Fiber. Followed by Clean Architecture.

## How to run ?

Firstly, create a config.env file to hold the environment variables by:
```makefile
make 
```
Secondly, edit your DB_USER and DB_PASS for the username and password of the database config. Also install necessary go pakags by:

``` go
go mod tidy
```

Finally, run the application by:
```makefile
make run
```

## References
https://eminetto.medium.com/clean-architecture-using-golang-b63587aa5e3f
https://github.com/percybolmer/ddd-go/tree/clean-architecture
https://github.com/ruslantsyganok/clean_arcitecture_golang_example
https://github.com/vinigracindo/fiber-gorm-clean-architecture
https://github.com/gofiber/fiber
https://gorm.io