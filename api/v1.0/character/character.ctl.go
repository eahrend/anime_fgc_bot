package character

import (
	"context"
	"encoding/json"
	"github.com/eahrend/anime_fgc_bot/common/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
)

func addCharacter(c *gin.Context) {
	db := c.MustGet("mongo").(*mongo.Client)
	characterData := models.StriveCharacterData{}
	err := json.NewDecoder(c.Request.Body).Decode(&characterData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": "failed to decode object",
		})
		return
	}
	characterName := characterData.CharacterNameSimplified
	gameName := characterData.GameNameSimplified
	collection := db.Database("fighters").Collection(gameName)
	res := collection.FindOneAndReplace(context.Background(), bson.M{"character_name_simplified": characterName}, &characterData)
	if res.Err() == mongo.ErrNoDocuments {
		_, err = collection.InsertOne(context.Background(), &characterData)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, map[string]string{
			"status": "success",
		})
		return
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
	return
}
func getCharacterNormal(c *gin.Context) {
	characterName := c.Param("name")
	gameName := c.Param("game")
	characterNormal := c.Param("normal")
	if characterName == "" || gameName == "" || characterNormal == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": "failed to get character or game name",
		})
		return
	}
	db := c.MustGet("mongo").(*mongo.Client)
	collection := db.Database("fighters").Collection(gameName)
	characterData := models.StriveCharacterData{}
	err := collection.FindOne(context.Background(), bson.M{"character_name_simplified": characterName}).Decode(&characterData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": "failed to get character or game name",
		})
		return
	}
	c.JSON(http.StatusOK, characterData.NormalMoves[strings.ToLower(characterNormal)])
}

func getCharacter(c *gin.Context) {
	characterName := c.Param("name")
	gameName := c.Param("game")
	if characterName == "" || gameName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": "failed to get character or game name",
		})
		return
	}
	db := c.MustGet("mongo").(*mongo.Client)
	collection := db.Database("fighters").Collection(gameName)
	characterData := models.StriveCharacterData{}
	err := collection.FindOne(context.Background(), bson.M{"character_name_simplified": characterName}).Decode(&characterData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": "failed to get character or game name",
		})
		return
	}
	c.JSON(http.StatusOK, characterData)
}
