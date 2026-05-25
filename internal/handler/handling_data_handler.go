package handler

import (
	"net/http"
	"tipe-handling/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type HandlingDataHandler struct {
	Service *service.HandlingSettingExcelService
}

func NewHandlingDataHandler(s *service.HandlingSettingExcelService) *HandlingDataHandler {
	return &HandlingDataHandler{
		Service: s,
	}
}

func (h *HandlingDataHandler) ExportRoboAiToExcel(ctx *gin.Context) {
	branches := ctx.QueryArray("branch")

	fileBytes, err := h.Service.ExportRoboAiData(ctx.Request.Context(), branches)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Header("Content-Disposition", "attachment; filename=robo_ai.xlsx")
	ctx.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", fileBytes)
}

func (h *HandlingDataHandler) UploadExcel(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file required"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()

	xl, err := excelize.OpenReader(f)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	rows, err := xl.GetRows("Sheet1")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = h.Service.ProcessExcel(c, rows)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

func (h *HandlingDataHandler) ExportExcludeContract(ctx *gin.Context) {
	fileBytes, err := h.Service.ExportExcludeContract(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Header("Content-Disposition", "attachment; filename=EXCLUDE_CONTRACT_PROD.xlsx")
	ctx.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", fileBytes)
}