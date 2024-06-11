package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)



func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.Run()

}

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
}


type Recipe struct {
	ID string `json: "id"`
	Name string `json: "name"`
	Tags []string `json:"tags"`
	Ingredients []string `json: "ingredients"`
	Instruction []string `json: "instruction"`
	PublishedAt time.Time `json: "publishedAt"`
}


func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindBodyWithJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
			return
	}

	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}


