package controller

import (
	"github.com/gin-gonic/gin"
	"harbor-image-delete/model"
	"harbor-image-delete/service"
	"net/http"
)

// DeleteFromProjectsAndRepositories 基于 Projects 和 Repositories 删除对应 Repositories 下多余镜像，默认保留最后 20 次
func DeleteFromProjectsAndRepositories(c *gin.Context) {
	params := &model.ArtifactsUrl{
		ProjectName:      "",
		RepositoriesName: "",
		Total:            20,
	}
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

	// 调用 service 方法，删除 Image 数据
	deleteData, err := service.DeleteFromProjectsAndRepositories(params, artifactsData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	// 统计需要删除 Image 总数
	total := len(deleteData)
	c.JSON(http.StatusOK, gin.H{
		"msg":   "删除 Image 成功",
		"data":  deleteData,
		"total": total,
	})

	defer service.SystemGcSchedule()

}

// DeleteFromProjects 基于 Projects 删除对应 Projects 所有 Repositories 多余镜像，默认保留最后 20 次
func DeleteFromProjects(c *gin.Context) {
	params := &model.ProjectsUrl{
		ProjectName: "",
		Total:       20,
	}

	// 绑定参数
	if err := c.ShouldBind(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 发起请求，解析 JSON 数据
	newParams := &model.Projects{
		Name:      params.ProjectName,
		ProjectId: 0,
	}
	repositoriesData, err := service.Repositories(newParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	// 调用 service 方法，删除 Image 数据
	deleteData, err := service.DeleteFromProjects(params, repositoriesData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	total := len(deleteData)
	c.JSON(http.StatusOK, gin.H{
		"msg":   "删除 Projects 下 Repositories 的 Images 成功",
		"data":  deleteData,
		"total": total,
	})

	defer service.SystemGcSchedule()
}

// DeleteALL 删除所有 Projects 下 Repositories 多余镜像，默认保留最后 20 次
func DeleteALL(c *gin.Context) {
	return
}
