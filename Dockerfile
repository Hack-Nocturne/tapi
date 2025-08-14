# Stage 1: Build the Go binary with sqlc
FROM golang:1.24-alpine AS builder

WORKDIR /

# Cache Go dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN rm -rf secrets

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o test

# Stage 2: Minimal runtime image
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /test /test

USER nonroot
ENTRYPOINT ["/test"]
