# simple-rest-service

Simple REST service is a sample project

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=fdaines_simple-rest-service&metric=alert_status)](https://sonarcloud.io/dashboard?id=fdaines_simple-rest-service)


## Requirements
- Go 1.15+

for check if Go is intalled, run: `go version`

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
To change the service version, edit the `sonar.projectVersion` property from `.sonarcloud.properties` file

## Architecture Tests
```
$ go get -u github.com/fdaines/arch-go
$ arch-go
```


