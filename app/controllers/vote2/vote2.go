package vote2

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"votePlatfom/app/model"
	"votePlatfom/pkg/response"
)

func GetActivityInfoById(c *gin.Context) {
	response.Context = c
	response.JsonResponseError("something wrong vote2 ...")
}

func TestFunc(c *gin.Context) {
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	response.Context = c
	response.JsonResponseSuccess(map[string]interface{}{
		"dbDatabase": dbDatabase,
		"dbUsername": dbUsername,
	})
}

func GetCategoryByActivityId(c *gin.Context) {
	id := c.DefaultQuery("activityId", "0")
	idInt, _ := strconv.Atoi(id)

	m := model.Vote2CategoryModel{}
	list := m.GetCategoryByActivityId(idInt)
	response.Context = c
	response.JsonResponseSuccess(list)
}
