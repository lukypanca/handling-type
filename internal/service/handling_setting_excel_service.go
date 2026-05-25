package service

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
	"tipe-handling/internal/helper"
	"tipe-handling/internal/mapper"
	"tipe-handling/internal/model"
	"tipe-handling/internal/repository/am"

	"github.com/xuri/excelize/v2"
)

type HandlingSettingExcelService struct {
	db     *sql.DB
	amrepo *am.HandlingSettingExcelRepository
}

func NewHandlingSettingExcelService(db *sql.DB, amrepo *am.HandlingSettingExcelRepository) *HandlingSettingExcelService {
	return &HandlingSettingExcelService{
		db:     db,
		amrepo: amrepo,
	}
}

func (s *HandlingSettingExcelService) ExportRoboAiData(ctx context.Context, branches []string) ([]byte, error) {

	data, err := s.amrepo.GetRoboAiData(ctx, branches)
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheet := "Sheet1"

	// header
	headers := []string{
		"AGING_DATE",
		"TGL_PROSES_AGING_DATE",
		"TGL_PPD",
		"MONTH_PPD_DATE",
		"YEAR_PPD_DATE",
		"TGL_PENGAJUAN_RESCHEDULE",
		"HARI_RESCHEDULE",
		"NAMA_AREA",
		"KODE_AREA",
		"NAMA_CABANG",
		"KODE_CABANG",
		"MUFNET",
		"KODE_MUFNET",
		"FINANCIAL_TYPE",
		"CONTRACTNO",
		"CUSTNAME",
		"PEKERJAAN",
		"PASANGAN",
		"DEALER",
		"MERK",
		"JENIS",
		"MODEL",
		"TAHUN_KENDARAAN",
		"NOPOL",
		"RANGKA",
		"MESIN",
		"STATUS_JF",
		"OBJECT",
		"PORTFOLIO",
		"PENGGOLONGAN_PRODUCT",
		"SALES_THROUGH",
		"MARKETING_INTERNAL",
		"MARKETING_HEAD",
		"INT_NAME",
		"CFO_NAME",
		"NAME_CA",
		"OTR",
		"EFF_RATE",
		"NETT_DP",
		"KE",
		"TENOR",
		"JATUH_TEMPO",
		"STATUS_KONTRAK",
		"HARI",
		"BUCKET_AWAL",
		"BUCKET_AKHIR",
		"PEMBAYARAN_M1",
		"PEMBAYARAN_M2",
		"PEMBAYARAN_M3",
		"PEMBAYARAN_TERAKHIR",
		"PAYMENT_TYPE",
		"APPL_PRINCIPAL_AMT",
		"ANGSURAN",
		"OST_DENDA",
		"TOTAL_TITIPAN",
		"BAL_INTR",
		"BAL_PRIN",
		"NO_OID",
		"OD_OID",
		"TIPE_RESTRUCTURE",
		"NO_CONTRACT_OLD",
		"BANK_PENDANAAN",
		"FLAG_AGING",
		"TGL_PROSES",
		"SETTLEMENT_NO",
		"CHANNEL",
		"NAMA_PRODUCT",
		"NAMA_REF_MITRA",
		"LAPANGAN_USAHA",
		"DUEINSTALL",
		"INTERNAL_NPK",
		"POSISI",
		"TAMBAHAN_TENOR",
		"TANGGAL_TARIK",
		"NAMA_PIC_TARIK",
	}

	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// helper
	getStr := func(s *string) string {
		if s != nil {
			return *s
		}
		return ""
	}

	getInt := func(i *int) interface{} {
		if i != nil {
			return *i
		}
		return ""
	}

	getTime := func(t *time.Time) interface{} {
		if t != nil {
			return t.Format("2006-01-02")
		}
		return ""
	}

	getFloat := func(f *float64) interface{} {
		if f != nil {
			return *f
		}
		return nil
	}


	// isi data
	for rowIndex, d := range data {
		row := rowIndex + 2

		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), getTime(d.AgingDate))
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), getTime(d.TglProsesAging))
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), getTime(d.TglPpd))
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), getInt(d.MonthPpdDate))
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), getStr(d.YearPpdDate))
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), getTime(d.TglPengajuanReschedule))
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), getStr(d.HariReschedule))
		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), getStr(d.NamaArea))
		f.SetCellValue(sheet, fmt.Sprintf("I%d", row), getStr(d.KodeArea))
		f.SetCellValue(sheet, fmt.Sprintf("J%d", row), getStr(d.NamaCabang))
		f.SetCellValue(sheet, fmt.Sprintf("K%d", row), getStr(d.KodeCabang))
		f.SetCellValue(sheet, fmt.Sprintf("L%d", row), getStr(d.Mufnet))
		f.SetCellValue(sheet, fmt.Sprintf("M%d", row), getStr(d.KodeMufnet))
		f.SetCellValue(sheet, fmt.Sprintf("N%d", row), getStr(d.FinancialType))
		f.SetCellValue(sheet, fmt.Sprintf("O%d", row), getStr(d.ContractNo))
		f.SetCellValue(sheet, fmt.Sprintf("P%d", row), getStr(d.CustName))
		f.SetCellValue(sheet, fmt.Sprintf("Q%d", row), getStr(d.Pekerjaan))
		f.SetCellValue(sheet, fmt.Sprintf("R%d", row), getStr(d.Pasangan))
		f.SetCellValue(sheet, fmt.Sprintf("S%d", row), getStr(d.Dealer))
		f.SetCellValue(sheet, fmt.Sprintf("T%d", row), getStr(d.Merk))
		f.SetCellValue(sheet, fmt.Sprintf("U%d", row), getStr(d.Jenis))
		f.SetCellValue(sheet, fmt.Sprintf("V%d", row), getStr(d.Model))
		f.SetCellValue(sheet, fmt.Sprintf("W%d", row), getStr(d.TahunKendaraan))
		f.SetCellValue(sheet, fmt.Sprintf("X%d", row), getStr(d.Nopol))
		f.SetCellValue(sheet, fmt.Sprintf("Y%d", row), getStr(d.Rangka))
		f.SetCellValue(sheet, fmt.Sprintf("AA%d", row), getStr(d.Mesin))
		f.SetCellValue(sheet, fmt.Sprintf("AB%d", row), getStr(d.StatusJf))
		f.SetCellValue(sheet, fmt.Sprintf("AC%d", row), getStr(d.Object))
		f.SetCellValue(sheet, fmt.Sprintf("AD%d", row), getStr(d.Portfolio))
		f.SetCellValue(sheet, fmt.Sprintf("AE%d", row), getStr(d.PenggolonganProduct))
		f.SetCellValue(sheet, fmt.Sprintf("AF%d", row), getStr(d.SalesThrough))
		f.SetCellValue(sheet, fmt.Sprintf("AG%d", row), getStr(d.MarketingInternal))
		f.SetCellValue(sheet, fmt.Sprintf("AH%d", row), getStr(d.MarketingHead))
		f.SetCellValue(sheet, fmt.Sprintf("AI%d", row), getStr(d.IntName))
		f.SetCellValue(sheet, fmt.Sprintf("AJ%d", row), getStr(d.CfoName))
		f.SetCellValue(sheet, fmt.Sprintf("AK%d", row), getStr(d.NameCa))
		f.SetCellValue(sheet, fmt.Sprintf("AL%d", row), getFloat(d.Otr))
		f.SetCellValue(sheet, fmt.Sprintf("AM%d", row), getFloat(d.EffRate))
		f.SetCellValue(sheet, fmt.Sprintf("AN%d", row), getFloat(d.NettDp))
		f.SetCellValue(sheet, fmt.Sprintf("AO%d", row), getInt(d.Ke))
		f.SetCellValue(sheet, fmt.Sprintf("AP%d", row), getInt(d.Tenor))
		f.SetCellValue(sheet, fmt.Sprintf("AQ%d", row), getTime(d.JatuhTempo))
		f.SetCellValue(sheet, fmt.Sprintf("AR%d", row), getStr(d.StatusKontrak))
		f.SetCellValue(sheet, fmt.Sprintf("AS%d", row), getInt(d.Hari))
		f.SetCellValue(sheet, fmt.Sprintf("AT%d", row), getStr(d.BucketAwal))
		f.SetCellValue(sheet, fmt.Sprintf("AU%d", row), getStr(d.BucketAkhir))
		f.SetCellValue(sheet, fmt.Sprintf("AV%d", row), getStr(d.PembayaranM1))
		f.SetCellValue(sheet, fmt.Sprintf("AW%d", row), getStr(d.PembayaranM2))
		f.SetCellValue(sheet, fmt.Sprintf("AX%d", row), getStr(d.PembayaranM3))
		f.SetCellValue(sheet, fmt.Sprintf("AY%d", row), getTime(d.PembayaranTerakhir))
		f.SetCellValue(sheet, fmt.Sprintf("AZ%d", row), getStr(d.PaymentType))
		f.SetCellValue(sheet, fmt.Sprintf("BA%d", row), getFloat(d.ApplPrincipalAmt))
		f.SetCellValue(sheet, fmt.Sprintf("BB%d", row), getFloat(d.Angsuran))
		f.SetCellValue(sheet, fmt.Sprintf("BC%d", row), getFloat(d.OstDenda))
		f.SetCellValue(sheet, fmt.Sprintf("BD%d", row), getFloat(d.TotalTitipan))
		f.SetCellValue(sheet, fmt.Sprintf("BE%d", row), getFloat(d.BalIntr))
		f.SetCellValue(sheet, fmt.Sprintf("BF%d", row), getFloat(d.BalPrin))
		f.SetCellValue(sheet, fmt.Sprintf("BG%d", row), getStr(d.NoOid))
		f.SetCellValue(sheet, fmt.Sprintf("BH%d", row), getInt(d.OdOid))
		f.SetCellValue(sheet, fmt.Sprintf("BI%d", row), getStr(d.TipeRestructure))
		f.SetCellValue(sheet, fmt.Sprintf("BJ%d", row), getStr(d.NoContractOld))
		f.SetCellValue(sheet, fmt.Sprintf("BK%d", row), getStr(d.BankPendanaan))
		f.SetCellValue(sheet, fmt.Sprintf("BL%d", row), getInt(d.FlagAging))
		f.SetCellValue(sheet, fmt.Sprintf("BM%d", row), getTime(d.TglProses))
		f.SetCellValue(sheet, fmt.Sprintf("BN%d", row), getStr(d.SettlementNo))
		f.SetCellValue(sheet, fmt.Sprintf("BO%d", row), getStr(d.Channel))
		f.SetCellValue(sheet, fmt.Sprintf("BP%d", row), getStr(d.NamaProduct))
		f.SetCellValue(sheet, fmt.Sprintf("BQ%d", row), getStr(d.NamaRefMitra))
		f.SetCellValue(sheet, fmt.Sprintf("BR%d", row), getStr(d.LapanganUsaha))
		f.SetCellValue(sheet, fmt.Sprintf("BS%d", row), getInt(d.DueInstall))
		f.SetCellValue(sheet, fmt.Sprintf("BT%d", row), getStr(d.InternalNpk))
		f.SetCellValue(sheet, fmt.Sprintf("BU%d", row), getInt(d.Posisi))
		f.SetCellValue(sheet, fmt.Sprintf("BV%d", row), getInt(d.TambahanTenor))
		f.SetCellValue(sheet, fmt.Sprintf("BW%d", row), getTime(d.TanggalTarik))
		f.SetCellValue(sheet, fmt.Sprintf("BX%d", row), getStr(d.NamaPicTarik))
	}

	// simpan ke buffer
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}

func getRow(row []string, i int) string {
	if i < len(row) {
		return strings.TrimSpace(row[i])
	}
	return ""
}

func cleanNumber(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ".", "")

	if s == "" || s == "-" {
		return "0"
	}
	return s
}

func isRowEmpty(row []string) bool {
	for _, v := range row {
		if strings.TrimSpace(v) != "" {
			return false
		}
	}
	return true
}

func (s *HandlingSettingExcelService) ProcessExcel(ctx context.Context, rows [][]string) error {

	var (
		data         []model.CmsRptAgingMis
		skippedRow   int
		totalRow     int
	)

	for i, row := range rows {

		if i == 0 {
			continue // header
		}

		if len(row) == 0 || isRowEmpty(row) {
			continue
		}

		totalRow++

		d := model.CmsRptAgingMis{
			AgingDate:              helper.ParseTime(getRow(row, 0)),
			TglProsesAging:        helper.ParseTime(getRow(row, 1)),
			TglPpd:                helper.ParseTime(getRow(row, 2)),
			MonthPpdDate:          helper.ParseInt(cleanNumber(getRow(row, 3))),
			YearPpdDate:           helper.ParseString(getRow(row, 4)),
			TglPengajuanReschedule:helper.ParseTime(getRow(row, 5)),
			HariReschedule:        helper.ParseString(getRow(row, 6)),
			NamaArea:              helper.ParseString(getRow(row, 7)),
			KodeArea:              helper.ParseString(getRow(row, 8)),
			NamaCabang:            helper.ParseString(getRow(row, 9)),
			KodeCabang:            helper.ParseString(getRow(row, 10)),
			Mufnet:                helper.ParseString(getRow(row, 11)),
			KodeMufnet:            helper.ParseString(getRow(row, 12)),
			FinancialType:         helper.ParseString(getRow(row, 13)),
			ContractNo:            helper.ParseString(getRow(row, 14)),
			CustName:              helper.ParseString(getRow(row, 15)),
			Pekerjaan:             helper.ParseString(getRow(row, 16)),
			Pasangan:              helper.ParseString(getRow(row, 17)),
			Dealer:                helper.ParseString(getRow(row, 18)),
			Merk:                  helper.ParseString(getRow(row, 19)),
			Jenis:                 helper.ParseString(getRow(row, 20)),
			Model:                 helper.ParseString(getRow(row, 21)),
			TahunKendaraan:        helper.ParseString(getRow(row, 22)),
			Nopol:                 helper.ParseString(getRow(row, 23)),
			Rangka:                helper.ParseString(getRow(row, 24)),
			Mesin:                 helper.ParseString(getRow(row, 25)),
			StatusJf:              helper.ParseString(getRow(row, 26)),
			Object:                helper.ParseString(getRow(row, 27)),
			Portfolio:             helper.ParseString(getRow(row, 28)),
			PenggolonganProduct:   helper.ParseString(getRow(row, 29)),
			SalesThrough:          helper.ParseString(getRow(row, 30)),
			MarketingInternal:     helper.ParseString(getRow(row, 31)),
			MarketingHead:         helper.ParseString(getRow(row, 32)),
			IntName:               helper.ParseString(getRow(row, 33)),
			CfoName:               helper.ParseString(getRow(row, 34)),
			NameCa:                helper.ParseString(getRow(row, 35)),

			// numeric safety layer
			Otr:                   helper.ParseFloat(cleanNumber(getRow(row, 36))),
			EffRate:               helper.ParseFloat(cleanNumber(getRow(row, 37))),
			NettDp:                helper.ParseFloat(cleanNumber(getRow(row, 38))),
			Ke:                    helper.ParseInt(cleanNumber(getRow(row, 39))),
			Tenor:                 helper.ParseInt(cleanNumber(getRow(row, 40))),
			JatuhTempo:            helper.ParseTime(getRow(row, 41)),

			StatusKontrak:        helper.ParseString(getRow(row, 42)),
			Hari:                 helper.ParseInt(cleanNumber(getRow(row, 43))),
			BucketAwal:            helper.ParseString(getRow(row, 44)),
			BucketAkhir:           helper.ParseString(getRow(row, 45)),
			PembayaranM1:          helper.ParseString(getRow(row, 46)),
			PembayaranM2:          helper.ParseString(getRow(row, 47)),
			PembayaranM3:          helper.ParseString(getRow(row, 48)),
			PembayaranTerakhir:    helper.ParseTime(getRow(row, 49)),
			PaymentType:           helper.ParseString(getRow(row, 50)),

			ApplPrincipalAmt:      helper.ParseFloat(cleanNumber(getRow(row, 51))),
			Angsuran:              helper.ParseFloat(cleanNumber(getRow(row, 52))),
			OstDenda:              helper.ParseFloat(cleanNumber(getRow(row, 53))),
			TotalTitipan:          helper.ParseFloat(cleanNumber(getRow(row, 54))),
			BalIntr:               helper.ParseFloat(cleanNumber(getRow(row, 55))),
			BalPrin:               helper.ParseFloat(cleanNumber(getRow(row, 56))),

			NoOid:                 helper.ParseString(getRow(row, 57)),
			OdOid:                 helper.ParseInt(cleanNumber(getRow(row, 58))),
			TipeRestructure:       helper.ParseString(getRow(row, 59)),
			NoContractOld:         helper.ParseString(getRow(row, 60)),
			BankPendanaan:         helper.ParseString(getRow(row, 61)),
			FlagAging:             helper.ParseInt(cleanNumber(getRow(row, 62))),
			TglProses:             helper.ParseTime(getRow(row, 63)),
			SettlementNo:          helper.ParseString(getRow(row, 64)),
			Channel:               helper.ParseString(getRow(row, 65)),
			LapanganUsaha:         helper.ParseString(getRow(row, 66)),
			NamaProduct:           helper.ParseString(getRow(row, 67)),
			NamaRefMitra:          helper.ParseString(getRow(row, 68)),
			DueInstall:            helper.ParseInt(cleanNumber(getRow(row, 69))),
			InternalNpk:           helper.ParseString(getRow(row, 70)),
			Posisi:                helper.ParseInt(cleanNumber(getRow(row, 71))),
			TambahanTenor:         helper.ParseInt(cleanNumber(getRow(row, 72))),
			TanggalTarik:          helper.ParseTime(getRow(row, 73)),
			NamaPicTarik:          helper.ParseString(getRow(row, 74)),
		}

		// 🔥 VALIDATION LAYER (WAJIB)
		if d.ContractNo == nil || *d.ContractNo == "" {
			skippedRow++
			continue
		}

		if d.Tenor == nil || *d.Tenor < 0 || *d.Tenor > 99999 {
			skippedRow++
			continue
		}

		data = append(data, d)
	}

	log.Printf("PROCESS DONE | TOTAL=%d | VALID=%d | SKIPPED=%d",
		totalRow, len(data), skippedRow)

	return s.amrepo.BatchInsertRobo(ctx, data)
}

// func (s *HandlingSettingExcelService) ProcessExcel(ctx context.Context, rows [][]string) error {

// 	var data []model.CmsRptAgingMis

// 	for i, row := range rows {
// 		if i == 0 {
// 			continue // skip header
// 		}

// 		// optional: skip row kosong
// 		if len(row) == 0 {
// 			continue
// 		}

// 		d := model.CmsRptAgingMis{
// 			AgingDate:              helper.ParseTime(getRow(row, 0)),
// 			TglProsesAging:        helper.ParseTime(getRow(row, 1)),
// 			TglPpd:                helper.ParseTime(getRow(row, 2)),
// 			MonthPpdDate:          helper.ParseInt(getRow(row, 3)),
// 			YearPpdDate:           helper.ParseString(getRow(row, 4)),
// 			TglPengajuanReschedule:helper.ParseTime(getRow(row, 5)),
// 			HariReschedule:        helper.ParseString(getRow(row, 6)),
// 			NamaArea:              helper.ParseString(getRow(row, 7)),
// 			KodeArea:              helper.ParseString(getRow(row, 8)),
// 			NamaCabang:            helper.ParseString(getRow(row, 9)),
// 			KodeCabang:            helper.ParseString(getRow(row, 10)),
// 			Mufnet:                helper.ParseString(getRow(row, 11)),
// 			KodeMufnet:            helper.ParseString(getRow(row, 12)),
// 			FinancialType:         helper.ParseString(getRow(row, 13)),
// 			ContractNo:            helper.ParseString(getRow(row, 14)),
// 			CustName:              helper.ParseString(getRow(row, 15)),
// 			Pekerjaan:             helper.ParseString(getRow(row, 16)),
// 			Pasangan:              helper.ParseString(getRow(row, 17)),
// 			Dealer:                helper.ParseString(getRow(row, 18)),
// 			Merk:                  helper.ParseString(getRow(row, 19)),
// 			Jenis:                 helper.ParseString(getRow(row, 20)),
// 			Model:                 helper.ParseString(getRow(row, 21)),
// 			TahunKendaraan:        helper.ParseString(getRow(row, 22)),
// 			Nopol:                 helper.ParseString(getRow(row, 23)),
// 			Rangka:                helper.ParseString(getRow(row, 24)),
// 			Mesin:                 helper.ParseString(getRow(row, 25)),
// 			StatusJf:              helper.ParseString(getRow(row, 26)),
// 			Object:                helper.ParseString(getRow(row, 27)),
// 			Portfolio:             helper.ParseString(getRow(row, 28)),
// 			PenggolonganProduct:   helper.ParseString(getRow(row, 29)),
// 			SalesThrough:          helper.ParseString(getRow(row, 30)),
// 			MarketingInternal:     helper.ParseString(getRow(row, 31)),
// 			MarketingHead:         helper.ParseString(getRow(row, 32)),
// 			IntName:               helper.ParseString(getRow(row, 33)),
// 			CfoName:               helper.ParseString(getRow(row, 34)),
// 			NameCa:                helper.ParseString(getRow(row, 35)),
// 			Otr:                   helper.ParseFloat(getRow(row, 36)),
// 			EffRate:               helper.ParseFloat(getRow(row, 37)),
// 			NettDp:                helper.ParseFloat(getRow(row, 38)),
// 			Ke:                    helper.ParseInt(getRow(row, 39)),
// 			Tenor:                 helper.ParseInt(getRow(row, 40)),
// 			JatuhTempo:            helper.ParseTime(getRow(row, 41)),
// 			StatusKontrak:        helper.ParseString(getRow(row, 42)),
// 			Hari:                 helper.ParseInt(getRow(row, 43)),
// 			BucketAwal:            helper.ParseString(getRow(row, 44)),
// 			BucketAkhir:           helper.ParseString(getRow(row, 45)),
// 			PembayaranM1:          helper.ParseString(getRow(row, 46)),
// 			PembayaranM2:          helper.ParseString(getRow(row, 47)),
// 			PembayaranM3:          helper.ParseString(getRow(row, 48)),
// 			PembayaranTerakhir:    helper.ParseTime(getRow(row, 49)),
// 			PaymentType:           helper.ParseString(getRow(row, 50)),
// 			ApplPrincipalAmt:      helper.ParseFloat(getRow(row, 51)),
// 			Angsuran:              helper.ParseFloat(getRow(row, 52)),
// 			OstDenda:              helper.ParseFloat(getRow(row, 53)),
// 			TotalTitipan:          helper.ParseFloat(getRow(row, 54)),
// 			BalIntr:               helper.ParseFloat(getRow(row, 55)),
// 			BalPrin:               helper.ParseFloat(getRow(row, 56)),
// 			NoOid:                 helper.ParseString(getRow(row, 57)),
// 			OdOid:                 helper.ParseInt(getRow(row, 58)),
// 			TipeRestructure:       helper.ParseString(getRow(row, 59)),
// 			NoContractOld:         helper.ParseString(getRow(row, 60)),
// 			BankPendanaan:         helper.ParseString(getRow(row, 61)),
// 			FlagAging:             helper.ParseInt(getRow(row, 62)),
// 			TglProses:             helper.ParseTime(getRow(row, 63)),
// 			SettlementNo:          helper.ParseString(getRow(row, 64)),
// 			Channel:               helper.ParseString(getRow(row, 65)),
// 			LapanganUsaha:         helper.ParseString(getRow(row, 66)),
// 			NamaProduct:           helper.ParseString(getRow(row, 67)),
// 			NamaRefMitra:          helper.ParseString(getRow(row, 68)),
// 			DueInstall:            helper.ParseInt(getRow(row, 69)),
// 			InternalNpk:           helper.ParseString(getRow(row, 70)),
// 			Posisi:                helper.ParseInt(getRow(row, 71)),
// 			TambahanTenor:         helper.ParseInt(getRow(row, 72)),
// 			TanggalTarik:          helper.ParseTime(getRow(row, 73)),
// 			NamaPicTarik:          helper.ParseString(getRow(row, 74)), // ✅ aman sekarang
// 		}

// 		data = append(data, d)
// 	}

// 	return s.amrepo.BatchInsertRobo(ctx, data)
// }

func (s *HandlingSettingExcelService) ExportExcludeContract(ctx context.Context) ([]byte, error) {

	data, err := s.amrepo.GetExcludeContractProd(ctx)
	fmt.Println("DATA LENGTH:", len(data))
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheet := "DataProd"
	f.NewSheet(sheet)
	f.DeleteSheet("Sheet1")

	for i,h := range mapper.HeaderExcludeContract() {
		cell, _ := excelize.CoordinatesToCellName(i+1,1)
		f.SetCellValue(sheet, cell, h)
	}

	mapper.DataExcludeContractMapper(f, sheet, data)

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}