package controller

import (
	"github.com/gin-gonic/gin"
	"harbor-image-delete/model"
	"harbor-image-delete/service"
	"net/http"
)

func Projects(c *gin.Context) {
	projectsUrl := "/api/v2.0/projects?"

	// 解析 JSON 数据
	projectsData := new([]model.Projects)

	projectsData, err := service.Projects(projectsUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取 Projects 列表成功",
		"data": projectsData,
	})
}
