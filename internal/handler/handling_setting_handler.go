package handler

import (
	"net/http"
	dto "tipe-handling/internal/dto/request"
	"tipe-handling/internal/model"
	"tipe-handling/internal/service"

	"github.com/gin-gonic/gin"
)

type HandlingSettingHandler struct {
	Service *service.HandlingSettingService
}

func NewHandlingSettingHandler(s *service.HandlingSettingService) *HandlingSettingHandler {
	return &HandlingSettingHandler{Service: s}
}

func (h *HandlingSettingHandler) GetAll(ctx *gin.Context) {
	data, err := h.Service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (h *HandlingSettingHandler) Create(ctx *gin.Context) {
	req := new(dto.CreateHandlingSettingRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// mapping DTO → model
	data := &model.HandlingSetting{
		DescHandling: req.DescHandling,
		TipeHandling: req.TipeHandling,
		StartOD:      req.StartOD,
		EndOD:        req.EndOD,
	}

	result, err := h.Service.Create(ctx.Request.Context(), data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}
