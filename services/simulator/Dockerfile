FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/services/simulator

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o simulator .

FROM alpine:3.22

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/services/simulator/simulator .

CMD ["./simulator"]