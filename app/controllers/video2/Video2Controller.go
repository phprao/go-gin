package video2

import (
	"github.com/gin-gonic/gin"
	"votePlatfom/pkg/response"
)

func GetActivityInfoById(c *gin.Context)  {
	response.JsonResponseError(c, "something wrong video2 ...", response.RESPONSE_CODE_ERROR)
}