# CurrencyExchange
A Simple Golang Server For Currency Exchange

Run locally
```
$ make install ENV=local
$ ./BusServer
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