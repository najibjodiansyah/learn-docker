# Howto

## Run the program using environment file

```bash
source template.env
go run main.go
```

## Run the program by using manually set environment

```bash
export DB_CONNECTION_STRING='username:password@tcp(127.0.0.1:3306)/database_name?charset=utf8&parseTime=True&loc=Local'
go run main.go
```
# learn-docker
