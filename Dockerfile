FROM golang:1.18-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build main.go

# Export necessary port.
EXPOSE 5000

# Command to run when starting the container.
ENTRYPOINT ["./main"]
