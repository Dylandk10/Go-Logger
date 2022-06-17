package main

import (
	"log"
	"os"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Filepath struct {
	errorPath string
}

type TestResponse struct {
	ID string `json:"id"`
	Response string `json:"Title"`
}

var response = []TestResponse {
	{ID: "1", Response: "Test Response"},
}

func testingFunc(filePath Filepath, s string) {
	if err := os.MkdirAll(filePath.errorPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(filePath.errorPath + "/Dat.txt")
	if err != nil {
		log.Fatalf("Fatal error creating file %s\n", err)
	}

	defer f.Close()

	writeData := []byte(s)

	_, err2 := f.Write(writeData)

	if err2 != nil {
		log.Fatalf("Error with writing data %s\n", err2)
	}
	fmt.Println("Ran testFunc")
}

func getTestRoute(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, response)
}

func main() {

	//load the file paths
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading env file")
	}

	errorFromEnv := os.Getenv("Error_File")

	filepath := Filepath{errorFromEnv}
	testingFunc(filepath, "testingString 2")

	//build the rest api
	router := gin.Default()
	router.GET("/test", getTestRoute)
	router.Run("localhost:8080")
}
