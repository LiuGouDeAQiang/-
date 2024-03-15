package router

import (
	"esports/server/app/logic"
	"github.com/gin-gonic/gin"
)

func New() {
	r := gin.Default()
	r.POST("/login", logic.DoLogin)
	r.POST("/create", logic.CreateUser)
	{
		//战队列表
		r.GET("/teamsList", logic.GetTeamsList)
		r.GET("/teamDetails", logic.GetTeamDetails)
		//战队成员
		r.GET("/teamsMember", logic.GetTeamsMember)
		r.GET("/Player", logic.GetPlayer)
	}

	{ //赛事列表:全球总决赛，职业联赛，冠军赛等
		r.GET("/gameList", logic.GetGamesList)
		//每赛季的具体赛程
		r.GET("/gamesCourse", logic.GetGamesCourse)
	}
	{
		//裁判列表
		r.GET("/coach", logic.GetRefereeMes)
	}

	{
		//发送通知
		r.GET("/notice", logic.Notice)
	}
	{
		//上传新闻
		r.POST("/saveText")
	}
	{
		//关于新闻的列表
		r.GET("/textList", logic.GetTextList)
		//关于新闻的列表
		r.GET("/textContent", logic.GetTextContent)
		//获取新闻评论
		r.GET("/commentGet", logic.GetComment)
		//保存文章
		r.POST("/save", logic.SaveText)
		//上传文章
		r.POST("/upload", logic.UploadText)
		//评论功能
		r.POST("/commentPost", logic.CommentText)

	}
	//定时器1；更新赛程信息
	//{
	//	// 定义定时器
	//	timer := time.NewTicker(12 * time.Hour)
	//	go func() {
	//		update.ReptileGame()
	//		for {
	//			select {
	//			case <-timer.C:
	//				update.ReptileGame()
	//			}
	//		}
	//	}()
	//}

	r.Run(":8080")
}
