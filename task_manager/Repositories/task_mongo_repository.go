package repositories

import (
	"context"
	domain "task_manager/Domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskMongoRepository struct {
	Collection *mongo.Collection
}

func (r *TaskMongoRepository) CreateTask(task *domain.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.Collection.InsertOne(ctx, task)
	return err
}

func (r *TaskMongoRepository) GetTaskByID(id string) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var task domain.Task
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = r.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskMongoRepository) GetAllTasks() ([]*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var tasks []*domain.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskMongoRepository) UpdateTask(task *domain.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(task.ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": task})
	return err
}

func (r *TaskMongoRepository) DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
