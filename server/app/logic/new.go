package logic

import (
	model "esports/server/app/models"
	"esports/server/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//新闻列表获取

func GetTextList(c *gin.Context) {
	text, err := model.GetTextList()
	if err != nil {
		c.JSON(200, tools.ECode{Message: "资源未找到"})
	}
	c.JSON(200, tools.ECode{Data: text})
}

// 获取新闻内容

func GetTextContent(c *gin.Context) {
	titleUid := c.Query("titleUid")
	fmt.Println("文章uid", titleUid)
	uid, _ := strconv.ParseInt(titleUid, 10, 0)
	text, err := model.GetTextContent(int(uid))
	if err != nil {
		c.JSON(200, tools.ECode{Message: "资源未找到"})
	}
	c.JSON(200, tools.ECode{Data: text})
}

//文章评论获取

//func GetComment(c *gin.Context) {
//	titleUid := c.Query("titleUid")
//	comment, err := model.GetComment(titleUid)
//	if err != nil {
//		c.JSON(200, tools.ECode{Message: "暂时没有评论"})
//	}
//	c.JSON(200, tools.ECode{Data: comment})
//}
//
//// 评论文章或回复
//
//func CommentText(c *gin.Context) {
//	var user model.Comment
//	if err := c.ShouldBind(&user); err != nil {
//		c.JSON(200, tools.ECode{
//			Code:    10001,
//			Message: err.Error(),
//		})
//		return
//	}
//
//}

// 新闻文章保存

func SaveText(c *gin.Context) {
	// 解析前端传递的JSON数据
	var requestData struct {
		UID  int32  `json:"uid"`
		Text string `json:"text"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// 创建数据库记录
	text := model.Text{
		UserUid: requestData.UID,
		Content: requestData.Text,
	}
	if err := model.SaveText(text); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save text"})
		return
	}

	// 返回保存成功的文本给前端
	c.JSON(200, gin.H{"message": "Text saved successfully", "text": text})
}

// 文章上传

func UploadText(c *gin.Context) {
	titleUid := c.Query("titleUid")
	err := model.UploadText(titleUid)
	if err != nil {
		c.JSON(200, tools.ECode{Code: 1, Message: "上传失败，请重试"})
	}
	c.JSON(200, tools.ECode{Message: "上传成功"})
}
