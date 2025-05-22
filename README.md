# URL Shortener Service

A modern URL shortener service built with Go, featuring REST APIs, routing, and Redis storage.

## Features

- Create short links that redirect to original URLs
- Click tracking and analytics
- RESTful API design
- Redis storage for high performance
- Modern routing with Chi router

## Prerequisites

- Go 1.21 or higher
- Redis server
- Make (optional, for using Makefile commands)

## Project Structure

```
.
├── cmd/
│   └── api/            # Application entry point
├── internal/
│   ├── config/         # Configuration management
│   ├── handlers/       # HTTP handlers
│   ├── middleware/     # Custom middleware
│   ├── models/         # Data models
│   └── storage/        # Database interactions
├── pkg/
│   └── utils/          # Utility functions
├── go.mod
├── go.sum
└── README.md
```

## Getting Started

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set up environment variables:
   ```bash
   cp .env.example .env
   ```
4. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## API Endpoints

- `POST /api/v1/shorten` - Create a new short URL
- `GET /{shortCode}` - Redirect to original URL
- `GET /api/v1/stats/{shortCode}` - Get URL statistics

## License

MIT
