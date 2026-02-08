// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v3"
)

// Database instance
var db *sql.DB

// Database settings
const (
	user     = "root"
	password = "password"
	dbname   = "fiber_demo"
)

// Employee struct
type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    int     `json:"age"`
}

// Connect function
func Connect() error {
	var err error
	// Use DSN string to open
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func main() {
	// Connect with database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()

	app.Get("/employee", func(c fiber.Ctx) error {
		rows, err := db.Query("SELECT id, name, salary, age FROM employees ORDER BY id")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer rows.Close()

		var employees []Employee
		for rows.Next() {
			var emp Employee
			if err := rows.Scan(&emp.ID, &emp.Name, &emp.Salary, &emp.Age); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			employees = append(employees, emp)
		}
		if err := rows.Err(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(employees)
	})

	app.Get("/employee/:id", func(c fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := strconv.Atoi(idParam)
		if err != nil || id <= 0 {
			return c.Status(fiber.StatusBadRequest).SendString("invalid employee id")
		}

		var emp Employee
		err = db.QueryRow("SELECT id, name, salary, age FROM employees WHERE id = ?", id).Scan(&emp.ID, &emp.Name, &emp.Salary, &emp.Age)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.SendStatus(fiber.StatusNotFound)
			}
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(emp)
	})

	app.Post("/employee", func(c fiber.Ctx) error {
		var emp Employee
		if err := c.Bind().Body(&emp); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		result, err := db.Exec("INSERT INTO employees (name, salary, age) VALUES (?, ?, ?)", emp.Name, emp.Salary, emp.Age)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		id, err := result.LastInsertId()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		emp.ID = int(id)

		return c.Status(fiber.StatusCreated).JSON(emp)
	})

	app.Put("/employee/:id", func(c fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := strconv.Atoi(idParam)
		if err != nil || id <= 0 {
			return c.Status(fiber.StatusBadRequest).SendString("invalid employee id")
		}

		var emp Employee
		if err := c.Bind().Body(&emp); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		result, err := db.Exec("UPDATE employees SET name=?, salary=?, age=? WHERE id=?", emp.Name, emp.Salary, emp.Age, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		if rowsAffected == 0 {
			return c.SendStatus(fiber.StatusNotFound)
		}

		emp.ID = id
		return c.JSON(emp)
	})

	app.Delete("/employee/:id", func(c fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := strconv.Atoi(idParam)
		if err != nil || id <= 0 {
			return c.Status(fiber.StatusBadRequest).SendString("invalid employee id")
		}

		result, err := db.Exec("DELETE FROM employees WHERE id = ?", id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		if rowsAffected == 0 {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.SendStatus(fiber.StatusNoContent)
	})

	log.Fatal(app.Listen(":3000"))
}
