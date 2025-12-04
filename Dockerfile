FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

FROM alpine:3.20

RUN apk add --no-cache ca-certificates bash coreutils && update-ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 3400 4400

CMD ["./main"]
