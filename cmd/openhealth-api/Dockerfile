FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd/openhealth-api
RUN CGO_ENABLED=0 GOOS=linux go build -o /openhealth-api .

FROM alpine:latest
RUN apk --no-cache add ca-certificates curl
WORKDIR /root/
COPY --from=builder /openhealth-api .
# COPY --from=builder /app/_migrations ./_migrations
CMD ["./openhealth-api"]
