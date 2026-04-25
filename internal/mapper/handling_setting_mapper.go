package mapper

import (
	dto "tipe-handling/internal/dto/request"
	"tipe-handling/internal/model"
)

func ToModel(req *dto.CreateHandlingSettingRequest) *model.HandlingSetting {
	return &model.HandlingSetting{
		DescHandling: req.DescHandling,
		TipeHandling: req.TipeHandling,
		StartOD:      req.StartOD,
		EndOD:        req.EndOD,
	}
}
