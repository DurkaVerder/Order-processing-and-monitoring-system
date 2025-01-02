FROM golang:1.23.3

WORKDIR /order-status-changer

COPY . .

RUN go mod tidy

RUN go build -o app ./cmd/app/main.go

CMD ["./app"]
