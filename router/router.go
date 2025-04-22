package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()

	initializeRouters(r)

	r.Run(":8080")
}
