package customerr

import (
	"testing"

	"github.com/auth0/go-auth0/authentication"
)

func TestGetLoginError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *APIError
	}{
		{
			name: "TestGetLoginError_403",
			args: args{
				err: &authentication.Error{
					StatusCode: 403,
				},
			},
			want: &APIError{
				StatusCode: 403,
				Err:        "invalid_login",
				Message:    "Username or password is invalid",
			},
		},
		{
			name: "TestGetLoginError_500",
			args: args{
				err: &authentication.Error{
					StatusCode: 500,
				},
			},
			want: &APIError{
				StatusCode: 500,
				Err:        "internal_server_error",
				Message:    "Internal server error occurred",
			},
		},
		{
			name: "TestGetLoginError_xxx",
			args: args{
				err: &authentication.Error{
					StatusCode: 404,
					Err:        "not_found",
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
			if got := GetLoginError(tt.args.err); got.Error() != tt.want.Error() {
				t.Errorf("GetSignUpError() = %v, want %v", got, tt.want)
			}
		})
	}
}
