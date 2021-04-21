# simple-rest-service

Simple REST service is a sample project

## Requirements
- Go 1.15+

## Running the service
```
$ go run main.go
```

To check if the service is running:
```
$ curl http://localhost:10000/promotions
```


## Unit Tests
```
$ go test ./...
$ go test -cover ./...
$ go test ./... -coverprofile=c.out && go tool cover -html=c.out
```

## Mutation Tests
```
$ go get -v github.com/zimmski/go-mutesting
$ go-mutesting ./...
```

## SonarQube Static Analysis
Dashboard: https://sonarcloud.io/dashboard?id=fdaines_simple-rest-service

## Architecture Tests
```
$ go get -u github.com/fdaines/arch-go
$ arch-go
```


