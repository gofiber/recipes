// ⚡️ Fiber: A Go-based web framework inspired by Express
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var mc *minio.Client

const (
	endpoint        = "127.0.0.1:9000" // MinIO endpoint (local MinIO server)
	accessKeyID     = "minioadmin"     // Access key for MinIO
	secretAccessKey = "minioadmin"     // Secret access key for MinIO
	useSSL          = false            // Whether to use SSL (set to true for secure connections)
	bucketName      = "fiber"          // Name of the bucket in MinIO
)

func main() {
	// Initialize the MinIO client
	if err := newMinioClient(); err != nil {
		log.Fatalf("Error initializing MinIO client: %v", err) // Exit if client setup fails
	}

	// Create a new Fiber instance for the web service
	app := fiber.New()

	// Define the route for uploading multiple files
	app.Post("/upload", func(c *fiber.Ctx) error {
		// Retrieve all files from the multipart form, under the field name "documents"
		multipartForm, err := c.MultipartForm()
		if err != nil {
			// Return a bad request response if there is an issue with retrieving the form data
			return c.Status(http.StatusBadRequest).SendString("Error retrieving form: " + err.Error())
		}

		files := multipartForm.File["documents"] // Extract the files with field name "documents"
		if len(files) == 0 {
			// Return a bad request response if no files are found in the request
			return c.Status(http.StatusBadRequest).SendString("No files found in request")
		}

		var uploadedFiles []fiber.Map // List to store information about successfully uploaded files
		var failedFiles []fiber.Map   // List to store information about files that failed to upload

		// Iterate over the files and upload each one
		for _, filePart := range files {

			// Validate the filename to ensure it is non-empty and contains only allowed characters
			err := validateFilename(filePart.Filename)
			if err != nil {
				// If validation fails, add the file to the failed list and continue
				failedFiles = append(failedFiles, fiber.Map{
					"filename": filePart.Filename,
					"error":    fmt.Sprintf("Invalid filename: %v", err),
				})
				continue // Skip to the next file
			}

			// Open the file for reading before upload
			file, err := filePart.Open()
			if err != nil {
				// If opening the file fails, add it to the failed list
				failedFiles = append(failedFiles, fiber.Map{
					"filename": filePart.Filename,
					"error":    fmt.Sprintf("Error opening file: %v", err),
				})
				continue // Skip to the next file
			}
			defer file.Close() // Ensure the file is closed after the upload

			// Upload the file to MinIO
			uploadInfo, err := mc.PutObject(
				c.Context(),
				bucketName,        // Bucket name
				filePart.Filename, // File name in the MinIO bucket
				file,              // File data to upload
				filePart.Size,     // File size
				minio.PutObjectOptions{ // Options for file upload
					PartSize:    5 * 1024 * 1024,            // Upload in chunks (5 MB per part)
					ContentType: "application/octet-stream", // Default content type for binary files
				},
			)
			if err != nil {
				// If the upload fails, add the file to the failed list
				failedFiles = append(failedFiles, fiber.Map{
					"filename": filePart.Filename,
					"error":    fmt.Sprintf("Error uploading file: %v", err),
				})
				continue // Skip to the next file
			}

			// Log upload details (ETag, Size)
			log.Printf("File uploaded successfully: ETag: %s, Size: %d", uploadInfo.ETag, uploadInfo.Size)

			// Generate a URL for the uploaded file
			protocol := "http"
			if c.Protocol() == "https" {
				protocol = "https"
			}
			fileURL := fmt.Sprintf("%s://%s/file/%s", protocol, c.Hostname(), filePart.Filename)

			uploadedFiles = append(uploadedFiles, fiber.Map{
				"filename": filePart.Filename,
				"url":      fileURL, // Include the URL to access the file
			})
		}

		// Return the results of the upload attempts
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"uploaded": uploadedFiles, // List of successfully uploaded files
			"failed":   failedFiles,   // List of files that failed to upload
		})
	})

	// Define the route to retrieve files by filename
	app.Get("/file/:filename", func(c *fiber.Ctx) error {
		// Get the filename from the URL parameter
		filename := c.Params("filename")

		// Validate the filename to ensure it is non-empty and contains only allowed characters
		if err := validateFilename(filename); err != nil {
			// Return a bad request response if the filename is invalid (empty or contains disallowed characters)
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		// Check if the file exists in the MinIO bucket
		_, err := mc.StatObject(c.Context(), bucketName, filename, minio.StatObjectOptions{})
		if err != nil {
			// Return a not found response if the file doesn't exist
			return c.Status(http.StatusNotFound).SendString("File not found")
		}

		// Retrieve the file from MinIO
		object, err := mc.GetObject(c.Context(), bucketName, filename, minio.GetObjectOptions{})
		if err != nil {
			// Return an internal server error if there's an issue retrieving the file
			return c.Status(http.StatusInternalServerError).SendString("Error retrieving file: " + err.Error())
		}

		// Set HTTP headers to indicate a file download
		c.Set(fiber.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", filename))
		c.Set(fiber.HeaderContentType, fiber.MIMEOctetStream)

		// Stream the file contents to the client
		return c.SendStream(object)
	})

	// Start the Fiber server on port 3000
	log.Fatal(app.Listen(":3000"))
}

// newMinioClient initializes a new MinIO client and ensures the "fiber" bucket exists.
func newMinioClient() error {
	// Initialize the MinIO client with the specified credentials and endpoint
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		// Return an error if the client initialization fails
		return fmt.Errorf("failed to create MinIO client: %w", err)
	}
	mc = minioClient

	// Ensure the "fiber" bucket exists, creating it if necessary
	location := "us-east-1"
	if err := ensureBucketExists(context.Background(), mc, bucketName, location); err != nil {
		return err
	}

	return nil
}

// ensureBucketExists checks if a bucket exists, and creates it if it does not.
func ensureBucketExists(ctx context.Context, client *minio.Client, bucketName, location string) error {
	// Check if the bucket already exists
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		// Return an error if there is an issue checking the bucket's existence
		return fmt.Errorf("error checking if bucket exists: %w", err)
	}

	// If the bucket doesn't exist, create it
	if !exists {
		if err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location}); err != nil {
			// Return an error if the bucket creation fails
			return fmt.Errorf("failed to create bucket: %w", err)
		}
		log.Printf("Bucket %s created successfully", bucketName)
	} else {
		log.Printf("Bucket %s already exists", bucketName)
	}

	return nil
}

// validateFilename checks if the filename contains only alphanumeric characters, underscores, dashes, and dots.
func validateFilename(filename string) error {
	// Check if the filename is empty
	if filename == "" {
		// Return an error if the filename is empty
		return fmt.Errorf("invalid filename: filename cannot be empty")
	}

	// Check filename length
	if len(filename) > 255 {
		return fmt.Errorf("invalid filename: exceeds maximum length of %d characters", 255)
	}

	// Prevent path traversal
	if strings.Contains(filename, "/") || strings.Contains(filename, "\\") || strings.HasPrefix(filename, "..") {
		return fmt.Errorf("invalid filename: path traversal is not allowed")
	}

	// Prevent hidden files (files starting with a dot)
	if strings.HasPrefix(filename, ".") {
		return fmt.Errorf("invalid filename: hidden files not allowed")
	}

	// Check each character in the filename for invalid characters
	for _, char := range filename {
		if !isAlphaNumericOrSpecial(char) {
			// Return an error if any invalid character is found
			return fmt.Errorf("invalid filename: contains invalid characters")
		}
	}
	return nil
}

// isAlphaNumericOrSpecial checks if a character is alphanumeric or a valid special character.
func isAlphaNumericOrSpecial(char rune) bool {
	// Validate the character (alphanumeric, dash, underscore, or dot)
	switch {
	case 'A' <= char && char <= 'Z':
		return true
	case 'a' <= char && char <= 'z':
		return true
	case '0' <= char && char <= '9':
		return true
	case char == '-' || char == '_':
		return true
	case char == '.':
		return true
	default:
		return false
	}
}
