// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

// Database instance
var db *sql.DB

// Database settings
const (
	host     = "localhost"
	port     = 5432 // Default port
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

// Employees struct
type Employees struct {
	Employees []Employee `json:"employees"`
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

	// Get all records from MySQL
	app.Get("/employee", func(c *fiber.Ctx) error {
		// Get Employee list from database
		rows, err := db.Query("SELECT id, name, salary, age FROM employees order by id")
		if err != nil {
			return c.Status(500).SendString(err.Error())
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
		// Return Employees in JSON format
		return c.JSON(result)
	})

	// Add record into MySQL
	app.Post("/employee", func(c *fiber.Ctx) error {
		// New Employee struct
		u := new(Employee)

		// Parse body into struct
		if err := c.BodyParser(u); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Insert Employee into database
		res, err := db.Query("INSERT INTO employees (NAME, SALARY, AGE) VALUES (?, ?, ?)", u.Name, u.Salary, u.Age)
		if err != nil {
			return err
		}

		// Print result
		log.Println(res)

		// Return Employee in JSON format
		return c.JSON(u)
	})

	// Update record into MySQL
	app.Put("/employee", func(c *fiber.Ctx) error {
		// New Employee struct
		u := new(Employee)

		// Parse body into struct
		if err := c.BodyParser(u); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Update Employee record in database
		res, err := db.Query("UPDATE employees SET name=?,salary=?,age=? WHERE id=?", u.Name, u.Salary, u.Age, u.ID)
		if err != nil {
			return err
		}

		// Print result
		log.Println(res)

		// Return Employee in JSON format
		return c.Status(201).JSON(u)
	})

	// Delete record from MySQL
	app.Delete("/employee", func(c *fiber.Ctx) error {
		// New Employee struct
		u := new(Employee)

		// Parse body into struct
		if err := c.BodyParser(u); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Delete Employee from database
		res, err := db.Query("DELETE FROM employees WHERE id = ?", u.ID)
		if err != nil {
			return err
		}

		// Print result
		log.Println(res)

		// Return Employee in JSON format
		return c.JSON("Deleted")
	})

	log.Fatal(app.Listen(":3000"))
}
