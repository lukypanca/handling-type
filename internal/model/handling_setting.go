package model

import "time"

type HandlingSetting struct {
	ID           int
	TipeHandling string
	DescHandling *string
	StartOD      int
	EndOD        int
	Status       string
	IsActive     int
	InsertDate   time.Time
	InsertBy     string
	UpdateDate   time.Time
	UpdateBy     string
}

type HandlingBranch struct {
	ID                int
	HandlingSettingId int
	KodeCabang        string
	NamaCabang        string
	KodeArea          string
	NamaArea          string
	InsertDate        time.Time
	InsertBy          string
	UpdateDate        time.Time
	UpdateBy          string
}

type HandlingObject struct {
	ID                int
	HandlingSettingId int
	ObjectCode        string
	ObjectGroup       string
	InsertDate        time.Time
	InsertBy          string
	UpdateDate        time.Time
	UpdateBy          string
}

type HandlingTipeNasabah struct {
	ID                int
	HandlingSettingId int
	TipeNasabahCode   string
	TipeNasabahDesc   string
	InsertDate        time.Time
	InsertBy          string
	UpdateDate        time.Time
	UpdateBy          string
}

type HandlingCollScoring struct {
	ID                int
	HandlingSettingId int
	CollScoringCode   string
	CollScoringDesc   string
	InsertDate        time.Time
	InsertBy          string
	UpdateDate        time.Time
	UpdateBy          string
}

type HandlingPaymentType struct {
	ID                int
	HandlingSettingId int
	PaymentTypeCode   string
	PaymentTypeDesc   string
	InsertDate        time.Time
	InsertBy          string
	UpdateDate        time.Time
	UpdateBy          string
}

type HandlingTipePembiayaan struct {
	ID                 int
	HandlingSettingId  int
	TipePembiayaanCode string
	TipePembiayaanDesc string
	InsertDate         time.Time
	InsertBy           string
	UpdateDate         time.Time
	UpdateBy           string
}

type HandlingSkemaPembiayaan struct {
	ID                  int
	HandlingSettingId   int
	SkemaPembiayaanCode string
	SkemaPembiayaanDesc string
	InsertDate          time.Time
	InsertBy            string
	UpdateDate          time.Time
	UpdateBy            string
}

type HandlingPenggolonganProduct struct {
	ID                      int
	HandlingSettingId       int
	PenggolonganProductCode string
	PenggolonganProductDesc string
	InsertDate              time.Time
	InsertBy                string
	UpdateDate              time.Time
	UpdateBy                string
}

type HandlingBankPendanaan struct {
	ID                int
	HandlingSettingId int
	BankCode          string
	BankDesc          string
	InsertDate        time.Time
	InsertBy          string
	UpdateDate        time.Time
	UpdateBy          string
}
