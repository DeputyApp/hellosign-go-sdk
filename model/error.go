package model

type ErrorResponse struct {
	Error    *Error    `json:"error"`
	Warnings []Warning `json:"warnings"`
}

type Error struct {
	Message string `json:"error_msg"`
	Name    string `json:"error_name"`
}
