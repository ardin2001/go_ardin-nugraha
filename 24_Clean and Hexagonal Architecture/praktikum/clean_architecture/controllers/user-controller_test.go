package controllers

import (
	"clean_hexagonal_architecture/middlewares"
	"clean_hexagonal_architecture/models"
	"clean_hexagonal_architecture/repositories"
	"clean_hexagonal_architecture/services"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = services.NewUserService(userRMock)
	jwtMock   = &middlewares.IjwtSMock{Mock: mock.Mock{}}
	userCTest = NewUserController(userSMock, jwtMock)
)

func TestGetUsersController_Success(t *testing.T) {
	users := []*models.User{
		{
			Email:    "tester123@gmail.com",
			Password: "123",
		},
		{
			Email:    "tester123@gmail.com",
			Password: "123",
		},
	}

	userRMock.Mock.On("GetUsersRepository").Return(users, nil)

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := userCTest.GetUsersController(c)
	assert.Nil(t, err)
}

func TestGetUsersController_Failure(t *testing.T) {
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = services.NewUserService(userRMock)
	userCTest = NewUserController(userSMock, jwtMock)
	userRMock.Mock.On("GetUsersRepository").Return(nil, errors.New("get all users failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := userCTest.GetUsersController(c)
	assert.Nil(t, err)
}

func TestCreateUserController_Success(t *testing.T) {
	user := models.User{
		Email:    "tester123@gmail.com",
		Password: "123",
	}

	userRMock.Mock.On("CreateRepository", user).Return(user, nil)
	jwtMock.Mock.On("CreateJWTToken").Return("token123", nil)

	rec := httptest.NewRecorder()

	userByte, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(userByte))

	req := httptest.NewRequest(http.MethodPost, "/users", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = userCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateUserController_Failure1(t *testing.T) {
	user := models.User{}

	userRMock.Mock.On("CreateRepository", user).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	userByte, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(userByte))

	req := httptest.NewRequest(http.MethodPost, "/users", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = userCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateUserController_Failure2(t *testing.T) {
	user := models.User{}

	userRMock.Mock.On("CreateRepository", user).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPost, "/users", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err := userCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateUserController_Failure3(t *testing.T) {
	jwtMock = &middlewares.IjwtSMock{Mock: mock.Mock{}}
	userCTest = NewUserController(userSMock, jwtMock)
	user := models.User{
		Email:    "wdwqdq@gmail.com",
		Password: "321dd",
	}

	userRMock.Mock.On("CreateRepository", user).Return(user, nil)
	jwtMock.Mock.On("CreateJWTToken").Return(nil, errors.New("error bang"))

	rec := httptest.NewRecorder()

	userByte, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(userByte))

	req := httptest.NewRequest(http.MethodPost, "/users", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = userCTest.CreateController(c)
	assert.Nil(t, err)
}
