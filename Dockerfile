FROM golang:1.16.3-alpine3.13 AS builder

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -o stock-api .


FROM apline:latest AS production

COPY --from=builder /app .

EXPOSE 8080

CMD ["./stock-api"]