package customerr

func GetCreateTigerRepoError() *APIError {
	return &APIError{
		StatusCode: 500,
		Err:        "internal_server_error",
		Message:    "Internal server error occurred",
	}
}

func GetInvalidTimeError() *APIError {
	return &APIError{
		StatusCode: 400,
		Err:        "invalid_time",
		Message:    "Invalid time format",
	}
}