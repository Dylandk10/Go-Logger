package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

type Filepath struct {
	errorPath string
}

func testingFunc(filePath Filepath, s string) {
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

}

func main() {

	//initalize bash
	cmd := exec.Command("/bin/sh", "./init.sh")
	cmd.Run()

	//load the file paths
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading env file")
	}

	errorFromEnv := os.Getenv("Error_File")

	filepath := Filepath{errorFromEnv}
	testingFunc(filepath, "testingString")

}
