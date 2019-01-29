FROM golang:1.10 AS builder

WORKDIR /$GOPATH/src/github.com/guygrigsby/cowsay-server

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /cowsay-server .

FROM ubuntu:18.04

RUN apt-get update && apt-get install -y cowsay
COPY --from=builder /cowsay-server ./
EXPOSE 8080
ENTRYPOINT ["./cowsay-server"]

