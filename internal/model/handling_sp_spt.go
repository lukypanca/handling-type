package model

import "time"

type HandlingSpSpt struct {
	ID             int
	TipeHandling   string
	TipeHandlingId int
	DescHandling   *string
	StartOD        int
	EndOD          int
	FlagRod        string
	Status         string
	IsActive       int
	InsertDate     time.Time
	InsertBy       string
	UpdateDate     time.Time
	UpdateBy       string
}

type HandlingBranchSpSpt struct {
	ID              int
	HandlingSpSptId int
	KodeCabang      string
	NamaCabang      string
	KodeArea        string
	NamaArea        string
	InsertDate      time.Time
	InsertBy        string
	UpdateDate      time.Time
	UpdateBy        string
}

type HandlingObjectSpSpt struct {
	ID              int
	HandlingSpSptId int
	ObjectCode      string
	ObjectGroup     string
	InsertDate      time.Time
	InsertBy        string
	UpdateDate      time.Time
	UpdateBy        string
}

type HandlingTipeNasabahSpSpt struct {
	ID              int
	HandlingSpSptId int
	TipeNasabahCode string
	TipeNasabahDesc string
	InsertDate      time.Time
	InsertBy        string
	UpdateDate      time.Time
	UpdateBy        string
}

type HandlingCollScoringSpSpt struct {
	ID              int
	HandlingSpSptId int
	CollScoringCode string
	CollScoringDesc string
	InsertDate      time.Time
	InsertBy        string
	UpdateDate      time.Time
	UpdateBy        string
}

type HandlingPaymentTypeSpSpt struct {
	ID              int
	HandlingSpSptId int
	PaymentTypeCode string
	PaymentTypeDesc string
	InsertDate      time.Time
	InsertBy        string
	UpdateDate      time.Time
	UpdateBy        string
}

type HandlingTipePembiayaanSpSpt struct {
	ID                 int
	HandlingSpSptId    int
	TipePembiayaanCode string
	TipePembiayaanDesc string
	InsertDate         time.Time
	InsertBy           string
	UpdateDate         time.Time
	UpdateBy           string
}

type HandlingSkemaPembiayaanSpSpt struct {
	ID                  int
	HandlingSpSptId     int
	SkemaPembiayaanCode string
	SkemaPembiayaanDesc string
	InsertDate          time.Time
	InsertBy            string
	UpdateDate          time.Time
	UpdateBy            string
}

type HandlingPenggolonganProductSpSpt struct {
	ID                      int
	HandlingSpSptId         int
	PenggolonganProductCode string
	PenggolonganProductDesc string
	InsertDate              time.Time
	InsertBy                string
	UpdateDate              time.Time
	UpdateBy                string
}

type HandlingBankPendanaanSpSpt struct {
	ID              int
	HandlingSpSptId int
	BankCode        string
	BankDesc        string
	InsertDate      time.Time
	InsertBy        string
	UpdateDate      time.Time
	UpdateBy        string
}
