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
HEALTHCHECK --interval=5s --timeout=5s --start-period=5s --retries=10 CMD curl -f http://localhost:8000/api/v1/healthcheck || exit 1
CMD ["./openhealth-api"]
