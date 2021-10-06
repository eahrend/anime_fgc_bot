package character

import (
	"context"
	"encoding/json"
	"fmt"
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
			"error": fmt.Sprintf("failed to decode object: %s", err.Error()),
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
	return
}

func getCharacterNormals(c *gin.Context) {
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
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"error": "failed to get character or game name",
		})
		return
	}
	c.JSON(http.StatusOK, characterData.NormalMoves)
	return
}

func getCharacterThrows(c *gin.Context) {
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
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"error": "failed to get character or game name",
		})
		return
	}
	c.JSON(http.StatusOK, characterData.Throws)
	return
}

func getCharacterSupers(c *gin.Context) {
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
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"error": "failed to get character or game name",
		})
		return
	}
	c.JSON(http.StatusOK, characterData.SuperMoves)
	return
}

func getCharacterSpecials(c *gin.Context) {
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
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"error": "failed to get character or game name",
		})
		return
	}
	c.JSON(http.StatusOK, characterData.SpecialMoves)
	return
}

func getCharacterThrow(c *gin.Context) {
	characterName := c.Param("name")
	gameName := c.Param("game")
	characterThrow := c.Param("throw")
	if characterName == "" || gameName == "" || characterThrow == "" {
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
	c.JSON(http.StatusOK, characterData.Throws[strings.ToLower(characterThrow)])
}

func getCharacterSuper(c *gin.Context) {
	characterName := c.Param("name")
	gameName := c.Param("game")
	characterSuper := c.Param("super")
	if characterName == "" || gameName == "" || characterSuper == "" {
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
	// check if we were sent the command or the name
	if val, ok := characterData.SuperMoves[strings.ToLower(characterSuper)]; ok {
		c.JSON(http.StatusOK, val)
		return
	} else if !ok {
		// check for the command name
		for _, move := range characterData.SuperMoves {
			if move.Name == strings.ToLower(characterSuper) {
				c.JSON(http.StatusOK, move)
				return
			}
		}
	}

	c.JSON(http.StatusNotFound, map[string]string{
		"error": "move not found",
	})
	return
}

func getCharacterSpecial(c *gin.Context) {
	characterName := c.Param("name")
	gameName := c.Param("game")
	characterSpecial := c.Param("special")
	if characterName == "" || gameName == "" || characterSpecial == "" {
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
	// check if we were sent the command or the name
	if val, ok := characterData.SpecialMoves[strings.ToLower(characterSpecial)]; ok {
		c.JSON(http.StatusOK, val)
		return
	} else if !ok {
		// check for the command name
		for _, move := range characterData.SpecialMoves {
			if move.Name == strings.ToLower(characterSpecial) {
				c.JSON(http.StatusOK, move)
				return
			}
		}
	}

	c.JSON(http.StatusNotFound, map[string]string{
		"error": "move not found",
	})
	return
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
