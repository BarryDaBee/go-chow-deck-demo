package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meal struct {
	mgm.DefaultModel `bson:",inline"`
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Name             string             `json:"name" bson:"name"`
	Description      string             `json:"description" bson:"description"`
	Price            float64            `json:"price" bson:"price"`
	Rating           Rating             `json:"rating" bson:"rating"`
	Category         string             `json:"category" bson:"category"`
}
