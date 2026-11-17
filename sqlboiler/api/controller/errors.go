package controller

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

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

// lookupFailed answers a query that found nothing or failed: 404 when the row does not exist,
// 500 for anything else, such as a database that is down.
func lookupFailed(c fiber.Ctx, err error, what string) error {
	if errors.Is(err, sql.ErrNoRows) {
		return c.Status(fiber.StatusNotFound).JSON(what + " not found")
	}
	return internalError(c, err)
}

// internalError logs the error and answers 500 without it. The messages of the database driver
// name tables and constraints, which the caller has no use for.
func internalError(c fiber.Ctx, err error) error {
	log.Printf("%s %q failed: %v", c.Method(), c.Path(), err)
	return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
}

// parseID reads the :id route parameter. The ids are 32-bit serial columns, so a larger number
// cannot exist and is rejected like any other value that is not an id.
func parseID(c fiber.Ctx) (int, error) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	return int(id), err
}
