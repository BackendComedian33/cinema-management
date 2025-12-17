FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./ticketing ./cmd

FROM alpine:edge AS final

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/ticketing .

COPY --from=builder /app/.env .

EXPOSE 3300

CMD ["./ticketing"]
