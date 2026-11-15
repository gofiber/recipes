package controller

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/lib/pq"
)

// foreignKeyViolation is the PostgreSQL error code for a foreign key that is broken: a post whose
// author does not exist, or an author that posts still refer to.
const foreignKeyViolation = "23503"

func isForeignKeyViolation(err error) bool {
	var pqErr *pq.Error
	return errors.As(err, &pqErr) && pqErr.Code == foreignKeyViolation
}

// internalError logs the error and answers 500 without it. The messages of the database driver
// name tables and constraints, which the caller has no use for.
func internalError(c fiber.Ctx, err error) error {
	log.Printf("%s %q failed: %v", c.Method(), c.Path(), err)
	return c.Status(fiber.StatusInternalServerError).SendString("internal server error")
}
