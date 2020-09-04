package video2

import (
	"github.com/gin-gonic/gin"
	"votePlatfom/pkg/response"
)

func GetActivityInfoById(c *gin.Context)  {
	response.Context = c
	response.JsonResponseError("something wrong video2 ...")
}