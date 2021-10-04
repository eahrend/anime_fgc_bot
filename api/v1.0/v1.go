package apiv1

import (
	"github.com/eahrend/anime_fgc_bot/api/v1.0/character"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes attaches /v1 routs
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	character.ApplyRoutes(v1)
}
