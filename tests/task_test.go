package tests

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"procaskill/db"
	"time"
)

func deleteTask() {
	tasks := db.GetCollection("tasks")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := tasks.DeleteOne(ctx, bson.M{"title": "Test Task"})
	if err != nil {
		panic("Error when trying to delete test task")
	}
}