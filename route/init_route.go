package route

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	SetupApiRouters(r)
	return r
}
