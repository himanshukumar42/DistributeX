FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main ./cmd/main.go
RUN ls -l /app  # Check if main is built successfully

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/Schema.sql .

RUN chmod +x main

EXPOSE 8080

CMD ["./main"]

