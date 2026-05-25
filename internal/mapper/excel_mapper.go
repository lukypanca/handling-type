package mapper

import (
	"fmt"
	"tipe-handling/internal/helper"
	"tipe-handling/internal/model"

	"github.com/xuri/excelize/v2"
)

func HeaderExcludeContract() []string {
	return []string{
		"CONTRACT_NO",
		"NAMA_NASABAH",
		"OBJECT_GROUP",
		"TIPE_NASABAH",
		"COLLECTION_SCORING",
		"PAYMENT_TYPE",
		"TIPE_PEMBIAYAN",
		"SKEMA_PEMBIAYAN",
		"PENGGOLONGAN_PRODUCT",
		"BANK_PENDANAAN",
		"MARKETING_PROGRAM",
		"TIPE_HANDLING",
		"STATUS_EXCLUDE",
	}
}

func DataExcludeContractMapper(f *excelize.File, sheet string, data []model.CmsArExcludeContract) {
	for rowIndex, d := range data {
		row := rowIndex + 2

		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), helper.GetStr(d.ContractNo))
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), helper.GetStr(d.NamaNasabah))
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), helper.GetStr(d.ObjectGroup))
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), helper.GetStr(d.TipeNasabah))
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), helper.GetStr(d.CollScoring))
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), helper.GetStr(d.PaymentType))
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), helper.GetStr(d.TipePembiayaan))
		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), helper.GetStr(d.SkemaPembiayaan))
		f.SetCellValue(sheet, fmt.Sprintf("I%d", row), helper.GetStr(d.PenggolonganProduct))
		f.SetCellValue(sheet, fmt.Sprintf("J%d", row), helper.GetStr(d.BankPendanaan))
		f.SetCellValue(sheet, fmt.Sprintf("K%d", row), helper.GetStr(d.MarketingProgram))
		f.SetCellValue(sheet, fmt.Sprintf("L%d", row), helper.GetStr(d.TipeHandling))
		f.SetCellValue(sheet, fmt.Sprintf("M%d", row), helper.GetStr(d.StatusExclude))

	}
}

func MapExcludeContractToModel(row []string) model.CmsArExcludeContract {
	return model.CmsArExcludeContract{
		ContractNo:	helper.ParseString(helper.GetRow(row, 0)),
		NamaNasabah: helper.ParseString(helper.GetRow(row, 1)),
		ObjectGroup: helper.ParseString(helper.GetRow(row, 2)),
		TipeNasabah: helper.ParseString(helper.GetRow(row, 3)),
		CollScoring: helper.ParseString(helper.GetRow(row, 4)),
		PaymentType: helper.ParseString(helper.GetRow(row, 5)),
		TipePembiayaan: helper.ParseString(helper.GetRow(row, 6)),
		SkemaPembiayaan: helper.ParseString(helper.GetRow(row, 7)),
		PenggolonganProduct: helper.ParseString(helper.GetRow(row, 8)),
		BankPendanaan: helper.ParseString(helper.GetRow(row, 9)),
		MarketingProgram: helper.ParseString(helper.GetRow(row, 10)),
		TipeHandling: helper.ParseString(helper.GetRow(row, 11)),
		StatusExclude: helper.ParseString(helper.GetRow(row, 12)),

	}
}