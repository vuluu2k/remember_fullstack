package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vuluu2k/remember_fullstack/server/model/apperrors"
)

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func bindData(c *gin.Context, req interface{}) bool {

	if err := c.ShouldBind(req); err != nil {
		log.Printf("Error binding data: %v\n", err)

		if errs, ok := err.(validator.ValidationErrors); ok {
			var invalidArgs []invalidArgument

			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					Field: err.Field(),
					Value: err.Value().(string),
					Tag:   err.Tag(),
					Param: err.Param(),
				})
			}

			err := apperrors.NewBadRequest("Invalid request parameters. See invalidArgs")

			c.JSON(err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})

			return false
		}

		fallback := apperrors.NewInternal()

		c.JSON(fallback.Status(), gin.H{
			"error": fallback,
		})

		return false
	}

	return true
}
