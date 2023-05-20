package entity

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone" validate:"e164"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region"`
	Email   string `json:"email" validate:"email"`
}
