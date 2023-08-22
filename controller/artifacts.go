package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"harbor-image-delete/model"
	"harbor-image-delete/service"
	"net/http"
)

func Artifacts(c *gin.Context) {

	params := new(model.ArtifactsUrl)
	//自动把context中request的请求体参数绑定到params上
	if err := c.ShouldBind(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	artifactsUrl := fmt.Sprintf("/api/v2.0/projects/%s/repositories/%s/artifacts?page_size=100", params.ProjectName, params.RepositoriesName)

	fmt.Println(artifactsUrl)
	fmt.Println(service.Artifacts(artifactsUrl))
	// 解析 JSON 数据
	//artifactsData := new([][]model.Artifacts)
	//
	//repositoriesData, err := service.Repositories(artifactsUrl)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"msg":  err.Error(),
	//		"data": nil,
	//	})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"msg":  "获取 Repositories 列表成功",
	//	"data": repositoriesData,
	//})

}
