FROM golang:1.24-alpine AS build

WORKDIR /app

# Install necessary packages
COPY . ./

RUN go mod download
RUN go build -o /app/goshort .

FROM alpine:3.21.3 AS production

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/goshort .

# Run the binary
RUN chmod +x /app/goshort

# Copy the startup script
COPY ./startup/run.sh /app/run.sh

# Change permissions for the startup script
RUN chmod +x /app/run.sh

# Expose the port the app runs on
EXPOSE 3000

CMD ["/app/run.sh"]