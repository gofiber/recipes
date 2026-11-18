package routes

import (
	"log"
	"strconv"

	"ent-mysql/database"
	"ent-mysql/ent"

	"github.com/gofiber/fiber/v3"
)

// bookID parses the :id route parameter. A value that is not an integer used to turn
// silently into 0, which then failed as a confusing "not found".
func bookID(c fiber.Ctx) (int, error) {
	return strconv.Atoi(c.Params("id"))
}

func GetBook(c fiber.Ctx) error {
	id, err := bookID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid book ID")
	}
	b, err := database.DBConn.Book.
		Get(c.Context(), id)
	if err != nil {
		// ent reports a missing row as a *NotFoundError, not as a nil result.
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON("Not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	log.Println("book data: ", b)
	return c.Status(fiber.StatusOK).JSON(b)
}

func GetAllBook(c fiber.Ctx) error {
	b, err := database.DBConn.Book.
		Query().All(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	log.Println("book data All: ", b)
	return c.Status(fiber.StatusOK).JSON(b)
}

func CreateBook(c fiber.Ctx) error {
	body := struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}{}
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if body.Title == "" || body.Author == "" {
		return c.Status(fiber.StatusBadRequest).JSON("Not enough data for creating")
	}
	b, err := database.DBConn.Book.
		Create().
		SetTitle(body.Title).
		SetAuthor(body.Author).
		Save(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	log.Println("book was created: ", b)
	return c.Status(fiber.StatusOK).JSON(b)
}

func DeleteBook(c fiber.Ctx) error {
	id, err := bookID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid book ID")
	}
	err = database.DBConn.Book.
		DeleteOneID(id).
		Exec(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON("Not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	log.Println("book was deleted")
	return c.Status(fiber.StatusOK).JSON(nil)
}

func UpdateBook(c fiber.Ctx) error {
	title := c.Query("title")
	author := c.Query("author")
	id, err := bookID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid book ID")
	}
	if title == "" || author == "" {
		return c.Status(fiber.StatusBadRequest).JSON("Not enough data for updating")
	}
	b, err := database.DBConn.Book.
		UpdateOneID(id).
		SetTitle(title).
		SetAuthor(author).
		Save(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON("Not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	log.Println("book was updated: ", b)
	return c.Status(fiber.StatusOK).JSON(b)
}
