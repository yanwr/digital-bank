FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

FROM scratch
WORKDIR /app
EXPOSE 8080
COPY --from=builder /app/main .
ENTRYPOINT [ "/app/main" ]