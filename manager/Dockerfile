FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o manager .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/manager .

ENTRYPOINT ["./manager"]