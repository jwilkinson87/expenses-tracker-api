FROM golang:latest as build

# Set environment variables for static binary
ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

RUN go build -o /api api/main.go
RUN go build -o /migration migration/main.go
 
FROM alpine:latest as run

# Copy the application executable from the build image
COPY --from=build /api /api
COPY --from=build /migration /migration
COPY --from=build /app/sql /app/sql

WORKDIR /app
EXPOSE 8080
CMD ["/api"]