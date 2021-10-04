package character

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes attaches /artifacts routs
func ApplyRoutes(r *gin.RouterGroup) {
	charactersAPI := r.Group("/character")
	{
		charactersAPI.GET("/:game/:name", getCharacter)
	}
}
