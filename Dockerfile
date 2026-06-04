FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o labelin cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/labelin .
COPY --from=builder /app/views ./views
COPY --from=builder /app/public ./public

EXPOSE 3000
CMD ["./labelin"]
