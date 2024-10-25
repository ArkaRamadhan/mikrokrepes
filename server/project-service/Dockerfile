# Dockerfile untuk project-service
FROM golang:1.23.2

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o project-service

CMD ["./project-service"]

