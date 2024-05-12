# CurrencyExchange
A Simple Golang Server For Currency Exchange

Run locally
```
$ make install ENV=local
$ ./Asiayo
```

Access the OpenAPI
```
localhost:12345/swagger
```

Build the docker image
```
$ sudo docker build . -t asiayo
```

Run the docker image
```
$ sudo docker run asiayo:latest
```

Run the test and export to txt file
```
$ go test -v ./routes/... > out.txt
```