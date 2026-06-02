FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/services/dashboard

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o dashboard .

FROM alpine:3.22

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/services/dashboard/dashboard .

EXPOSE 8082

CMD ["./dashboard"]