package main

import (
	"fmt"
	"regexp"
)

func ValidatePostChat(id string, text string) bool {
	if len(id) <= 0 || len(text) <= 0 {
		return true
	}
	validator := "/[$\\(\\)<>]/"
	res1, res1Error:= regexp.MatchString(id, validator)
	res2, res2Error := regexp.MatchString(text, validator)
	fmt.Println(res1)
	fmt.Println(res2)
	if res1  || res2 || res1Error != nil || res2Error != nil{
		return true
	}
	
	//all text pass return false
	return false
}