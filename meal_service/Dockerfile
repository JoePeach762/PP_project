FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/internal/pb/swagger ./internal/pb/swagger

RUN adduser -D -s /bin/sh appuser && \
    chown -R appuser:appuser /app

USER appuser
CMD ["./app"]