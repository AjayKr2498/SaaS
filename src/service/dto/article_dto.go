package dto

// ErrorDTO - Display Errors Message
type ErrorDTO struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

// SuccessDTO - Display Success Message
type SuccessDTO struct {
	Message string `json:"message"`
}

// Article
type Article struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Status string `json:"status"`
}

// Articles
type Articles struct {
	Articles []Article `json:"articles"`
	Total    int       `json:"total"`
}
