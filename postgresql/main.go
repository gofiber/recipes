// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
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
	ID     string `json: "id"`
	Name   string `json: "name"`
	Salary string `json: "salary"`
	Age    string `json: "age"`
}

// Employees struct
type Employees struct {
	Employees []Employee `json: "employees"`
}

// Connect function
func Connect() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
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
	app.Get("/employee", func(c *fiber.Ctx) {
		// Insert Employee into database
		rows, err := db.Query("SELECT id, name, salary, age FROM employees order by id")
		if err != nil {
			c.Status(500).Send(err)
			return
		}
		defer rows.Close()
		result := Employees{}

		for rows.Next() {
			employee := Employee{}
			err := rows.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age)
			// Exit if we get an error
			if err != nil {
				c.Status(500).Send(err)
				return
			}
			// Append Employee to Employees
			result.Employees = append(result.Employees, employee)
		}
		// Return Employees in JSON format
		if err := c.JSON(result); err != nil {
			c.Status(500).Send(err)
			return
		}
	})

	// Add record into postgreSQL
	app.Post("/employee", func(c *fiber.Ctx) {
		// New Employee struct
		u := new(Employee)
		// Parse body into struct
		if err := c.BodyParser(u); err != nil {
			c.Status(400).Send(err)
			return
		}
		// Insert Employee into database
		res, err := db.Query("INSERT INTO employees (name, salary, age)VALUES ($1, $2, $3)", u.Name, u.Salary, u.Age)
		if err != nil {
			c.Status(500).Send(err)
			return
		}
		// Print result
		log.Println(res)
		// Return Employee in JSON format
		if err := c.JSON(u); err != nil {
			c.Status(500).Send(err)
			return
		}
	})

	// Update record into postgreSQL
	app.Put("/employee", func(c *fiber.Ctx) {
		// New Employee struct
		u := new(Employee)
		// Parse body into struct
		if err := c.BodyParser(u); err != nil {
			c.Status(400).Send(err)
			return
		}
		// Insert Employee into database
		res, err := db.Query("UPDATE employees SET name=$1,salary=$2,age=$3 WHERE id=$5", u.Name, u.Salary, u.Age, u.ID)
		if err != nil {
			c.Status(500).Send(err)
			return
		}
		// Print result
		log.Println(res)
		// Return Employee in JSON format
		if err := c.Status(201).JSON(u); err != nil {
			c.Status(500).Send(err)
			return
		}
	})

	// Delete record from postgreSQL
	app.Delete("/employee", func(c *fiber.Ctx) {
		// New Employee struct
		u := new(Employee)
		// Parse body into struct
		if err := c.BodyParser(u); err != nil {
			c.Status(400).Send(err)
			return
		}
		// Insert Employee into database
		res, err := db.Query("DELETE FROM employees WHERE id = $1", u.ID)
		if err != nil {
			c.Status(500).Send(err)
			return
		}
		// Print result
		log.Println(res)
		// Return Employee in JSON format
		if err := c.JSON("Deleted"); err != nil {
			c.Status(500).Send(err)
			return
		}
	})

	app.Listen(3000)
}
