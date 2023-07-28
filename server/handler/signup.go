package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vuluu2k/remember_fullstack/server/model"
	"github.com/vuluu2k/remember_fullstack/server/model/apperrors"
)

type signUpReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var req signUpReq

	if ok := bindData(c, &req); !ok {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.UserService.SignUp(c, u)

	if err != nil {
		log.Printf("Failed to sign up user: %v \n", err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	tokens, err := h.TokenService.NewPairFromUser(c, u, "")

	if err != nil {
		log.Printf("Failed to sign up user: %v \n", err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}
