FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -tags netgo -trimpath -o go-proxy-newbing main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/go-proxy-newbing .

EXPOSE 8080
CMD ["/app/go-proxy-newbing"]