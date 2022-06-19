package structs

//to show file paths for logging
type Filepath struct {
	ErrorPath string
	HttpError string
	ValidationError string
	FileOperationError string
}

//post chat request
type ChatResponse struct {
	ID string `json:"chatID" form:"ID"`
	Text string `json:"text" form:"Text"`

}
