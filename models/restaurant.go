package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	mgm.DefaultModel `bson:",inline"`
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Location         Location           `json:"location,inline" bson:"location,inline"`
	Menu             []Meal             `json:"menu" bson:"menu"`
	Tags             []string           `json:"tags" bson:"tags"`
}

func (r *Restaurant) Create() error {
	err := mgm.Coll(r).Create(r)

	if err != nil {
		return err
	}

	return nil
}
