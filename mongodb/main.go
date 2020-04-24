// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber
package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gofiber/fiber"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

// Database settings (insert your own database name and connection URI)
const dbName = "fiber_test"
const mongoURI = "mongodb://user:password@localhost:27017/" + dbName

// Employee struct
type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

// Connect configures the MongoDB client and initializes the database connection.
// Source: https://www.mongodb.com/blog/post/quick-start-golang--mongodb--starting-and-setup
func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func main() {
	// Connect to the database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()

	// Get all employee records from MongoDB
	// Docs: https://docs.mongodb.com/manual/reference/command/find/
	app.Get("/employee", func(c *fiber.Ctx) {
		// get all records as a cursor
		query := bson.D{{}}
		cursor, err := mg.Db.Collection("employees").Find(c.Fasthttp, query)
		if err != nil {
			c.Status(500).Send(err)
			return
		}

		var employees []Employee = make([]Employee, 0)

		// iterate the cursor and decode each item into an Employee
		if err := cursor.All(c.Fasthttp, &employees); err != nil {
			c.Status(500).Send(err)
			return
		}

		// return employees list in JSON format
		if err := c.JSON(employees); err != nil {
			c.Status(500).Send(err)
			return
		}
	})

	// Insert a new employee into MongoDB
	// Docs: https://docs.mongodb.com/manual/reference/command/insert/
	app.Post("/employee", func(c *fiber.Ctx) {
		collection := mg.Db.Collection("employees")

		// New Employee struct
		employee := new(Employee)
		// Parse body into struct
		if err := c.BodyParser(employee); err != nil {
			c.Status(400).Send(err)
			return
		}

		// force MongoDB to always set its own generated ObjectIDs
		employee.ID = ""

		// insert the record
		insertionResult, err := collection.InsertOne(c.Fasthttp, employee)
		if err != nil {
			c.Status(500).Send(err)
			return
		}

		// get the just inserted record in order to return it as response
		filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
		createdRecord := collection.FindOne(c.Fasthttp, filter)

		// decode the Mongo record into Employee
		createdEmployee := &Employee{}
		createdRecord.Decode(createdEmployee)

		// return the created Employee in JSON format
		if err := c.Status(201).JSON(createdEmployee); err != nil {
			c.Status(500).Send(err)
			return
		}
	})

	// Update an employee record in MongoDB
	// Docs: https://docs.mongodb.com/manual/reference/command/findAndModify/
	app.Put("/employee/:id", func(c *fiber.Ctx) {
		idParam := c.Params("id")
		employeeID, err := primitive.ObjectIDFromHex(idParam)

		// the provided ID might be invalid ObjectID
		if err != nil {
			c.Status(400).Send()
			return
		}

		employee := new(Employee)
		// Parse body into struct
		if err := c.BodyParser(employee); err != nil {
			c.Status(400).Send(err)
			return
		}

		// Find the employee and update its data
		query := bson.D{{Key: "_id", Value: employeeID}}
		update := bson.D{
			{Key: "$set",
				Value: bson.D{
					{Key: "name", Value: employee.Name},
					{Key: "age", Value: employee.Age},
					{Key: "salary", Value: employee.Salary},
				},
			},
		}
		err = mg.Db.Collection("employees").FindOneAndUpdate(c.Fasthttp, query, update).Err()

		if err != nil {
			// ErrNoDocuments means that the filter did not match any documents in the collection
			if err == mongo.ErrNoDocuments {
				c.Status(404).Send()
				return
			}
			c.Status(500).Send()
			return
		}

		// return the updated employee
		employee.ID = idParam
		if err := c.Status(200).JSON(employee); err != nil {
			c.Status(500).Send(err)
			return
		}
	})

	// Delete an employee from MongoDB
	// Docs: https://docs.mongodb.com/manual/reference/command/delete/
	app.Delete("/employee/:id", func(c *fiber.Ctx) {
		employeeID, err := primitive.ObjectIDFromHex(
			c.Params("id"),
		)

		// the provided ID might be invalid ObjectID
		if err != nil {
			c.Status(400).Send()
			return
		}

		// find and delete the employee with the given ID
		query := bson.D{{Key: "_id", Value: employeeID}}
		result, err := mg.Db.Collection("employees").DeleteOne(c.Fasthttp, &query)

		if err != nil {
			c.Status(500).Send()
			return
		}

		// the employee might not exist
		if result.DeletedCount < 1 {
			c.Status(404).Send()
			return
		}

		// the record was deleted
		c.Status(204).Send()
	})

	app.Listen(3000)
}
