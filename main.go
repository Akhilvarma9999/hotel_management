package main

import (
	"context"
	"flag"
	"learngo/hotel-resevation/api"
	"learngo/hotel-resevation/db"

	"log"

	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = db.DBURI

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	//app.Listen(":5000") this is hardcoded address 5000 so instead
	listenAddr := flag.String("listenAddr", ":5000", "The listen adress of the api server")
	flag.Parse()
	//if make run now it uses default 5000 port, we can do this
	//make build
	//./bin/api --listenAddr :7000

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(config)
	//here suppose our frontend group want our api and we create versin of it
	apiv1 := app.Group("/api/v1")
	apiv1.Put("user/:id", userHandler.HandlePutUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	app.Listen(*listenAddr)

}
