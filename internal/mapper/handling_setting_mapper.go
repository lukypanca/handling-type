package mapper

import (
	"tipe-handling/internal/dto/request"
	dto "tipe-handling/internal/dto/request"
	"tipe-handling/internal/model"
)

func ToHandlingSettingModel(req *dto.CreateHandlingSettingRequest) *model.HandlingSetting {
	return &model.HandlingSetting{
		TipeHandling: req.TipeHandling,
		DescHandling: &req.DescHandling,
		StartOD:      req.StartOD,
		EndOD:        req.EndOD,
		Status:       req.Status,
		IsActive:     req.IsActive,
	}
}

func ToHandlingBranchModel(req request.BranchRequest, handlingID int) *model.HandlingBranch {

	return &model.HandlingBranch{
		HandlingSettingId: handlingID,
		KodeCabang:        req.KodeCabang,
		NamaCabang:        req.NamaCabang,
		KodeArea:          req.KodeArea,
		NamaArea:          req.NamaArea,
	}
}

func ToHandlingObjectModel(req request.ObjectGroupRequest, handlingID int) *model.HandlingObject {

	return &model.HandlingObject{
		HandlingSettingId: handlingID,
		ObjectCode:        req.ObjectCode,
		ObjectGroup:       req.ObjectGroup,
	}
}

func ToHandlingTipeNasabah(req request.TipeNasabahRequest, handlingID int) *model.HandlingTipeNasabah {

	return &model.HandlingTipeNasabah{
		HandlingSettingId: handlingID,
		TipeNasabahCode:   req.TipeNasabahCode,
		TipeNasabahDesc:   req.TipeNasabahDesc,
	}
}

func ToHandlingCollScoring(req request.CollScoringRequest, handlingID int) *model.HandlingCollScoring {

	return &model.HandlingCollScoring{
		HandlingSettingId: handlingID,
		CollScoringCode:   req.CollScoringCode,
		CollScoringDesc:   req.CollScoringDesc,
	}
}

func ToHandlingPaymentType(req request.PaymentTypeRequest, handlingID int) *model.HandlingPaymentType {

	return &model.HandlingPaymentType{
		HandlingSettingId: handlingID,
		PaymentTypeCode:   req.PaymentTypeCode,
		PaymentTypeDesc:   req.PaymentTypeDesc,
	}
}

func ToHandlingTipePembiayaan(req request.TipePembiayaanRequest, handlingID int) *model.HandlingTipePembiayaan {

	return &model.HandlingTipePembiayaan{
		HandlingSettingId:  handlingID,
		TipePembiayaanCode: req.TipePembiayaanCode,
		TipePembiayaanDesc: req.TipePembiayaanDesc,
	}
}

func ToHandlingSkemaPembiayaan(req request.SkemaPembiayaanRequest, handlingID int) *model.HandlingSkemaPembiayaan {

	return &model.HandlingSkemaPembiayaan{
		HandlingSettingId:   handlingID,
		SkemaPembiayaanCode: req.SkemaPembiayaanCode,
		SkemaPembiayaanDesc: req.SkemaPembiayaanDesc,
	}
}

func ToHandlingPenggolonganProduct(req request.GolProdukRequest, handlingID int) *model.HandlingPenggolonganProduct {

	return &model.HandlingPenggolonganProduct{
		HandlingSettingId:       handlingID,
		PenggolonganProductCode: req.PenggolonganProductCode,
		PenggolonganProductDesc: req.PenggolonganProductDesc,
	}
}

func ToHandlingBankPendanaan(req request.BankPendanaanRequest, handlingID int) *model.HandlingBankPendanaan {

	return &model.HandlingBankPendanaan{
		HandlingSettingId: handlingID,
		BankCode:          req.BankCode,
		BankDesc:          req.BankDesc,
	}
}
