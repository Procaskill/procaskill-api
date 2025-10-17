package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	StartDate   string             `json:"start_date" bson:"start_date"`
	Duration    string             `json:"duration" bson:"duration"`
	DaysOfWeek  []string           `json:"days_of_week" bson:"days_of_week"`
}