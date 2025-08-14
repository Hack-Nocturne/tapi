# Simple Go Application

A simple Go web application that prints environment variables and provides HTTP endpoints.

## Features

- **Environment Variables Display**: Prints all environment variables in a prettified format on startup
- **Hello World Endpoint**: Simple HTTP endpoint that returns a JSON response
- **Environment Variables API**: HTTP endpoint that returns all environment variables as JSON

## Endpoints

- `GET /` - Hello World endpoint
- `GET /env` - Returns all environment variables as JSON

## Running the Application

1. Make sure you have Go installed (version 1.21 or later)
2. Navigate to the project directory
3. Run the application:

```bash
go run main.go
```

The application will:
1. Print all environment variables to the console in a prettified format
2. Start an HTTP server on port 8080 (or the port specified in the `PORT` environment variable)

## Example Usage

### Start the server
```bash
go run main.go
```

### Access the Hello World endpoint
```bash
curl http://localhost:8080/
```

### Access the Environment Variables endpoint
```bash
curl http://localhost:8080/env
```

## Environment Variables

- `PORT`: Specify the port for the HTTP server (defaults to 8080)

## Build and Run

To build a binary:
```bash
go build -o simple-go-app main.go
./simple-go-app
```

## Output Format

### Console Output
Environment variables are printed in a sorted, formatted table with truncated values for readability.

### JSON Responses
All HTTP endpoints return properly formatted JSON responses with appropriate headers.
