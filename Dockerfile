FROM golang:1.24-alpine AS build

WORKDIR /app

# Install necessary packages
COPY . ./

RUN go mod download
RUN go build -o /app/main .

FROM alpine:3.21.3 AS production

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/main .

# Run the binary
RUN chmod +x /app/main

# Expose the port the app runs on
EXPOSE 3000

CMD ["/app/main"]