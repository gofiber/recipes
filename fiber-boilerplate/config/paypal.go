package config

type PaypalConfiguration struct {
	Paypal_Client_ID string
	Paypal_Secret    string
}

var PaypalConfig *PaypalConfiguration //nolint:gochecknoglobals

func LoadPaypalConfig() {
	loadDefaultPaypalConfig()
	ViperConfig.Unmarshal(&PaypalConfig)
}

func loadDefaultPaypalConfig() {
	ViperConfig.SetDefault("PAYPAL_CLIENT_ID", "smtp.mailtrap.io")
	ViperConfig.SetDefault("PAYPAL_SECRET", "2525")
}
