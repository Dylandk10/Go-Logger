package main

import (
	"log"
	"os"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"example.test/structs"
	"example.test/testRoute"

)


//get the chat message and return it
func postChat(c *gin.Context) {
	id := c.PostForm("ID")
	text := c.PostForm("Text")
	fmt.Println(text)
	if len(id) == 0 || len(text) == 0 {
		log.Fatalf("bad request missing form Data")
		c.String(http.StatusBadRequest, "Error missing post data")
	}

	chatResponse := structs.ChatResponse{ID: id, Text: text}
	c.IndentedJSON(http.StatusCreated, chatResponse)
}


//main to set path varibales from env and activate the listener
func main() {

	//load the file paths
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading env file")
	}

	port := os.Getenv("Port")

	//testing routes
	//testRoute.TestingFunc(&filepath, "testingString 2")




	//build the rest api
	router := gin.Default()
	//test path 
	router.GET("/test", testRoute.GetTestRoute)

	//real paths 
	router.POST("/chat", postChat)
	router.Run("localhost:8080")
	fmt.Printf("Server running on port: %s\n", port)
}
