package controller

import (
	"github.com/gin-gonic/gin"
	"harbor-image-delete/service"
	"net/http"
)

func SystemGcSchedule(c *gin.Context) {
	jobStatus, err := service.SystemGcSchedule()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":       "清理磁盘成功",
		"jobstatus": jobStatus,
	})
}
