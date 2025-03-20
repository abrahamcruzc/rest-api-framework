FROM golang:1.23-alpine

WORKDIR /app
COPY . .

RUN go build -o sicei .

EXPOSE 8080
CMD ["./sicei"]
