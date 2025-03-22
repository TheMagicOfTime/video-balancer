FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/video-balancer ./cmd/video-balancer

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/video-balancer .

EXPOSE 50051

CMD ["./video-balancer"]