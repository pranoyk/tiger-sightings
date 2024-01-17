package middleware

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/auth0/go-auth0/management"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userEmail, err := tokenValid(c)
		if err != nil {
			fmt.Printf("error validating token: %+v\n", err)
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Set("email", userEmail)
		c.Next()
	}
}

func tokenValid(c *gin.Context) (string, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		file, err := os.ReadFile("./keys/public.pem")
		if err != nil {
			return nil, fmt.Errorf("error reading public key: %v", err)
		}
		block, _ := pem.Decode(file)
		var cert *x509.Certificate
		cert, err = x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("error parsing certificate: %v", err)
		}
		pub := cert.PublicKey.(*rsa.PublicKey)
		return pub, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		return "", fmt.Errorf("token is invalid")
	}
	userId := claims["sub"].(string)

	return getEmail(userId)
}

func getEmail(userId string) (string, error) {
	managementApi, err := management.New(os.Getenv("AUTH0_DOMAIN"),
		management.WithClientCredentials(context.Background(), os.Getenv("AUTH0_CLIENT_ID"), os.Getenv("AUTH0_CLIENT_SECRET")),
	)
	if err != nil {
		return "", fmt.Errorf("error creating management api client: %v", err)
	}
	user, err := managementApi.User.Read(context.Background(), userId)
	if err != nil {
		return "", fmt.Errorf("error reading user: %v", err)
	}
	return *user.Email, nil
}

func extractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
