// ⚡️ Fiber: A Go-based web framework inspired by Express
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/minio"
)

const maxFileSize = 10 * 1024 * 1024 // 10MB

func main() {
	// Initialize the MinIO client
	store := minio.New(minio.Config{
		Endpoint: getEnv("FIBER_MINIO_ENDPOINT", "localhost:9000"),
		Secure:   getEnv("FIBER_MINIO_USE_SSL", "false") == "true",
		Bucket:   getEnv("FIBER_MINIO_BUCKET", "fiber-bucket"),
		Region:   getEnv("FIBER_MINIO_REGION", "us-east-1"),
		Credentials: minio.Credentials{
			AccessKeyID:     getEnv("FIBER_MINIO_ACCESS_KEY", "minioadmin"),
			SecretAccessKey: getEnv("FIBER_MINIO_SECRET_KEY", "minioadmin"),
		},
	})

	// If the bucket doesn't exist, attempt to create it
	if err := store.CheckBucket(); err != nil {
		if err := store.CreateBucket(); err != nil {
			log.Fatalf("failed to create bucket: %v", err)
		}
	}

	// Create a new Fiber instance
	app := fiber.New()

	// Define the route to handle file uploads
	app.Post("/upload", func(c *fiber.Ctx) error {
		// Check file size before processing
		if c.Request().Header.ContentLength() > maxFileSize {
			return c.Status(http.StatusRequestEntityTooLarge).JSON(fiber.Map{
				"message": "file too large",
			})
		}

		// Get the file from the form data (input field name: "document")
		formFile, err := c.FormFile("document")
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "error retrieving file",
			})
		}

		// Double-check file size
		if formFile.Size > maxFileSize {
			return c.Status(http.StatusRequestEntityTooLarge).JSON(fiber.Map{
				"message": "file too large",
			})
		}

		// Validate the filename to ensure it is non-empty and contains only allowed characters
		filename := formFile.Filename
		if err := validateFilename(filename); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		// Open the file for reading before upload
		file, err := formFile.Open()
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "error opening file",
			})
		}
		defer file.Close() // Ensure the file is closed after the upload

		// Detect content type
		buffer := make([]byte, 512)
		_, err = file.Read(buffer)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "error reading file",
			})
		}
		minio.ConfigDefault.PutObjectOptions.ContentType = http.DetectContentType(buffer)

		// Reset file pointer
		_, err = file.Seek(0, 0)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "error resetting file",
			})
		}

		// Upload the file to MinIO
		uploadInfo, err := store.Conn().PutObject(
			c.Context(),
			minio.ConfigDefault.Bucket,           // Bucket name
			filename,                             // File name in the MinIO bucket
			file,                                 // File data to upload
			formFile.Size,                        // File size
			minio.ConfigDefault.PutObjectOptions, // content type for binary files
		)
		// If the upload fails, return Error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "error uploading file",
			})
		}

		// Log upload details (ETag, Size)
		log.Printf("file uploaded: ETag: %s, Size: %d", uploadInfo.ETag, uploadInfo.Size)

		// Create a URL to access the uploaded file
		protocol := "http"
		if c.Protocol() == "https" {
			protocol = "https"
		}
		fileURL := fmt.Sprintf("%s://%s/file/%s", protocol, c.Hostname(), filename)

		// Return a successful response with the file details and URL
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message":  "file uploaded successfully",
			"fileName": filename,
			"url":      fileURL,
		})
	})

	// Define the route to retrieve files by filename
	app.Get("/file/:filename", func(c *fiber.Ctx) error {
		// Get the filename from the URL parameter
		filename := c.Params("filename")

		// Validate the filename
		if err := validateFilename(filename); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		// Check if the file exists in the MinIO bucket
		_, err := store.Conn().StatObject(c.Context(), minio.ConfigDefault.Bucket, filename, minio.ConfigDefault.GetObjectOptions)
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "file not found",
			})
		}

		// Retrieve the file from MinIO
		object, err := store.Conn().GetObject(c.Context(), minio.ConfigDefault.Bucket, filename, minio.ConfigDefault.GetObjectOptions)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "error retrieving file",
			})
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

// validateFilename checks if the filename contains only alphanumeric characters, underscores, dashes, and dots.
func validateFilename(filename string) error {
	// Check if the filename is empty
	if filename == "" {
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
