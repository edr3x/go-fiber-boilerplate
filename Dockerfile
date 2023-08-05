FROM golang:1.20.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
