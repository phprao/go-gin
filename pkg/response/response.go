package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const RESPONSE_CODE_SUCCESS = 200
const RESPONSE_CODE_NOT_FOUND = 404
const RESPONSE_CODE_SYSTEM = 500
const RESPONSE_CODE_ERROR = 5000

func JsonResponseError(c *gin.Context, msg string, code int) {
	result := make(map[string]interface{})
	result["code"] = code
	result["msg"] = msg
	result["data"] = ""

	c.JSON(http.StatusOK, result)
}

func JsonResponseSuccess(c *gin.Context, data interface{}) {
	result := make(map[string]interface{})
	result["code"] = 0
	result["msg"] = ""
	result["data"] = data

	c.JSON(http.StatusOK, result)
}
