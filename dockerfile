#build
FROM golang:1.21 AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./IMS ./cmd/main.go


#run
FROM alpine:latest

WORKDIR /app
COPY --from=builder /build/IMS ./IMS
EXPOSE 8080
CMD ./IMS
