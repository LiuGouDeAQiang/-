package logic

import (
	model "esports/server/app/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 赛程

func GetGamesCourse(c *gin.Context) {
	gameName1 := c.Query("game")

	fmt.Println(gameName1)
	course := model.GetGamesCourse(gameName1)

	c.JSON(200, gin.H{"data": course})
}

// 赛程列表

func GetGamesList(c *gin.Context) {
	List := model.GetGamesList()
	c.JSON(200, gin.H{"data": List})
}

// 赛事裁判

func GetRefereeMes(c *gin.Context) {
	List := model.GetRefereeMes()
	c.JSON(200, gin.H{"data": List})
}
