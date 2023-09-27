FROM golang:1.20 AS build
WORKDIR /app

# Install required modules
COPY go.mod go.sum ./
RUN go mod download

# Build the service
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/panel-service ./cmd/panel-service

# Runtime environment (for minifying image size)
FROM gcr.io/distroless/static-debian12

COPY --from=build /bin/panel-service .
EXPOSE 9090
CMD ["./panel-service"]