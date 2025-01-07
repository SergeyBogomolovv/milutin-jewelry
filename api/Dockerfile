FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /main cmd/main.go

FROM alpine:latest

RUN apk add --no-cache libc6-compat

COPY --from=builder /main /main

CMD ["/main"]