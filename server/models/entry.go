package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID       primitive.ObjectID `bson:"id`
	Dish     *string            `json:"dish"`
	Desc     *string            `json:"desc"`
	Fat      *float64           `json:"fat"`
	Protein  *float64           `json:"protein"`
	Carbs    *float64           `json:"carbs"`
	Calories *float64           `json:"calories"`
}
