package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vuluu2k/remember_fullstack/server/model"
	"github.com/vuluu2k/remember_fullstack/server/model/apperrors"
	"github.com/vuluu2k/remember_fullstack/server/model/mocks"
)

func TestSignUp(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Email and Password Required", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), mock.Anything).Return(nil)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email": "",
		})

		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(reqBody))

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		mockUserService.AssertNotCalled(t, "SignUp")
	})

	t.Run("Invalid Email", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), mock.Anything).Return(nil)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    "vuluu04032000@gmail",
			"password": "SuperKeyPass123",
		})

		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(reqBody))

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		mockUserService.AssertNotCalled(t, "SignUp")
	})

	t.Run("Password too short", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), mock.Anything).Return(nil)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    "vuluu04032000@gmail.com",
			"password": "1234",
		})

		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(reqBody))

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		mockUserService.AssertNotCalled(t, "SignUp")
	})

	t.Run("Password too long", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), mock.Anything).Return(nil)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    "vuluu04032000@gmail.com",
			"password": "1234dsfsfsdafdsafafsafdsafsafsfafafsafdsafsafsdfdfsfadfsfsdafdfasfsfads",
		})

		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(reqBody))

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		mockUserService.AssertNotCalled(t, "SignUp")
	})

	t.Run("Error returned from UserService", func(t *testing.T) {
		u := &model.User{
			Email:    "vuluu040320@gmail.com",
			Password: "SuperKeyPass123",
		}

		mockUserService := new(mocks.MockUserService)

		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), u).Return(apperrors.NewConflict("User Already Exists", u.Email))

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    "vuluu040320@gmail.com",
			"password": "SuperKeyPass123",
		})

		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(reqBody))

		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusConflict, rr.Code)

		mockUserService.AssertExpectations(t)
	})
}
