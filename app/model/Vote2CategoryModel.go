package model

import (
	"github.com/jinzhu/gorm"
	"votePlatfom/pkg/utils/mysqlUtil"
)

type Vote2CategoryModel struct {
	ID         int    `json:"id" form:"id" primaryKey:"true"`
	Name       string `json:"name" form:"name" binding:"required"`
	ActivityId int    `json:"activity_id" form:"activity_id" binding:"required"`
	Level      int    `json:"level" form:"level"`
	Sort       int    `json:"sort" form:"sort" binding:"required"`
}

var db *gorm.DB

// 设置表名
func (Vote2CategoryModel) TableName() string {
	return "vote2_category"
}

func (m *Vote2CategoryModel) GetCategoryByActivityId(activityId int) []Vote2CategoryModel {
	db = mysqlUtil.GetMysqlConn()

	list := make([]Vote2CategoryModel, 0)
	db.Where("activity_id = ?", activityId).Find(&list)

	//_, err := redisUtil.Conn().Do("set", "gin-name", "1111111")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//res, err := redisUtil.Conn().Do("get", "gin-name")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("%s\n", res)
	//
	//token := jwtUtil.CreateToken(120, "15919906312")
	//fmt.Println(token)
	//info, _ := jwtUtil.ParseToken(token)
	//fmt.Println(info.UserId, info.Phone, info.ExpiresAt)

	return list
}
