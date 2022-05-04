package domain

type Product struct {
	Code  string `json:"code" bson:"code"`
	Name  string `json:"name" bson:"name"`
	Price string `json:"price" bson:"price"`
}
