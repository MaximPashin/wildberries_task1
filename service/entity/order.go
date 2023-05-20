package entity

import (
	"time"
)

type Order struct {
	ID              string    `json:"order_uid" validate:"required"`
	TrackNum        string    `json:"track_number"`
	Entry           string    `json:"entry"`
	Delivery        Delivery  `json:"delivery" validate:"required,dive"`
	Payment         Payment   `json:"payment" validate:"required,dive"`
	Items           []Item    `json:"items" validate:"required,dive"`
	Locale          string    `json:"locale" validate:"bcp47_language_tag"`
	Sign            string    `json:"internal_signature"`
	CustometID      string    `json:"customer_id"`
	DeliveryService string    `json:"delivery_service"`
	ShardKey        string    `json:"shardkey"`
	SmID            int       `json:"sm_id"`
	DateCreated     time.Time `json:"date_created"`
	OofShard        string    `json:"oof_shard"`
}
