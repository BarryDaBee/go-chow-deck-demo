package models

type Location struct {
	Latitude  float64 `json:"lat" bson:"lat"`
	Longitude float64 `json:"lng" bson:"lng"`
}
