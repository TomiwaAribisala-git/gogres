FROM golang:latest as builder

WORKDIR /app

COPY ./go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o stock-api


FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/stock-api /stock-api

EXPOSE 8080

ENTRYPOINT ["/stock-api"]