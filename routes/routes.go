package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"votePlatfom/app/controllers/video2"
	"votePlatfom/app/controllers/vote2"
	"votePlatfom/app/middlewares"
)

func RegisterRoutes(GRouter *gin.Engine) {
	registerRoutesVideo2(GRouter)
	registerRoutesVote2(GRouter)
	registerRoutesTest(GRouter)
}

func registerRoutesTest(GRouter *gin.Engine) {
	GRouter.GET("/test", func(c *gin.Context) {
		seconds := c.DefaultQuery("sleep", "0")
		secondsInt, _ := strconv.Atoi(seconds)

		time.Sleep(time.Duration(secondsInt) * time.Second)
		//c.String(http.StatusOK, "Welcome Gin Server")

		c.JSON(http.StatusOK, map[string]interface{}{"code": 4000, "msg": "Welcome Gin Server ..."})
	})
}

func registerRoutesVideo2(GRouter *gin.Engine) {
	// 路由组
	userGroup := GRouter.Group("/video2")

	userGroup.GET("/GetActivityInfoById", video2.GetActivityInfoById)
}

func registerRoutesVote2(GRouter *gin.Engine) {
	// 路由组
	userGroup := GRouter.Group("/vote2", middlewares.Vote2StatActive())

	userGroup.GET("/GetActivityInfoById", vote2.GetActivityInfoById)
	userGroup.GET("/TestFunc", vote2.TestFunc)
	userGroup.GET("/GetCategoryByActivityId", vote2.GetCategoryByActivityId)
}