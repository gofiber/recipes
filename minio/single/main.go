// ⚡️ Fiber: A Go-based web framework inspired by Express
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var mc *minio.Client

var (
	// MinIO configuration with defaults if environment variables are not set
	endpoint        = getEnv("FIBER_MINIO_ENDPOINT", "127.0.0.1:9000")          // MinIO endpoint (local MinIO server)
	accessKeyID     = getEnv("FIBER_MINIO_ACCESS_KEY", "minioadmin")            // Access key for MinIO
	secretAccessKey = getEnv("FIBER_MINIO_SECRET_KEY", "minioadmin")            // Secret access key for MinIO
	useSSL, _       = strconv.ParseBool(getEnv("FIBER_MINIO_USE_SSL", "false")) // Whether to use SSL (set to true for secure connections)
	bucketName      = getEnv("FIBER_MINIO_BUCKET_NAME", "fiber")                // Name of the bucket in MinIO
	location        = getEnv("FIBER_MINIO_LOCATION", "us-east-1")               // MinIO location
)

func main() {
	// Initialize the MinIO client
	if err := newMinioClient(); err != nil {
		log.Fatalf("Error initializing MinIO client: %v", err) // Exit if client setup fails
	}

	// Create a new Fiber instance for the web service
	app := fiber.New()

	// Define the route to handle file uploads
	app.Post("/upload", func(c *fiber.Ctx) error {
		// Get the file from the form data (input field name: "document")
		formFile, err := c.FormFile("document")
		if err != nil {
			// Return a bad request response if file retrieval fails
			return c.Status(http.StatusBadRequest).SendString("Error retrieving file: " + err.Error())
		}

		// Validate the filename to ensure it is non-empty and contains only allowed characters
		filename := formFile.Filename
		if err := validateFilename(filename); err != nil {
			// Return a bad request response if the filename is invalid
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		// Open the file for reading before upload
		file, err := formFile.Open()
		if err != nil {
			// Return an internal server error if the file can't be opened
			return c.Status(http.StatusInternalServerError).SendString("Error opening file: " + err.Error())
		}
		defer file.Close() // Ensure the file is closed after the upload

		// Upload the file to MinIO
		uploadInfo, err := mc.PutObject(
			c.Context(),
			bucketName,    // Bucket name
			filename,      // File name in the MinIO bucket
			file,          // File data to upload
			formFile.Size, // File size
			minio.PutObjectOptions{
				PartSize:    5 * 1024 * 1024,            // Chunk size for large files (5 MB per part)
				ContentType: "application/octet-stream", // Default content type for binary files
			},
		)
		if err != nil {
			// Return an internal server error if the upload fails
			return c.Status(http.StatusInternalServerError).SendString("Error uploading file: " + err.Error())
		}

		// Log upload details (ETag, Size)
		log.Printf("File uploaded: ETag: %s, Size: %d", uploadInfo.ETag, uploadInfo.Size)

		// Create a URL to access the uploaded file
		protocol := "http"
		if c.Protocol() == "https" {
			protocol = "https"
		}
		fileURL := fmt.Sprintf("%s://%s/file/%s", protocol, c.Hostname(), filename)

		// Return a successful response with the file details and URL
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message":  "File uploaded successfully",
			"fileName": filename,
			"url":      fileURL,
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

// Get environment variable or return default value if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
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
