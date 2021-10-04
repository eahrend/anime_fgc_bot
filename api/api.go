package api

import (
	apiv1 "github.com/eahrend/anime_fgc_bot/api/v1.0"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		apiv1.ApplyRoutes(api)
	}
}
