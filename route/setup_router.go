package route

import (
	"github.com/gin-gonic/gin"
	"harbor-image-delete/controller"
	"net/http"
)

func SetupApiRouters(r *gin.Engine) {
	r.GET("/testapi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "testapi success!",
			"data": nil,
		})
	})
	// GET
	r.GET("/Projects", controller.Projects)
	r.GET("/Repositories", controller.Repositories)
	r.GET("/Artifacts", controller.Artifacts)

	// DELETE
	r.DELETE("/DeleteFromProjectsAndRepositories", controller.DeleteFromProjectsAndRepositories)
	r.DELETE("/DeleteFromProjects", controller.DeleteFromProjects)
}
