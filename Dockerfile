# Use the official Golang image as the build stage
FROM golang:1.24-alpine as builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o app

# Use a minimal base image for the final build
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
