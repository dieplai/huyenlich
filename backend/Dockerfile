FROM golang:1.23-alpine AS builder
WORKDIR /app/backend
COPY backend/go.mod ./
RUN go mod download
COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:3.20
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/backend/server .
COPY data /app/data
ENV DATA_DIR=/app/data
EXPOSE 8080
CMD ["./server"]
