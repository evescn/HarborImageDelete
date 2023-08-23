package controller

import (
	"github.com/gin-gonic/gin"
	"harbor-image-delete/service"
	"net/http"
)

func Projects(c *gin.Context) {
	projectsUrl := "/api/v2.0/projects?"

	// 调用 service 方法，获取 Projects 列表
	projectsData, err := service.Projects(projectsUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	// 统计 Projects 列表总数
	total := len(*projectsData)
	c.JSON(http.StatusOK, gin.H{
		"msg":   "获取 Projects 列表成功",
		"data":  projectsData,
		"total": total,
	})
}
