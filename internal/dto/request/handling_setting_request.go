package dto

type CreateHandlingSettingRequest struct {
	BranchType          string   `json:"branch_type"`
	BranchList          []string `json:"branch_list"`
	ObjectGroup         []string `json:"object_group"`
	TIpeNasabah         []string `json:"tipe_nasabah"`
	StartOD             int      `json:"start_od"`
	EndOD               int      `json:"end_od"`
	CollectionScoring   []string `json:"collection_scoring"`
	PaymentType         []string `json:"payment_type"`
	TipePembiayaan      []string `json:"tipe_pembiayaan"`
	SkemaPembiayaan     []string `json:"skema_pembiayaan"`
	PenggolonganProduct []string `json:"penggolongan_product"`
	BankPendanaan       []string `json:"bank_pendanaan"`
	TipeHandling        string   `json:"tipe_handling"`
	DescHandling        *string  `json:"desc_handling"`
	Status              string   `json:"status"`
}
