FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM scratch
WORKDIR /app
EXPOSE 8080
COPY --from=builder /app/server /server
ENTRYPOINT [ "/server" ]