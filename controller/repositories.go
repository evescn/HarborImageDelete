package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"harbor-image-delete/model"
	"harbor-image-delete/service"
	"net/http"
)

func Repositories(c *gin.Context) {

	params := new(model.Projects)
	//自动把context中request的请求体参数绑定到params上
	if err := c.ShouldBind(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	repositoriesUrl := fmt.Sprintf("/api/v2.0/projects/%s/repositories?page_size=100", params.Name)

	// 解析 JSON 数据
	repositoriesData := new([]model.Repositories)
	
	repositoriesData, err := service.Repositories(repositoriesUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取 Repositories 列表成功",
		"data": repositoriesData,
	})

}
