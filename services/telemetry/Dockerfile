FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/services/telemetry

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o telemetry .

FROM alpine:3.22

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/services/telemetry/telemetry .

EXPOSE 8080

CMD ["./telemetry"]