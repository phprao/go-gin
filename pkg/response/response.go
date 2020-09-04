package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const RESPONSE_CODE_SUCCESS = 200
const RESPONSE_CODE_NOT_FOUND = 404
const RESPONSE_CODE_SYSTEM = 500
const RESPONSE_CODE_ERROR = 5000

var Context *gin.Context

func JsonResponseError(msg string, code ...int) {
	result := make(map[string]interface{})
	if len(code) == 0 {
		code = append(code, RESPONSE_CODE_ERROR)
	}
	result["code"] = code[0]
	result["msg"] = msg
	result["data"] = ""

	Context.JSON(http.StatusOK, result)
}

func JsonResponseSuccess(data interface{}) {
	result := make(map[string]interface{})
	result["code"] = 0
	result["msg"] = ""
	result["data"] = data

	Context.JSON(http.StatusOK, result)
}
