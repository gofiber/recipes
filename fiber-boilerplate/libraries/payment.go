package libraries

import (
	"fmt"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/config"
	"github.com/itsursujit/fiber-boilerplate/models"
	"github.com/plutov/paypal/v3"
)

type Order struct {
	InvoiceID         string
	Amount            string
	Currency          string
	PayPalOrderDetail *paypal.Order
	OrderUrl          string
}

func ConnectToPaypal() {

	var err error
	if Paypal != nil {
		return
	}
	if config.AppConfig.App_Env == "prod" {
		Paypal, err = paypal.NewClient(config.PaypalConfig.Paypal_Client_ID, config.PaypalConfig.Paypal_Secret, paypal.APIBaseLive)
	} else {
		Paypal, err = paypal.NewClient(config.PaypalConfig.Paypal_Client_ID, config.PaypalConfig.Paypal_Secret, paypal.APIBaseSandBox)
	}
	if err != nil {
		// bootstrap.Send("s.baniya.np@gmail.com", "Payment Service - Paypal not working", "Payment not working", "", "")
		Log.Fatal().Err(err).Msg("Cannot connect to Paypal")
	} else {
		Log.Info().Msg("Paypal Connected")
	}
}

func GetOrder(id string) (*paypal.Order, error) {
	return Paypal.GetOrder(id)
}

func CreateOrder(o *models.Payment, user *models.User) error {
	ConnectToPaypal()

	_, err := Paypal.GetAccessToken()
	if err != nil {
		fmt.Println(err)
	}
	order, err := Paypal.CreateOrder(
		paypal.OrderIntentCapture,
		[]paypal.PurchaseUnitRequest{
			{
				ReferenceID: fmt.Sprintf("%d", o.ID),
				Amount: &paypal.PurchaseUnitAmount{
					Value:    o.Amount,
					Currency: o.Currency,
				},
			},
		},
		&paypal.CreateOrderPayer{
			Name: &paypal.CreateOrderPayerName{
				GivenName: user.FirstName,
				Surname:   user.LastName,
			},
			EmailAddress: user.Email,
		},
		&paypal.ApplicationContext{
			BrandName: config.AppConfig.App_Name,
			ReturnURL: fmt.Sprintf("%s/paypal/response", config.AppConfig.App_Url),
			CancelURL: fmt.Sprintf("%s/paypal/cancel", config.AppConfig.App_Url),
		},
	)
	if err != nil {
		return err
	}
	Paypal.CaptureOrder(order.ID, paypal.CaptureOrderRequest{})
	o.PayPalOrderDetail = order
	o.GatewayOrderID = order.ID
	o.GatewayOrderStatus = order.Status
	o.Status = "PROCESSING"
	return nil
}
