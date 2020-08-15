FROM golang:alpine as builder
COPY ./ /app
WORKDIR /app
RUN  CGO_ENABLED=0 GOOS=linux go build -o backoffice-api
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app .
CMD ["./backoffice-api"]
