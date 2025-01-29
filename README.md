# Simple Info API - HNG Stage one

A lightweight Go API that provides basic user information including email, current UTC timestamp, and GitHub URL. Built with native Go HTTP packages and featuring CORS support.

## Project Overview

This API serves a simple GET endpoint that returns user information in JSON format, including a dynamically generated UTC timestamp in ISO 8601 format.

## Features

- GET endpoint returning user information
- Dynamic UTC timestamp generation
- CORS support
- Clean JSON response format
- Error handling

## Prerequisites

- Go 1.16 or higher
- Git (for cloning the repository)

## Installation & Setup locally

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/your-repo
   cd your-repo
   ```

2. Build the project:
   ```bash
   go build
   ```

3. Run the server:
   ```bash
   ./your-repo
   ```
   The server will start on port 8087.

## API Documentation

### Endpoint

```
GET http://localhost:8087
```

### Response Format

```json
{
  "email": "your-email@example.com",
  "current_datetime": "2025-01-30T09:30:00Z",
  "github_url": "https://github.com/yourusername/your-repo"
}
```

#### Fields

- `email`: Your email address
- `current_datetime`: Current UTC timestamp in ISO 8601 format
- `github_url`: Your GitHub repository URL

### Status Codes

- 200: Successful request
- 405: Method not allowed
- 500: Internal server error

### Example Usage

Using curl:
```bash
curl http://localhost:8087
```

Using JavaScript fetch:
```javascript
fetch('http://localhost:8087')
  .then(response => response.json())
  .then(data => console.log(data));
```

## CORS Support

The API includes CORS headers to allow cross-origin requests:
- Allowed Methods: GET, OPTIONS
- Allowed Headers: Content-Type, Authorization
- Allowed Origins: * (all origins - can be restricted in production)

## Project Structure

```
.
├── main.go          # Main application file
├── README.md        # This documentation
└── go.mod          # Go module file
```

## Development

Want to contribute? Great! Here are the steps:

1. Fork the repo
2. Create a new branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Need Golang Developers?

Looking to hire skilled Golang developers? Check out the talented pool of developers at [HNG Golang Developers](https://hng.tech/hire/golang-developers).


## Contact

Your Name - stephennwac007@gmail.com
Project Link: [https://github.com/stephennwachukwu/stageone](https://github.com/stephennwachukwu/stageone)