package controller

import (
	"fmt"
	"movie-review/model"

	"github.com/gin-gonic/gin"
)

func writeResponse(c *gin.Context, data interface{}, err error) {
	var res model.Response
	if err != nil {
		res = model.Response{
			Code:    400,
			Message: fmt.Sprint(err),
			Data:    nil,
		}
	}

	if err == nil {
		res = model.Response{
			Code:    200,
			Message: "Ok",
			Data:    data,
		}
	}

	c.JSON(res.Code, res)
}
