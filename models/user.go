package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EmailAddress     string             `json:"email_address" bson:"email_address"`
	FirstName        string             `json:"first_name" bson:"first_name"`
	LastName         string             `json:"last_name" bson:"last_name"`
	Phone            string             `json:"phone" bson:"phone"`
	ReferralCode     string             `json:"referral_code" bson:"referral_code"`
	Password         []byte             `json:"-"`
}
