package logic

import (
	model "esports/server/app/models"
	"esports/server/app/tools"
	"github.com/gin-gonic/gin"
)

func GetTeamsList(c *gin.Context) {
	teams, err := model.GetTeamsList()
	if err != nil {
		c.JSON(200, tools.ECode{
			Message: "战队列表获取失败",
		})
		return
	}

	c.JSON(200, gin.H{"teams": teams})
}
func GetTeamDetails(c *gin.Context) {
	name := c.Query("team")
	details, err := model.GetTeamDetails(name)
	if err != nil {
		c.JSON(200, tools.ECode{
			Message: "战队详情",
		})
		return
	}
	c.JSON(200, tools.ECode{Data: details})
}

func GetTeamsMember(c *gin.Context) {
	title := c.Query("team")

	teams, err := model.GetTeamPlayers(title)
	if err != nil {
		c.JSON(200, tools.ECode{
			Message: "战队成员获取失败",
		})
		return
	}
	c.JSON(200, tools.ECode{Data: teams})
}
func GetPlayer(c *gin.Context) {

	teams, err := model.GetPlayers()
	if err != nil {
		c.JSON(200, tools.ECode{
			Message: "战队成员获取失败",
		})
		return
	}
	c.JSON(200, tools.ECode{Data: teams})
}
