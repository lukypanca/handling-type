package handler

import (
	"tipe-handling/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type ExcludeContractHandler struct {
	Service *service.ExcludeContractService
}

func NewExcludeContractHandler(s *service.ExcludeContractService) *ExcludeContractHandler {
	return &ExcludeContractHandler{
		Service: s,
	}
}

func (h *ExcludeContractHandler) UploadExcludeContract(ctx *gin.Context) {

	// 1. ambil file dari request
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{"error": "file wajib diupload"})
		return
	}

	f, err := file.Open()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()

	// 2. baca excel
	xl, err := excelize.OpenReader(f)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "file bukan excel valid"})
		return
	}

	sheet := xl.GetSheetName(0)

	rows, err := xl.GetRows(sheet)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 3. panggil service
	err = h.Service.InsertExcludeContractFromExcel(ctx.Request.Context(), rows)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "gagal insert data",
			"error":   err.Error(),
		})
		return
	}

	// 4. success
	ctx.JSON(200, gin.H{
		"message": "upload berhasil",
	})
}
