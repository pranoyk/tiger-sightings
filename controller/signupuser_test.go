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

func TestRegisterUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSignUpService := mocks.NewMockSignUpUser(mockCtrl)
	userController := &SignUpUserController{
		Service: mockSignUpService,
		User:    &model.SignUpUserRequest{},
	}

	mockSignUpService.EXPECT().
		SignUp(gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	router := gin.Default()
	router.POST("/register", userController.RegisterUser)

	jsonStr := []byte(`{"username":"testuser","password":"testpass","email":"abc@def.com"}`)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user successfully created!")
}

func TestRegisterUserMissingFieldsError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSignUpService := mocks.NewMockSignUpUser(mockCtrl)
	userController := &SignUpUserController{
		Service: mockSignUpService,
		User:    &model.SignUpUserRequest{},
	}

	router := gin.Default()
	router.POST("/register", userController.RegisterUser)

	jsonStr := []byte(`{"username":"testuser","password":"","email":"abc@def.com"}`)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRegisterUserServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSignUpService := mocks.NewMockSignUpUser(mockCtrl)
	userController := &SignUpUserController{
		Service: mockSignUpService,
		User:    &model.SignUpUserRequest{},
	}

	mockSignUpService.EXPECT().
		SignUp(gomock.Any(), gomock.Any()).
		Return(&customerr.APIError{
			StatusCode: 400,
			Err:        "invalid_signup",
			Message:    "User already exists",
		}).
		Times(1)

	router := gin.Default()
	router.POST("/register", userController.RegisterUser)

	jsonStr := []byte(`{"username":"testuser","password":"testpass","email":"abc@def.com"}`)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "User already exists")
}
