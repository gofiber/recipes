package controllers

import (
	"github.com/gofiber/fiber"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/auth"
	"github.com/itsursujit/fiber-boilerplate/libraries"
	"github.com/itsursujit/fiber-boilerplate/models"
	"strconv"
)

func PlaceOrderFromPaypal(c *fiber.Ctx) {
	var order models.Payment
	c.BodyParser(&order)
	if amount, err := strconv.ParseFloat(order.Amount, 32); err != nil || order.Amount == "" || amount < 5 {
		c.SendStatus(400)
		c.JSON(fiber.Map{
			"validationError": true,
		})
		return
	}
	c.JSON(fiber.Map{
		"validationError": false,
	})
	pm, err := models.GetPaymentMethodBySlug("paypal")
	if err != nil {
		Flash.WithError(c, fiber.Map{
			"message": "Cannot make order at the moment",
		})
		c.Redirect("/")
		return
	}
	user, _ := auth.User(c)
	if user.EmailVerified != true {
		c.SendStatus(400)
		c.JSON(fiber.Map{
			"validationError": true,
			"message":         "Email is not verified, ",
		})
		return
	}
	order.PaymentMethodID = pm.ID
	order.Currency = pm.Currency
	order.UserID = user.ID
	order.Status = "PENDING"
	order.Create()
	err = libraries.CreateOrder(&order, user)
	DB.Save(&order)

	if err != nil {
		Flash.WithError(c, fiber.Map{
			"message": "Cannot make order at the moment",
		})
		c.Redirect("/")
		return
	}
	c.JSON(fiber.Map{
		"ack": true,
		"data": fiber.Map{
			"id": order.PayPalOrderDetail.ID,
		},
	})
}

func PostOrderResponseFromPaypal(c *fiber.Ctx) {

}

func ValidateOrderFromPaypal(c *fiber.Ctx) {
	amount := c.Params("amount")

	if _, err := strconv.ParseFloat(amount, 32); err != nil || amount == "" {
		c.SendStatus(400)
		c.JSON(fiber.Map{
			"validationError": true,
		})
		return
	}
	c.JSON(fiber.Map{
		"validationError": false,
	})
}

func GetOrderDetailFromPaypal(c *fiber.Ctx) {
	order, err := libraries.GetOrder(c.Params("id"))
	if err != nil {
		c.SendStatus(500)
		c.JSON(fiber.Map{
			"error":   true,
			"message": err,
		})
		return
	}
	c.JSON(order)
}

func PostOrderCancelResponseFromPaypal(c *fiber.Ctx) {
	order, err := libraries.GetOrder(c.Params("id"))
	if err != nil {
		c.SendStatus(500)
		c.JSON(fiber.Map{
			"error":   true,
			"message": err,
		})
		return
	}
	p, err := models.GetPaymentByGatewayOrderID(order.ID)
	if err != nil {
		c.SendStatus(500)
		c.JSON(fiber.Map{
			"error":   true,
			"message": err,
		})
		return
	}
	p.PayPalOrderDetail = order
	p.GatewayOrderID = order.ID
	p.UpdatePaymentStatusByGatewayOrderID("CANCELED")
	c.JSON(order)
}

func PostOrderSuccessResponseFromPaypal(c *fiber.Ctx) {
	order, err := libraries.GetOrder(c.Params("id"))
	if err != nil {
		c.SendStatus(500)
		c.JSON(fiber.Map{
			"error":   true,
			"message": err,
		})
		return
	}
	p, err := models.GetPaymentByGatewayOrderID(order.ID)
	p.PayPalOrderDetail = order
	p.GatewayOrderID = order.ID
	p.UpdatePaymentStatusByGatewayOrderID("APPROVED")
	user, _ := models.GetUserById(p.UserID)
	user.AddAmount(p.Amount)
	c.JSON(p)
}
