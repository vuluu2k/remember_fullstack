package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vuluu2k/remember_fullstack/server/model"
)

type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
}

type Config struct {
	R            *gin.Engine
	UserService  model.UserService
	TokenService model.TokenService
}

func NewHandler(c *Config) {
	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	}

	g := c.R.Group(os.Getenv("AUTH_API_URL"))
	g.GET("/me", h.Me)
	g.POST("/sign-up", h.SignUp)
	g.POST("/sign-in", h.SignIn)
	g.POST("/sign-out", h.SignOut)
	g.POST("/token", h.Token)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)
}

func (h *Handler) SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's sign in",
	})
}

func (h *Handler) SignOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's sign out",
	})
}

func (h *Handler) Token(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's token",
	})
}

func (h *Handler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's image",
	})
}

func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's delete image",
	})
}

func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's details",
	})
}
