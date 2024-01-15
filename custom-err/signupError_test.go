package customerr

import (
	"testing"

	"github.com/auth0/go-auth0/authentication"
)

func TestGetSignUpError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *APIError
	}{
		{
			name: "TestGetSignUpError",
			args: args{
				err: &authentication.Error{
					StatusCode: 400,
					Err:        "invalid_signup",
					Message:    "Invalid signup",
				},
			},
			want: &APIError{
				StatusCode: 400,
				Err:        "invalid_signup",
				Message:    "User already exists",
			},
		},
		{
			name: "TestGetSignUpError",
			args: args{
				err: &authentication.Error{
					StatusCode: 500,
					Err:        "Internal Server Error",
					Message:    "",
				},
			},
			want: &APIError{
				StatusCode: 500,
				Err:        "internal_server_error",
				Message:    "Internal server error occurred",
			},
		},
		{
			name: "TestGetSignUpError",
			args: args{
				err: &authentication.Error{
					StatusCode: 404,
					Err:        "not_found",
					Message:    "Not found",
				},
			},
			want: &APIError{
				StatusCode: 500,
				Err:        "not_found",
				Message:    "Internal server error occurred",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSignUpError(tt.args.err); got.Error() != tt.want.Error() {
				t.Errorf("GetSignUpError() = %v, want %v", got, tt.want)
			}
		})
	}
}