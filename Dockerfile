FROM golang:1.24.0-alpine3.20 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .
COPY internal/ internal/

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

###############################################
FROM alpine:3.20

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
