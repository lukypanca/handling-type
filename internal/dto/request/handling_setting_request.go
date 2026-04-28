package request

type CreateHandlingSettingRequest struct {
	BranchType   string `json:"branch_type"`
	DescHandling string `json:"desc_handling"`
	TipeHandling string `json:"tipe_handling"`
	StartOD      int    `json:"start_od"`
	EndOD        int    `json:"end_od"`
	Status       string `json:"status"`
	IsActive     int    `json:"is_active"`

	Branches        []BranchRequest          `json:"branches"`
	ObjectGroups    []ObjectGroupRequest     `json:"object_groups"`
	TipeNasabah     []TipeNasabahRequest     `json:"tipe_nasabah"`
	CollScoring     []CollScoringRequest     `json:"coll_scoring"`
	PaymentTypes    []PaymentTypeRequest     `json:"payment_types"`
	TipePembiayaan  []TipePembiayaanRequest  `json:"tipe_pembiayaan"`
	SkemaPembiayaan []SkemaPembiayaanRequest `json:"skema_pembiayaan"`
	GolProduk       []GolProdukRequest       `json:"gol_produk"`
	BankPendanaan   []BankPendanaanRequest   `json:"bank_pendanaan"`
}

type CreateHandlingSpSptRequest struct {
	BranchType   string `json:"branch_type"`
	DescHandling string `json:"desc_handling"`
	TipeHandling string `json:"tipe_handling"`
	StartOD      int    `json:"start_od"`
	EndOD        int    `json:"end_od"`
	FlagRod      string `json:"flag_rod"`
	Status       string `json:"status"`
	IsActive     int    `json:"is_active"`

	Branches        []BranchRequest          `json:"branches"`
	ObjectGroups    []ObjectGroupRequest     `json:"object_groups"`
	TipeNasabah     []TipeNasabahRequest     `json:"tipe_nasabah"`
	CollScoring     []CollScoringRequest     `json:"coll_scoring"`
	PaymentTypes    []PaymentTypeRequest     `json:"payment_types"`
	TipePembiayaan  []TipePembiayaanRequest  `json:"tipe_pembiayaan"`
	SkemaPembiayaan []SkemaPembiayaanRequest `json:"skema_pembiayaan"`
	GolProduk       []GolProdukRequest       `json:"gol_produk"`
	BankPendanaan   []BankPendanaanRequest   `json:"bank_pendanaan"`
}

type BranchRequest struct {
	KodeCabang string `json:"kode_cabang"`
	NamaCabang string `json:"nama_cabang"`
	KodeArea   string `json:"kode_area"`
	NamaArea   string `json:"nama_area"`
}

type ObjectGroupRequest struct {
	ObjectCode  string `json:"object_code"`
	ObjectGroup string `json:"object_group"`
}

type TipeNasabahRequest struct {
	TipeNasabahCode string `json:"tipe_nasabah_code"`
	TipeNasabahDesc string `json:"tipe_nasabah_desc"`
}

type CollScoringRequest struct {
	CollScoringCode string `json:"coll_scoring_code"`
	CollScoringDesc string `json:"coll_scoring_desc"`
}

type PaymentTypeRequest struct {
	PaymentTypeCode string `json:"payment_type_code"`
	PaymentTypeDesc string `json:"payment_type_desc"`
}

type TipePembiayaanRequest struct {
	TipePembiayaanCode string `json:"tipe_pembiayaan_code"`
	TipePembiayaanDesc string `json:"tipe_pembiayaan_desc"`
}

type SkemaPembiayaanRequest struct {
	SkemaPembiayaanCode string `json:"skema_pembiayaan_code"`
	SkemaPembiayaanDesc string `json:"skema_pembiayaan_desc"`
}

type GolProdukRequest struct {
	PenggolonganProductCode string `json:"penggolongan_product_code"`
	PenggolonganProductDesc string `json:"penggolongan_product_desc"`
}

type BankPendanaanRequest struct {
	BankCode string `json:"bank_code"`
	BankDesc string `json:"bank_desc"`
}
