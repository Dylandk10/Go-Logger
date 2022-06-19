package main 

import (
	"fmt"
	"os"
	"errors"
	"example.test/structs"
)


/*
	Error Types
	http error = 1
	validation error = 2
	file operation error 3 
*/
type errorType struct {
	httpError int
	validationError int
	fileOperationError int
}

func ErrorType() *errorType {
	return &errorType{httpError: 1, validationError: 2, fileOperationError: 3}
}


func HandleError(id int) {
	filePath := structs.Filepath{
		ErrorPath: "/Logs/Error", 
		HttpError: "/http/http-errors.txt",
		ValidationError: "/validation/validation-errors.txt",
		FileOperationError: "/file-operations/file-operation-erros.txt",
	}
	switch(id) {
		case 1:
			path := filePath.ErrorPath + filePath.HttpError
			doesExist := doesFileExist(path)

			if doesExist {
				fmt.Println("The path exist")
			} else {
				fmt.Printf("Path does not exist")
			}
		case 2:
			path := filePath.ErrorPath + filePath.ValidationError
			doesExist := doesFileExist(path)

			if doesExist {
				fmt.Println("The path exist")
			} else {
				fmt.Printf("Path does not exist")
			}
			
		case 3:
			path := filePath.ErrorPath + filePath.FileOperationError
			doesExist := doesFileExist(path)
			
			if doesExist {
				fmt.Println("The path exist")
			} else {
				fmt.Printf("Path does not exist")
			}
	}
}

func doesFileExist(filePathString string) bool {
	if _, err := os.Stat(filePathString); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
