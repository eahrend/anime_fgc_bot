package main

import (
	"github.com/eahrend/anime_fgc_bot/api"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(gin.Recovery())
	api.ApplyRoutes(app)
	panic(app.Run("0.0.0.0:8080"))
}
