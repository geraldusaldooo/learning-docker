FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Copy go.mod and execute to download Go modules
COPY go.mod go.sum ./
RUN go mod download


COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go_compose_service


EXPOSE 77

# Run
CMD [ "/go_compose_service" ]