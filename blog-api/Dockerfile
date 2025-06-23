# 1. Use official Golang image
FROM golang:1.22-alpine

# 2. Set environment
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# 3. Set working directory
WORKDIR /app

# 4. Copy go mod and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# 5. Copy the entire project
COPY . .

# 6. Build the app (main.go is in cmd/blog-api)
RUN go build -o blog-api ./cmd/blog-api

# 7. Run the binary
CMD ["./blog-api"]
