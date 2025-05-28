# GoShort - A URL Shortening Service

# Overview
GoShort is a URL shortening service written in Go, supported by Docker Compose. It provides a simple API to create and manage short URLs.

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

## Build the application
```bash
go build -o bin/goshort .
```

## Run Database Migrations (Locally)
```bash
bin/goshort migrate
```

## Run the Application (Locally)
```bash
bin/goshort start
```

## Run the entire stack (MySQL + GoShort + PHPMyAdmin) using Docker
```bash
docker-compose up -d
```