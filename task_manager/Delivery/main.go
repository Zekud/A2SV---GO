package main

import (
	"context"
	"log"
	"os"
	"task_manager/Delivery/controllers"
	"task_manager/Delivery/routers"
	repositories "task_manager/Repositories"
	usecases "task_manager/Usecases"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(dbName)
	taskRepo := &repositories.TaskMongoRepository{Collection: db.Collection("tasks")}
	userRepo := &repositories.UserMongoRepository{Collection: db.Collection("users")}
	taskUsecase := &usecases.TaskUsecase{TaskRepo: taskRepo}
	userUsecase := &usecases.UserUsecase{UserRepo: userRepo}
	taskController := &controllers.TaskController{TaskUsecase: taskUsecase}
	userController := &controllers.UserController{UserUsecase: userUsecase}
	r := routers.SetupRouter(taskController, userController)
	r.Run(":8080")
}
