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

	mg = MongoInstance{
		Client: client,
		Db:     client.Database(dbName),
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
		cursor, err := mg.Db.Collection("employees").Find(c.Context(), bson.D{})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer cursor.Close(c.Context())

		var employees []Employee
		if err := cursor.All(c.Context(), &employees); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(employees)
	})

	// Get once employee records from MongoDB
	// Docs: https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-read-documents
	app.Get("/employee/:id", func(c *fiber.Ctx) error {
		id, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.SendStatus(400)
		}

		var employee Employee
		err = mg.Db.Collection("employees").FindOne(c.Context(), bson.M{"_id": id}).Decode(&employee)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.SendStatus(404)
			}
			return c.SendStatus(500)
		}

		return c.JSON(employee)
	})

	// Insert a new employee into MongoDB
	// Docs: https://docs.mongodb.com/manual/reference/command/insert/
	app.Post("/employee", func(c *fiber.Ctx) error {
		var employee Employee
		if err := c.BodyParser(&employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		employee.ID = ""
		result, err := mg.Db.Collection("employees").InsertOne(c.Context(), employee)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		employee.ID = result.InsertedID.(primitive.ObjectID).Hex()
		return c.Status(201).JSON(employee)
	})

	// Update an employee record in MongoDB
	// Docs: https://docs.mongodb.com/manual/reference/command/findAndModify/
	app.Put("/employee/:id", func(c *fiber.Ctx) error {
		id, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.SendStatus(400)
		}

		var employee Employee
		if err := c.BodyParser(&employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		update := bson.M{"$set": bson.M{
			"name":   employee.Name,
			"age":    employee.Age,
			"salary": employee.Salary,
		}}

		err = mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), bson.M{"_id": id}, update).Err()
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.SendStatus(404)
			}
			return c.SendStatus(500)
		}

		employee.ID = c.Params("id")
		return c.JSON(employee)
	})

	// Delete an employee from MongoDB
	// Docs: https://docs.mongodb.com/manual/reference/command/delete/
	app.Delete("/employee/:id", func(c *fiber.Ctx) error {
		id, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.SendStatus(400)
		}

		result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), bson.M{"_id": id})
		if err != nil {
			return c.SendStatus(500)
		}

		if result.DeletedCount == 0 {
			return c.SendStatus(404)
		}

		return c.SendStatus(204)
	})

	log.Fatal(app.Listen(":3000"))
}
