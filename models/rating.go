package models

type Rating struct {
	Score float32 `json:"score" bson:"score"`
	Count int     `json:"count" bson:"count"`
}
