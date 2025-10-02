package mailApp

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ConsigmentNote string             `bson:"connote,omitempty" json:"connote,omitempty"`
	SenderName     string             `bson:"sender_name,omitempty" json:"sender_name,omitempty"`
	ReceiverName   string             `bson:"receiver_name,omitempty" json:"receiver_name,omitempty"`
	PhoneNumber    string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	ItemContent    string             `bson:"item_content,omitempty" json:"item_content,omitempty"`
	DeliveryStatus string             `bson:"delivery_status,omitempty" json:"delivery_status,omitempty"`
	CreatedAt      primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt      primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
