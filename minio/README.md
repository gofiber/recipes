---
title: MinIO
keywords: [minio, file upload, file download]
description: A simple application for uploading and downloading files from MinIO.
---

# MinIO File Upload & Download Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/minio) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/minio)

This example demonstrates a simple Go Fiber application that includes modules for uploading both single and multiple files, as well as downloading files from MinIO. Each module provides REST API endpoints for file upload and retrieval, serving as a foundation for applications requiring file storage and access.

## Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/dl/): (version 1.22 or higher) installed
- [minio](https://min.io/download): MinIO running on your local machine or a remote server
- [Git](https://git-scm.com/downloads)

## Project Structure

- `single/main.go`: Example for uploading and downloading a single file to/from MinIO.

- `multiple/main.go`: Example for uploading multiple files to MinIO and downloading files from MinIO.

- `go.mod`: Go module file managing project dependencies.

## Getting Started

### 1. Clone the Repository

Clone the repository and navigate to the example directory:

```bash
git clone https://github.com/gofiber/recipes.git
cd recipes/minio
```

### 2. Install Dependencies

Use Go’s module system to install dependencies:

```bash
go mod download
```

## Running the Examples

### Uploading and Downloading a Single File

1. Go to the `single` directory:

   ```bash
   cd single
   ```

2. Start the application:

   ```bash
   go run main.go
   ```

3. Upload a file using `curl` or `Postman`:
   ```bash
   curl -F "document=@/path/to/your/file" http://localhost:3000/upload
   ```
4. Download the file by specifying its name in the request:

   ```bash
   curl -O http://localhost:3000/file/<filename>
   ```

### Uploading Multiple Files and Downloading Files

1. Go to the `multiple` directory:

   ```bash
   cd multiple
   ```

2. Start the application:

   ```bash
   go run main.go
   ```

3. Upload multiple files using `curl` or `Postman`:

   ```bash
   curl -F "documents=@/path/to/your/file1" -F "documents=@/path/to/your/file2" http://localhost:3000/upload
   ```

4. Download a file by specifying its name in the request.

   ```bash
   curl -O http://localhost:3000/file/<filename>
   ```

## Code Overview

### `single/main.go`

- Defines routes to handle a single file upload and download.

- Includes error handling for file validation, MinIO connection, and bucket management.

### `multiple/main.go`

- Handles uploading multiple files in a single request and allows for file downloads.

- Validates each file and provides detailed responses for both successful and failed uploads.

## Conclusion

This example offers a approach for managing file uploads and downloads with Go Fiber and MinIO. It can be expanded to support additional features, such as adding metadata, handling large files, or restricting access to files.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Fiber storage](https://github.com/gofiber/storage)
- [MinIO Documentation](https://min.io/docs/)
