# json-storage-api
JSON strage api with Golang

# Execute
```
$ go run main.go db.go
```

# test
## create stock
```
$ curl -X POST -H "Content-Type: application/json" localhost:8090/api/v1/stocks
```

### response
```json
{"error":null,"result":{"ID":4,"UUID":"1a04e33f41988d98fda0dd94d7e61a82","JSON":"{}"}}
```

## get stock
```
$ curl localhost:8090/api/v1/stocks/1a04e33f41988d98fda0dd94d7e61a82
```

### response
```json
{"error":null,"result":{"ID":4,"UUID":"1a04e33f41988d98fda0dd94d7e61a82","JSON":"{}"}}
```

## update stock
```
$ curl -X POST -H "Content-Type: application/x-www-form-urlencoded" --data-urlencode "json=\"{\"id\":\"test\",\"post\":\"test2\"}\"" localhost:8090/api/v1/stocks/1a04e33f41988d98fda0dd94d7e61a82/put
```

### response
```json
{"error":null,"result":true}
```

## Deltete stock
```
$ curl -X POST -H "Content-Type: application/x-www-form-urlencoded" localhost:8090/api/v1/stocks/1a04e33f41988d98fda0dd94d7e61a82/delete
```

### response
```json
{"error":null,"result":true}
```

# build
## In executing server
`go env`

## Do it
`env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build`

# deploy
# Do it
`nohup ./json-storage-api &`

## kill
`ps aux | grep json-storage-api`