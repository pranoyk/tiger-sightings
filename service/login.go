package service

// go:generate mockgen -source=./service/login.go -destination=./mocks/mock_login.go -package=mocks github.com/pranoyk/tiger-sightings/service Login

import (
	"context"
	"fmt"
	"os"

	"github.com/auth0/go-auth0/authentication"
	"github.com/auth0/go-auth0/authentication/oauth"
	customerr "github.com/pranoyk/tiger-sightings/custom-err"
	"github.com/pranoyk/tiger-sightings/model"
)

type login struct {
	domain       string
	clientID     string
	clientSecret string
	audience     string
}

type Login interface {
	LogIn(ctx context.Context, user *model.LoginRequest) (string, *customerr.APIError)
}

func NewLogin() Login {
	return &login{
		domain:       os.Getenv("AUTH0_DOMAIN"),
		clientID:     os.Getenv("AUTH0_CLIENT_ID"),
		clientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		audience:     os.Getenv("AUTH0_AUDIENCE"),
	}
}

func (l *login) LogIn(ctx context.Context, user *model.LoginRequest) (string, *customerr.APIError) {
	authAPI, err := authentication.New(
		ctx,
		l.domain,
		authentication.WithClientID(l.clientID),
		authentication.WithClientSecret(l.clientSecret),
	)
	if err != nil {
		fmt.Printf("failed to initialize the auth0 authentication API client: %+v \n", err)
		return "", &customerr.APIError{
			StatusCode: 500,
			Err:        "internal_server_error",
			Message:    "Internal server error occurred",
		}
	}

	tokenSet, err := authAPI.OAuth.LoginWithPassword(ctx, oauth.LoginWithPasswordRequest{
		ClientAuthentication: oauth.ClientAuthentication{
			ClientID:    l.clientID,
			ClientSecret: l.clientSecret,
		},
		Username:             user.Username,
		Password:             user.Password,
		Audience:             l.audience,
	}, oauth.IDTokenValidationOptions{})
	if err != nil {
		fmt.Printf("failed to login user: %+v \n", err)
		return "", customerr.GetLoginError(err)
	}

	fmt.Printf("tokenset: %+v \n", tokenSet)

	return tokenSet.AccessToken, nil
}
