package customerr

func GetInvalidTimeError() *APIError {
	return &APIError{
		StatusCode: 400,
		Err:        "invalid_time",
		Message:    "Invalid time format",
	}
}

func GetTigersRepoError() *APIError {
	return &APIError{
		StatusCode: 500,
		Err:        "internal_server_error",
		Message:    "Internal server error occurred",
	}
}

func GetInvalidCursorError() *APIError {
	return &APIError{
		StatusCode: 400,
		Err:        "invalid_cursor",
		Message:    "Invalid cursor",
	}
}