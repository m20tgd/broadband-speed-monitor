# Stage 1: Build the Go binary
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY src ./src
RUN go build -o broadband-speed-monitor ./src

# Stage 2: Create a minimal runtime image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/broadband-speed-monitor .
COPY .env .env
RUN chmod +x ./broadband-speed-monitor
CMD ["./broadband-speed-monitor"]