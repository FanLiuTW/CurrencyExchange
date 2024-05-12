# build stage
FROM golang:1.17 AS builder
WORKDIR /go/src/Asiayo
COPY . .
RUN make install ENV=local

# final stage
FROM ubuntu:20.04
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /opt/zoneinfo.zip
ENV ZONEINFO /opt/zoneinfo.zip
WORKDIR /root/
COPY --from=builder /go/src/Asiayo/Asiayo .
EXPOSE 4510
ENTRYPOINT ["./Asiayo"]
