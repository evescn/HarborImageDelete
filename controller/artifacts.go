package controller

import (
	"github.com/gin-gonic/gin"
	"harbor-image-delete/model"
	"harbor-image-delete/service"
	"net/http"
)

func Artifacts(c *gin.Context) {
	params := new(model.ArtifactsUrl)
	// 绑定参数
	if err := c.ShouldBind(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 调用 service 方法，获取 Artifacts 列表
	artifactsData, err := service.Artifacts(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	// 统计 Artifacts 列表总数
	total := len(*artifactsData)
	c.JSON(http.StatusOK, gin.H{
		"msg":   "获取 Artifacts 列表成功",
		"data":  artifactsData,
		"total": total,
	})

}
