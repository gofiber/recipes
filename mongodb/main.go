// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gofiber/fiber/v2"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

// Database settings (insert your own database name and connection URI)
const (
	dbName   = "fiber_test"
	mongoURI = "mongodb://user:password@localhost:27017/" + dbName
)

// Employee struct
type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

// Connect configures the MongoDB client and initializes the database connection.
// Source: https://www.mongodb.com/docs/drivers/go/current/quick-start/
func Connect() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

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
	app.Get("/employee", func(c *fiber.Ctx) error {
		// get all records as a cursor
		query := bson.D{{}}
		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var employees []Employee = make([]Employee, 0)

		// iterate the cursor and decode each item into an Employee
		if err := cursor.All(c.Context(), &employees); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		// return employees list in JSON format
		return c.JSON(employees)
	})

	// Get once employee records from MongoDB
	// Docs: https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-read-documents
	app.Get("/employee/:id", func(c *fiber.Ctx) error {
		// get id by params
		params := c.Params("id")

		_id, err := primitive.ObjectIDFromHex(params)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		filter := bson.D{{Key: "_id", Value: _id}}

		var result Employee

		if err := mg.Db.Collection("employees").FindOne(c.Context(), filter).Decode(&result); err != nil {
			return c.Status(500).SendString("Something went wrong.")
		}

		return c.Status(fiber.StatusOK).JSON(result)
	})

	// Insert a new employee into MongoDB
	// Docs: https://docs.mongodb.com/manual/reference/command/insert/
	app.Post("/employee", func(c *fiber.Ctx) error {
		collection := mg.Db.Collection("employees")

		// New Employee struct
		employee := new(Employee)
		// Parse body into struct
		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// force MongoDB to always set its own generated ObjectIDs
		employee.ID = ""

		// insert the record
		insertionResult, err := collection.InsertOne(c.Context(), employee)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		// get the just inserted record in order to return it as response
		filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
		createdRecord := collection.FindOne(c.Context(), filter)

		// decode the Mongo record into Employee
		createdEmployee := &Employee{}
		createdRecord.Decode(createdEmployee)

		// return the created Employee in JSON format
		return c.Status(201).JSON(createdEmployee)
	})

	// Update an employee record in MongoDB
	// Docs: https://docs.mongodb.com/manual/reference/command/findAndModify/
	app.Put("/employee/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		employeeID, err := primitive.ObjectIDFromHex(idParam)
		// the provided ID might be invalid ObjectID
		if err != nil {
			return c.SendStatus(400)
		}

		employee := new(Employee)
		// Parse body into struct
		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Find the employee and update its data
		query := bson.D{{Key: "_id", Value: employeeID}}
		update := bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{Key: "name", Value: employee.Name},
					{Key: "age", Value: employee.Age},
					{Key: "salary", Value: employee.Salary},
				},
			},
		}
		err = mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()
		if err != nil {
			// ErrNoDocuments means that the filter did not match any documents in the collection
			if err == mongo.ErrNoDocuments {
				return c.SendStatus(404)
			}
			return c.SendStatus(500)
		}

		// return the updated employee
		employee.ID = idParam
		return c.Status(200).JSON(employee)
	})

	// Delete an employee from MongoDB
	// Docs: https://docs.mongodb.com/manual/reference/command/delete/
	app.Delete("/employee/:id", func(c *fiber.Ctx) error {
		employeeID, err := primitive.ObjectIDFromHex(
			c.Params("id"),
		)
		// the provided ID might be invalid ObjectID
		if err != nil {
			return c.SendStatus(400)
		}

		// find and delete the employee with the given ID
		query := bson.D{{Key: "_id", Value: employeeID}}
		result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), &query)
		if err != nil {
			return c.SendStatus(500)
		}

		// the employee might not exist
		if result.DeletedCount < 1 {
			return c.SendStatus(404)
		}

		// the record was deleted
		return c.SendStatus(204)
	})

	log.Fatal(app.Listen(":3000"))
}
