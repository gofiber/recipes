# Autocert Example

This example demonstrates how to set up a secure Go Fiber application using Let's Encrypt for automatic TLS certificate management with `autocert`.

## Description

This project provides a starting point for building a secure web application with automatic TLS certificate management using Let's Encrypt. It leverages Fiber for the web framework and `autocert` for certificate management.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [Git](https://git-scm.com/downloads)

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/autocert
    ```

2. Install the dependencies:
    ```bash
    go mod download
    ```

3. Update the `HostPolicy` in `main.go` with your domain:
    ```go
    m := &autocert.Manager{
        Prompt: autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("yourdomain.com"), // Replace with your domain
        Cache: autocert.DirCache("./certs"),
    }
    ```

4. Run the application:
    ```bash
    go run main.go
    ```

The application should now be running on `https://localhost`.

## Example Usage

1. Open your browser and navigate to `https://yourdomain.com` (replace with your actual domain).

2. You should see the message: `This is a secure server 👮`.

## Conclusion

This example provides a basic setup for a Go Fiber application with automatic TLS certificate management using Let's Encrypt. It can be extended and customized further to fit the needs of more complex applications.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Let's Encrypt Documentation](https://letsencrypt.org/docs/)
- [Autocert Documentation](https://pkg.go.dev/golang.org/x/crypto/acme/autocert)
