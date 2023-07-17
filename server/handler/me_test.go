package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/vuluu2k/remember_fullstack/model"
	"github.com/vuluu2k/remember_fullstack/model/mocks"
)

func TestMe(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserResp := &model.User{
			UID:   uid,
			Email: "vuluu040320@gmail.com",
			Name:  "Vũ Lưu",
		}

		mockUserService := new(mocks.MockUserService)

		mockUserService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockUserResp, nil)
		rr := httptest.NewRecorder()

		router := gin.Default()
	})
}
