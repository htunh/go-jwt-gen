# JWT Payment Token Example

This repository contains a Go application that demonstrates how to create and decode a JWT token for payment processing. The application creates a JWT token with payment details, sends it to a payment gateway via an HTTP POST request, and then decodes the response token to extract the payment details.

## Prerequisites

To run this application, you need to have the following installed on your machine:

- [Go](https://golang.org/dl/) (version 1.16 or later)

## Libraries Used

The application uses the following Go libraries:

- `encoding/json` for JSON encoding and decoding
- `fmt` for formatted I/O
- `log` for logging errors
- `net/http` for making HTTP requests
- `strings` for string manipulation
- `time` for handling time and dates
- `github.com/dgrijalva/jwt-go` for creating and decoding JWT tokens

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/go-jwt-gen.git
    cd go-jwt-gen
    ```

2. Install the required Go libraries:

    ```sh
    go get github.com/dgrijalva/jwt-go
    ```

## Usage

To run the application, execute the following command:

```sh
go run main.go
