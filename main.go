package main

import (
	"time"

	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()
	router.Run()

}


type Recipe struct {
	Name string `json: "name"`
	Tags []string `json:"tags"`
	Ingredients []string `json: "ingredients"`
	Instruction []string `json: "instruction"`
	PublishedAt time.Time `json: "publishedAt"`
}