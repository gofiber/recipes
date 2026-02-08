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

	"github.com/gofiber/fiber/v3"
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
	app.Get("/employee", func(c fiber.Ctx) error {
		cursor, err := mg.Db.Collection("employees").Find(c.RequestCtx(), bson.D{})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer cursor.Close(c.RequestCtx())

		var employees []Employee
		if err := cursor.All(c.RequestCtx(), &employees); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(employees)
	})

	// Get once employee records from MongoDB
	// Docs: https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-read-documents
	app.Get("/employee/:id", func(c fiber.Ctx) error {
		id, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		var employee Employee
		err = mg.Db.Collection("employees").FindOne(c.RequestCtx(), bson.M{"_id": id}).Decode(&employee)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.SendStatus(fiber.StatusNotFound)
			}
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(employee)
	})

	// Insert a new employee into MongoDB
	// Docs: https://docs.mongodb.com/manual/reference/command/insert/
	app.Post("/employee", func(c fiber.Ctx) error {
		var employee Employee
		if err := c.Bind().Body(&employee); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		employee.ID = ""
		result, err := mg.Db.Collection("employees").InsertOne(c.RequestCtx(), employee)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		insertedID, ok := result.InsertedID.(primitive.ObjectID)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).SendString("failed to convert inserted ID to ObjectID")
		}

		employee.ID = insertedID.Hex()
		return c.Status(fiber.StatusCreated).JSON(employee)
	})

	// Update an employee record in MongoDB
	// Docs: https://docs.mongodb.com/manual/reference/command/findAndModify/
	app.Put("/employee/:id", func(c fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		var employee Employee
		if err := c.Bind().Body(&employee); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		update := bson.M{"$set": bson.M{
			"name":   employee.Name,
			"age":    employee.Age,
			"salary": employee.Salary,
		}}
		err = mg.Db.Collection("employees").FindOneAndUpdate(c.RequestCtx(), bson.M{"_id": id}, update).Err()
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.SendStatus(fiber.StatusNotFound)
			}
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// return the updated employee
		employee.ID = idParam
		return c.Status(fiber.StatusOK).JSON(employee)
	})

	// Delete an employee from MongoDB
	// Docs: https://docs.mongodb.com/manual/reference/command/delete/
	app.Delete("/employee/:id", func(c fiber.Ctx) error {
		id, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		result, err := mg.Db.Collection("employees").DeleteOne(c.RequestCtx(), bson.M{"_id": id})
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if result.DeletedCount == 0 {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.SendStatus(fiber.StatusNoContent)
	})

	log.Fatal(app.Listen(":3000"))
}

// fiber:context-methods migrated
