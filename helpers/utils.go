package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WrapResponse(c *gin.Context, data interface{}, err error) {
	type response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	d := response{
		Status:  "ok",
		Message: "",
		Data:    []struct{}{},
	}
	if data != nil {
		d.Data = data
	}
	if err != nil {
		d.Status = "failed"
		d.Message = err.Error()
	}
	c.JSON(http.StatusOK, d)
}
