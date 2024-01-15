package service

// go:generate mockgen -source=./service/signup.go -destination=./mocks/mock_signup.go -package=mocks github.com/pranoyk/tiger-sightings/service SignUpUser

import (
	"context"
	"fmt"
	"os"

	"github.com/auth0/go-auth0/authentication"
	"github.com/auth0/go-auth0/authentication/database"
	customerr "github.com/pranoyk/tiger-sightings/custom-err"
	"github.com/pranoyk/tiger-sightings/model"
)

type signUpUser struct {
	domain       string
	clientID     string
	clientSecret string
}

type SignUpUser interface {
	SignUp(ctx context.Context, user *model.SignUpUserRequest) (string, *customerr.APIError)
}

func NewSignUpUser() *signUpUser {
	return &signUpUser{
		domain:       os.Getenv("AUTH0_DOMAIN"),
		clientID:     os.Getenv("AUTH0_CLIENT_ID"),
		clientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
	}
}

func (su *signUpUser) SignUp(ctx context.Context, user *model.SignUpUserRequest) (string, *customerr.APIError) {
	authAPI, err := authentication.New(
		ctx,
		su.domain,
		authentication.WithClientID(su.clientID),
		authentication.WithClientSecret(su.clientSecret),
	)
	if err != nil {
		fmt.Printf("failed to initialize the auth0 authentication API client: %+v", err)
		return "", &customerr.APIError{
			StatusCode: 500,
			Err:        "internal_server_error",
			Message:    "Internal server error occurred",
		}
	}

	userData := database.SignupRequest{
		Connection: "Username-Password-Authentication",
		Username:   user.Username,
		Password:   user.Password,
		Email:      user.Email,
	}

	createdUser, err := authAPI.Database.Signup(ctx, userData)
	if err != nil {
		fmt.Printf("failed to sign user up: %+v", err)
		return "", customerr.GetSignUpError(err)
	}

	return createdUser.ID, nil
}
