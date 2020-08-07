FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main .

# Build a small image
FROM scratch

COPY --from=builder /build/main /

# Command to run
ENTRYPOINT ["/main"]
