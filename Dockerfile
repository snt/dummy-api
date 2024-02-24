FROM golang:1.22.0-alpine3.19 AS builder

COPY go.mod *.go /app/

WORKDIR /app

RUN go mod tidy

RUN go build -ldflags '-w -s' -a -installsuffix cgo -o dummy-api

FROM scratch

WORKDIR /app

COPY --from=builder /app/dummy-api .

EXPOSE 8080

CMD ["/app/dummy-api"]

