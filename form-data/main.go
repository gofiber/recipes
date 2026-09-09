// This recipe answers gofiber/recipes#2707 wishlist item 1: how to parse
// application/x-www-form-urlencoded bodies into Go structs, slices and maps,
// including nested and repeated fields, using Fiber v3's Bind().Body().
package main

import (
	"log"
	"net/url"

	"github.com/gofiber/fiber/v3"
)

// ContactForm is a flat form: every field is a single scalar value.
type ContactForm struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// TagsForm has a repeated field. `tags=go&tags=web` binds into []string.
type TagsForm struct {
	Tags []string `json:"tags"`
}

// Address is nested inside Order via the dotted key path "address.street",
// "address.city".
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

// Item is one element of Order.Items, a slice of structs. The decoder needs
// an index to tell elements apart: "items.0.name", "items.0.qty",
// "items.1.name", "items.1.qty".
type Item struct {
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

// Order combines a nested struct field (Address) and a slice of structs
// (Items) in one top-level struct.
type Order struct {
	Customer string  `json:"customer"`
	Address  Address `json:"address"`
	Items    []Item  `json:"items"`
}

func main() {
	app := fiber.New()

	app.Post("/struct", bindStruct)
	app.Post("/slice", bindSlice)
	app.Post("/nested", bindNested)
	app.Post("/map", bindMap)

	log.Fatal(app.Listen(":3000"))
}

// bindStruct binds a flat form into a struct.
func bindStruct(c fiber.Ctx) error {
	var form ContactForm
	if err := c.Bind().Body(&form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(form)
}

// bindSlice binds a repeated field into a []string.
func bindSlice(c fiber.Ctx) error {
	var form TagsForm
	if err := c.Bind().Body(&form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(form)
}

// bindNested binds a nested struct field and a slice of structs in one form.
func bindNested(c fiber.Ctx) error {
	var order Order
	if err := c.Bind().Body(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(order)
}

// bindMap demonstrates parsing dynamic, caller-chosen field names into a
// map[string][]string using the standard library's net/url.ParseQuery.
// Alternatively, Fiber's Bind().Body(&mapTarget) also supports decoding
// form data directly into maps via its built-in form binder.
func bindMap(c fiber.Ctx) error {
	values, err := url.ParseQuery(string(c.Body()))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(values)
}
