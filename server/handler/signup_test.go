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

	t.Run("Successful Token Creation", func(t *testing.T) {
		u := &model.User{
			Email:    "vuluu040320@gmail.com",
			Password: "SuperKeyPass123",
		}

		mockTokenResp := &model.TokenPair{
			TokenID:      "tokenId",
			RefreshToken: "refreshToken",
		}

		mockUserService := new(mocks.MockUserService)
		mockTokenService := new(mocks.MockTokenService)

		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), u).Return(nil)
		mockTokenService.On("NewPairFromUser", mock.AnythingOfType("*gin.Context"), u, "").Return(mockTokenResp, nil)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:            router,
			UserService:  mockUserService,
			TokenService: mockTokenService,
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

		respBody, err := json.Marshal(gin.H{
			"tokens": mockTokenResp,
		})

		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())

		mockUserService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})

	t.Run("Failed Token Creation", func(t *testing.T) {
		u := &model.User{
			Email:    "vuluu040320@gmail.com",
			Password: "SuperKeyPass123",
		}

		mockErrorResponse := apperrors.NewInternal()
		mockUserService := new(mocks.MockUserService)
		mockTokenService := new(mocks.MockTokenService)

		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), u).Return(nil)
		mockTokenService.On("NewPairFromUser", mock.AnythingOfType("*gin.Context"), u, "").Return(nil, mockErrorResponse)

		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:            router,
			UserService:  mockUserService,
			TokenService: mockTokenService,
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

		respBody, err := json.Marshal(gin.H{
			"error": mockErrorResponse,
		})

		assert.NoError(t, err)

		assert.Equal(t, mockErrorResponse.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())

		mockUserService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})
}
