package handler

import (
	"net/http"
	"tipe-handling/internal/dto/request"
	"tipe-handling/internal/dto/response"
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

	var req request.CreateHandlingSettingRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,
			response.Response[any]{
				Success: false,
				Message: "invalid request",
				Error:   err.Error(),
			},
		)
		return
	}

	result, err := h.Service.Create(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			response.Response[any]{
				Success: false,
				Message: "failed to create handling setting",
				Error:   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusCreated,
		response.Response[*response.CreateHandlingSettingResponse]{
			Success: true,
			Message: "success",
			Data:    result,
		},
	)
}
