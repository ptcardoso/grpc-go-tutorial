FROM golang:latest AS builder

LABEL maintainer="Pedro Cardoso <pmtcardoso@gmail.com>"

WORKDIR /
COPY . .

RUN go build -o main main.go

EXPOSE 8080

CMD ["./main"]