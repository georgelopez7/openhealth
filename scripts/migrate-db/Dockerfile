FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/scripts/migrate-db
RUN CGO_ENABLED=0 GOOS=linux go build -o /migrate-db .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /migrate-db .
COPY --from=builder /app/_migrations ./_migrations
CMD ["./migrate-db"]
