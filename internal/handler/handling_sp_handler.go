package handler

import (
	"net/http"
	"tipe-handling/internal/dto/request"
	"tipe-handling/internal/dto/response"
	"tipe-handling/internal/service"

	"github.com/gin-gonic/gin"
)

type HandlingSpHandler struct {
	Service      *service.HandlingSpSptService
}

func NewHandlingSpHandler(s *service.HandlingSpSptService) *HandlingSpHandler {
	return &HandlingSpHandler{
		Service:      s,
	}
}


func (h *HandlingSpHandler) CreateSp(ctx *gin.Context) {
	var req request.CreateHandlingSpSptRequest

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
				Message: "failed to create handling sp spt",
				Error:   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusCreated,
		response.Response[*response.CreateHandlingSpSptResponse]{
			Success: true,
			Message: "success",
			Data: result,
		},
	)
}



