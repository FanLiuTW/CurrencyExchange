# CurrencyExchange
A Simple Golang Server For Currency Exchange

Run locally
```
$ make install ENV=local
$ ./Asiayo
```

Access the OpenAPI
```
localhost:4510/swagger
```

Build the docker image
```
$ sudo docker build . -t asiayo
```

Run the docker image
```
$ sudo docker run --rm -d -p 4510:4510 asiayo:latest
```

Run the test and export to txt file
```
$ go test -v ./routes/... > out.txt
```