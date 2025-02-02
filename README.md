# hu

hu is a lightweight package that provides simple HTTP utility functions to ease the creation of REST API responses. It offers a couple of helper functions that handle writing standard HTTP responses, as well as encoding responses to JSON.

## Features

- Send JSON responses with custom HTTP status codes.
- Easily emit HTTP responses with a specific status code.
- Helper for easily responding with a Bad Request response.

## Installation

To add hu to your project, you can simply use `go get`:

```bash
go get github.com/flaticols/hu
```

## Usage

Import the package in your Go application:

```go
import "github.com/flaticols/hu"
```

### Sending a JSON Response

Use the function RespJSON to send a JSON-encoded response:

```go
data := map[string]string{
    "message": "hello world",
}
err := hu.RespJSON(w, http.StatusOK, data)
if err != nil {
    // handle error
}
```

This function sets a proper Content-Type header (`application/json`), marshals the data to JSON, sets the status code, and writes the response body.

### Sending a Response without a body

If you need to just send a specific status code without a response body, use:

```go
hu.Resp(w, http.StatusNoContent)
```

### Responding with a Bad Request

To send a standardized JSON response for a bad request, use:

```go
data := map[string]string{
    "error": "invalid request parameters",
}
err := hu.RespBad(w, data)
if err != nil {
    // handle error
}
```

This helper is equivalent to calling RespJSON with HTTP Status 400 (Bad Request).

## Testing

The package comes with a comprehensive test suite ensuring that all helper functions work as expected. To run the tests, execute:

```bash
go test ./...
```

The tests cover:

- Successful JSON responses.
- Handling errors in JSON marshaling.
- Simple status code responses.
- Specialized handling of bad requests.

## Contributing

Contributions are welcome! If you have ideas for improvements or find a bug, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
