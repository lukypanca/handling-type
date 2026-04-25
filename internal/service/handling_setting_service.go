package service

import (
	"context"
	dto "tipe-handling/internal/dto/request"
	"tipe-handling/internal/mapper"
	"tipe-handling/internal/model"
	"tipe-handling/internal/repository"
)

type HandlingSettingService struct {
	amRepo *repository.HandlingSettingRepository
}

func NewHandlingSettingService(repo *repository.HandlingSettingRepository) *HandlingSettingService {
	return &HandlingSettingService{
		amRepo: repo,
	}
}

func (s *HandlingSettingService) GetAll(ctx context.Context) ([]model.HandlingSetting, error) {
	return s.amRepo.FindAll(ctx)
}

func (s *HandlingSettingService) Create(
	ctx context.Context,
	req *dto.CreateHandlingSettingRequest,
) (*model.HandlingSetting, error) {

	data := mapper.ToModel(req)

	id, err := s.amRepo.Save(ctx, data)
	if err != nil {
		return nil, err
	}

	data.ID = id
	return data, nil
}
