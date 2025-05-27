# GoShort - A URL Shortening Service

# Overview
GoShort is a URL shortening service written in Go. It provides a simple API to create and manage short URLs.

# Features
- Create short URLs
- Redirect to original URLs
- Custom short URLs
- Integration with Docker for easy deployment

# Installation
## Prerequisites
- Go 1.18 or later
- Docker (optional, for containerized deployment)

## Clone the repository
```bash
git clone git@github.com:InspectorGadget/goshort.git
cd goshort
```

## Build the application (Local)
```bash
go mod download
go build -o goshort ./cmd/goshort
```

## Run the application (Local)
```bash
./goshort
```

OR

## Run with Docker Compose (Docker)
```bash
docker-compose up -d
```
