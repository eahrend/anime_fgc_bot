package character

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes attaches /artifacts routs
func ApplyRoutes(r *gin.RouterGroup) {
	charactersAPI := r.Group("/character")
	{
		charactersAPI.GET("/:game/:name", getCharacter)
		charactersAPI.GET("/:game/:name/normals/:normal", getCharacterNormal)
		charactersAPI.GET("/:game/:name/normals/", getCharacterNormals)
		charactersAPI.GET("/:game/:name/specials/:special", getCharacterSpecial)
		charactersAPI.GET("/:game/:name/specials/", getCharacterSpecials)
		charactersAPI.GET("/:game/:name/supers/:super", getCharacterSuper)
		charactersAPI.GET("/:game/:name/supers/", getCharacterSupers)
		charactersAPI.GET("/:game/:name/throws/:throw", getCharacterThrow)
		charactersAPI.GET("/:game/:name/throws/", getCharacterThrows)
		charactersAPI.POST("/", addCharacter)
	}
}
