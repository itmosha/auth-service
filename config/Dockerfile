# Builder
FROM golang:1.22.1-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o binaryapp cmd/service/main.go

# Runner
FROM alpine AS runner
WORKDIR /app

COPY ./.env /app/.env
COPY ./config/config.yaml /app/config/config.yaml
COPY --from=builder /app/binaryapp .

CMD ["./binaryapp"]