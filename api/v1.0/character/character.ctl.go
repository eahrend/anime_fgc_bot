package character

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eahrend/anime_fgc_bot/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func getCharacter(c *gin.Context) {
	characterName := c.Param("name")
	gameName := c.Param("game")
	if characterName == "" || gameName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": "failed to get character or game name",
		})
		return
	}
	fileName := fmt.Sprintf("./data/%s/%s.json", gameName, characterName)
	b, err := os.ReadFile(fileName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("failed to get character data:%s", err.Error()),
		})
		return
	}
	cd := models.StriveCharacterData{}
	err = json.NewDecoder(bytes.NewBuffer(b)).Decode(&cd)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	c.JSONP(200, cd)
	return
}
