// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Database instance
var db *sql.DB

// Database settings
const (
	host     = "localhost"
	port     = 5432 // Default port
	user     = "postgres"
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

// Employees struct
type Employees struct {
	Employees []Employee `json:"employees"`
}

// Connect function
func Connect() error {
	var err error
	db, err = sql.Open("pgx", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
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

	// Get all records from postgreSQL
	app.Get("/employee", func(c fiber.Ctx) error {
		// Select all Employee(s) from database
		rows, err := db.Query("SELECT id, name, salary, age FROM employees order by id")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer rows.Close()
		result := Employees{}

		for rows.Next() {
			employee := Employee{}
			if err := rows.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age); err != nil {
				return err // Exit if we get an error
			}

			// Append Employee to Employees
			result.Employees = append(result.Employees, employee)
		}
		if err := rows.Err(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		// Return Employees in JSON format
		return c.JSON(result)
	})

	// Add record into postgreSQL
	app.Post("/employee", func(c fiber.Ctx) error {
		// New Employee struct
		u := new(Employee)

		// Parse body into struct
		if err := c.Bind().Body(u); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		// Insert Employee into database
		if _, err := db.Exec("INSERT INTO employees (name, salary, age) VALUES ($1, $2, $3)", u.Name, u.Salary, u.Age); err != nil {
			return err
		}

		// Return Employee in JSON format
		return c.JSON(u)
	})

	// Update record into postgreSQL
	app.Put("/employee/:id", func(c fiber.Ctx) error {
		// New Employee struct
		u := new(Employee)

		// Parse body into struct
		if err := c.Bind().Body(u); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		// Update Employee into database
		if _, err := db.Exec("UPDATE employees SET name=$1, salary=$2, age=$3 WHERE id=$4", u.Name, u.Salary, u.Age, c.Params("id")); err != nil {
			return err
		}

		// Return Employee in JSON format
		return c.Status(fiber.StatusCreated).JSON(u)
	})

	// Delete record from postgreSQL
	app.Delete("/employee/:id", func(c fiber.Ctx) error {
		// Delete Employee from database
		if _, err := db.Exec("DELETE FROM employees WHERE id = $1", c.Params("id")); err != nil {
			return err
		}

		// Return result in JSON format
		return c.JSON("Deleted")
	})

	log.Fatal(app.Listen(":3000"))
}
