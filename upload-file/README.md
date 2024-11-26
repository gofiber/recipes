---
title: File Upload
keywords: [file upload, upload, form, multipart]
---

# File Upload Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/upload-file) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/upload-file)

This example demonstrates how to handle file uploads using Go Fiber.

## Description

This project provides a basic setup for handling file uploads in a Go Fiber application. It includes examples for uploading single and multiple files, as well as saving files to different directories.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [Git](https://git-scm.com/downloads)

## Project Structure

- `single/main.go`: Example for uploading a single file to the root directory.
- `single_relative_path/main.go`: Example for uploading a single file to a relative path.
- `multiple/main.go`: Example for uploading multiple files.
- `go.mod`: The Go module file.

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/upload-file
    ```

2. Install the dependencies:

    ```bash
    go mod download
    ```

## Running the Examples

### Single File Upload

1. Navigate to the `single` directory:

    ```bash
    cd single
    ```

2. Run the application:

    ```bash
    go run main.go
    ```

3. Use a tool like `curl` or Postman to upload a file:

    ```bash
    curl -F "document=@/path/to/your/file" http://localhost:3000/
    ```

### Single File Upload with Relative Path

1. Navigate to the `single_relative_path` directory:

    ```bash
    cd single_relative_path
    ```

2. Run the application:

    ```bash
    go run main.go
    ```

3. Use a tool like `curl` or Postman to upload a file:

    ```bash
    curl -F "document=@/path/to/your/file" http://localhost:3000/
    ```

### Multiple File Upload

1. Navigate to the `multiple` directory:

    ```bash
    cd multiple
    ```

2. Run the application:

    ```bash
    go run main.go
    ```

3. Use a tool like `curl` or Postman to upload multiple files:

    ```bash
    curl -F "documents=@/path/to/your/file1" -F "documents=@/path/to/your/file2" http://localhost:3000/
    ```

## Code Overview

### `single/main.go`

Handles uploading a single file to the root directory.

### `single_relative_path/main.go`

Handles uploading a single file to a relative path.

### `multiple/main.go`

Handles uploading multiple files.

## Conclusion

This example provides a basic setup for handling file uploads in a Go Fiber application. It can be extended and customized further to fit the needs of more complex applications.

## References

- [Fiber Documentation](https://docs.gofiber.io)
