FROM golang:1.23.0-alpine
WORKDIR /app

# When using COPY with more than one source file, 
# the destination must be a directory and end with a / or a \
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 12345

# this is a binary file run directly, CMD ["go", "run" "./main"] will throw error
ENTRYPOINT ["./main"]