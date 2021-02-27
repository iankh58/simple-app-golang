package routers

import(
	"github.com/gin-gonic/gin"
	"../controllers"
)

func InitRouter() *gin.Engine {
	// Set the router as the default one shipped with Gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Setup route group for the API
	api := router.Group("/api")

	api.PUT("/users/process", controllers.Process)

	return router
}