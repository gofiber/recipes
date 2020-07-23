package models

/*LoginResponse : Struct for JWT response*/
type LoginResponse struct {
	Token string `json:"jwt,omitempty"`
}
