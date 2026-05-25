package router

import (
	"tipe-handling/internal/handler"

	"github.com/gin-gonic/gin"
)

type EngineWrapper struct {
	*gin.Engine
}

func SetupRoutes(h *handler.HandlingSettingHandler,
	he *handler.HandlingDataHandler,
	ec *handler.ExcludeContractHandler,
	hs *handler.HandlingSpHandler) *EngineWrapper {

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "OK"})
	})

	r.GET("/handling", h.GetAll)
	r.POST("/handling", h.Create)
	r.GET("/getHandlingExcel", he.ExportRoboAiToExcel) //http://localhost:64387/getHandlingExcel?branch=0106
	r.GET("/uploadExcel", he.UploadExcel)
	r.GET("/getExcludeContractProd", he.ExportExcludeContract)
	r.GET("/uploadExcludeContractExcel", ec.UploadExcludeContract)
	r.POST("/createSp", hs.CreateSp)


	return &EngineWrapper{r}
}
