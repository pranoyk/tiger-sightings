package customerr

import "github.com/auth0/go-auth0/authentication"

func GetSignUpError(err error) *APIError {
	apiError := err.(*authentication.Error)
	switch apiError.StatusCode {
		case 400:
			return &APIError{
				StatusCode: 400,
				Err:        "invalid_signup",
				Message:    "User already exists",
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