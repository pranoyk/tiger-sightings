package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	customerr "github.com/pranoyk/tiger-sightings/custom-err"
	"github.com/pranoyk/tiger-sightings/mocks"
	"github.com/pranoyk/tiger-sightings/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockLoginService := mocks.NewMockLogin(mockCtrl)
	loginController := &LoginUserController{
		Service: mockLoginService,
		User:    &model.LoginRequest{},
	}

	mockLoginService.EXPECT().
		LogIn(gomock.Any(), gomock.Any()).
		Return("access_token", nil).
		Times(1)

	router := gin.Default()
	router.POST("/login", loginController.Login)

	jsonStr := []byte(`{"username":"testuser","password":"testpass"}`)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "access_token")
}

func TestLoginInvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockLoginService := mocks.NewMockLogin(mockCtrl)
	loginController := &LoginUserController{
		Service: mockLoginService,
		User:    &model.LoginRequest{},
	}

	router := gin.Default()
	router.POST("/login", loginController.Login)

	jsonStr := []byte(`{"username":"testuser","password":""}`)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLoginServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockLoginService := mocks.NewMockLogin(mockCtrl)
	loginController := &LoginUserController{
		Service: mockLoginService,
		User:    &model.LoginRequest{},
	}

	mockLoginService.EXPECT().
		LogIn(gomock.Any(), gomock.Any()).
		Return("", &customerr.APIError{
			StatusCode: 403,
			Err:        "invalid_login",
			Message:    "Username or password is invalid",
		}).
		Times(1)

	router := gin.Default()
	router.POST("/login", loginController.Login)

	jsonStr := []byte(`{"username":"testuser","password":"testpass"}`)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}
