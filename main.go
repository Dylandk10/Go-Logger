package main

import (
	"log"
	"os"
	//"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"example.test/structs"
	"example.test/testRoute"

)

func getChat(c *gin.Context) {
	
}


//main to set path varibales from env and activate the listener
func main() {

	//load the file paths
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading env file")
	}

	errorPathFromEnv := os.Getenv("Error_Path")
	userPathFromEnv := os.Getenv("User_Path")

	//test the routes
	filepath := structs.Filepath{ErrorPath: errorPathFromEnv, UserPath: userPathFromEnv}
	testRoute.TestingFunc(&filepath, "testingString 2")




	//build the rest api
	router := gin.Default()
	//test path 
	router.GET("/test", testRoute.GetTestRoute)

	//real paths 
	router.GET("/chat", getChat)
	router.Run("localhost:8080")
	fmt.Println("Server running on port: 8080")
}
