package router

import (
	"tipe-handling/internal/handler"

	"github.com/gin-gonic/gin"
)

type EngineWrapper struct {
	*gin.Engine
}

func SetupRoutes(h *handler.HandlingSettingHandler) *EngineWrapper {

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "OK"})
	})

	r.GET("/handling", h.GetAll)
	r.POST("/handling", h.Create)

	return &EngineWrapper{r}
}