package vote2

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"votePlatfom/app/model"
	"votePlatfom/pkg/response"
)

func GetActivityInfoById(c *gin.Context) {
	response.JsonResponseError(c, "something wrong vote2 ...", response.RESPONSE_CODE_ERROR)
}

func TestFunc(c *gin.Context) {
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")

	response.JsonResponseSuccess(c, map[string]interface{}{
		"dbDatabase": dbDatabase,
		"dbUsername": dbUsername,
	})
}

func GetCategoryByActivityId(c *gin.Context) {
	id := c.DefaultQuery("activityId", "0")
	idInt, _ := strconv.Atoi(id)

	m := model.Vote2CategoryModel{}
	list := m.GetCategoryByActivityId(idInt)

	response.JsonResponseSuccess(c, list)
}
