package servers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Any("/fuck", func(context *gin.Context) {
		println("shit")
	})
	return router
}
