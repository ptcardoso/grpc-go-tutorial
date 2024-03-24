FROM golang:latest AS builder

LABEL maintainer="Pedro Cardoso <pmtcardoso@gmail.com>"

WORKDIR /
COPY . .

RUN go build -o main main.go

EXPOSE 50051

CMD ["./main"]