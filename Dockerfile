FROM docker.io/guygrigsby/golang:1.11-alpine as builder

WORKDIR /$GOPATH/src/github.com/guygrigsby/cowsay-server
RUN apk update && apk --no-cache add ca-certificates dep git && update-ca-certificates
COPY . ./

RUN dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /cowsay-server .

FROM ubuntu:18.04

RUN apt-get update && apt-get install -y cowsay ca-certificates && update-ca-certificates
COPY --from=builder /cowsay-server ./
EXPOSE 80 443
ENTRYPOINT ["./cowsay-server"]

