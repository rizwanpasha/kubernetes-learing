FROM golang:1.23.0-alpine
WORKDIR /app
COPY go.mod main.go ./
RUN go build -o main .
ENTRYPOINT ["./main"]

