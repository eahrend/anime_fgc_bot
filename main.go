package main

import (
	"context"
	"github.com/eahrend/anime_fgc_bot/api"
	"github.com/eahrend/anime_fgc_bot/common/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func main() {
	mongoURL := os.Getenv("MONGODB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	app := gin.Default()
	app.Use(gin.Recovery())
	app.Use(middleware.UseMongo(client))
	api.ApplyRoutes(app)
	panic(app.Run("0.0.0.0:8080"))
}
