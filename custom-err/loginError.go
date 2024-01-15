package customerr

import "github.com/auth0/go-auth0/authentication"

func GetLoginError(err error) *APIError {
	apiError := err.(*authentication.Error)
	switch apiError.StatusCode {
		case 403:
			return &APIError{
				StatusCode: 403,
				Err:        "invalid_login",
				Message:    "Username or password is invalid",
			}
		case 500:
			return &APIError{
				StatusCode: 500,
				Err:        "internal_server_error",
				Message:    "Internal server error occurred",
			}
		default:
			return &APIError{
				StatusCode: 500,
				Err:        apiError.Err,
				Message:    "Internal server error occurred",
			}
	}
}