package testRoute

import (
	"log"
	"os"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"example.test/structs"
)
//testing struct and response for testingFunc
type TestResponse struct {
	ID string `json:"id"`
	Response string `json:"Title"`
}

var response = []TestResponse {
	{ID: "1", Response: "Test Response"},
}

//debug the code writing path 
func TestingFunc(filePath *structs.Filepath, s string) {
	if err := os.MkdirAll(filePath.ErrorPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(filePath.ErrorPath + "/Dat.txt")
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

func GetTestRoute(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, response)
}
