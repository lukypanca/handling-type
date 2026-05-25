package am

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"tipe-handling/internal/helper"
	"tipe-handling/internal/model"
)

type HandlingSettingExcelRepository struct {
	db *sql.DB
}

func NewHandlingSettingExcelRepository(db *sql.DB) *HandlingSettingExcelRepository {
	return &HandlingSettingExcelRepository{db: db}
}

func (r *HandlingSettingExcelRepository) GetRoboAiData(ctx context.Context, branches []string) ([]model.CmsRptAgingMis, error) {

	inClause, args := helper.BuildInClause(branches)
	query := fmt.Sprintf(`
			SELECT 
				AG.AGING_DATE,
				AG.TGL_PROSES_AGING,
				AG.TGL_PPD,
				AG.MONTH_PPD_DATE,
				AG.YEAR_PPD_DATE,
				AG.TGL_PENGAJUAN_RESCHEDULE,
				AG.HARI_RESCHEDULE,
				AG.NAMA_AREA,
				AG.KODE_AREA,
				AG.NAMA_CABANG,
				AG.KODE_CABANG,
				AG.MUFNET,
				AG.KODE_MUFNET,
				AG.FINANCIAL_TYPE,
				AG.CONTRACTNO,
				AG.CUSTNAME,
				AG.PEKERJAAN,
				AG.PASANGAN,
				AG.DEALER,
				AG.MERK,
				AG.JENIS,
				AG.MODEL,
				AG.TAHUN_KENDARAAN,
				AG.NOPOL,
				AG.RANGKA,
				AG.MESIN,
				AG.STATUS_JF,
				AG.OBJECT,
				AG.PORTFOLIO,
				AG.PENGGOLONGAN_PRODUCT,
				AG.SALES_THROUGH,
				AG.MARKETING_INTERNAL,
				AG.MARKETING_HEAD,
				AG.INT_NAME,
				AG.CFO_NAME,
				AG.NAME_CA,
				AG.OTR,
				AG.EFF_RATE,
				AG.NETT_DP,
				AG.KE,
				AG.TENOR,
				AG.JATUH_TEMPO,
				AG.STATUS_KONTRAK,
				AG.HARI,
				AG.BUCKET_AWAL,
				AG.BUCKET_AKHIR,
				AG.PEMBAYARAN_M1,
				AG.PEMBAYARAN_M2,
				AG.PEMBAYARAN_M3,
				AG.PEMBAYARAN_TERAKHIR,
				AG.PAYMENT_TYPE,
				AG.APPL_PRINCIPAL_AMT,
				AG.ANGSURAN,
				AG.OST_DENDA,
				AG.TOTAL_TITIPAN,
				AG.BAL_INTR,
				AG.BAL_PRIN,
				AG.NO_OID,
				AG.OD_OID,
				AG.TIPE_RESTRUCTURE,
				AG.NO_CONTRACT_OLD,
				AG.BANK_PENDANAAN,
				AG.FLAG_AGING,
				AG.TGL_PROSES,
				AG.SETTLEMENT_NO,
				AG.CHANNEL,
				AG.NAMA_PRODUCT,
				AG.NAMA_REF_MITRA,
				AG.LAPANGAN_USAHA,
				AG.DUEINSTALL,
				AG.INTERNAL_NPK,
				AG.POSISI,
				AG.TAMBAHAN_TENOR,
				AG.TANGGAL_TARIK,
				AG.NAMA_PIC_TARIK
			FROM MUFAM.CMS_RPT_AGING_MIS AG
			INNER JOIN MUFAM.CONTRACT_MASTER CM
				ON AG.KODE_CABANG = CM.BRANCH_CODE
			AND AG.CONTRACTNO = CM.CONTRACT_NO
			INNER JOIN MUFAM.CONTRACT_LOAN_STRUCTURE CLS
				ON CM.CONTRACT_ID = CLS.CONTRACT_ID
			LEFT JOIN MUFAM.STG_PINALTY_AMOUNT_CMS SPA
				ON SPA.BR_ID = CM.BRANCH_CODE
			AND SPA.CONT_NO = CM.CONTRACT_NO
			INNER JOIN MUFAM.APPLICATIONS APP
				ON APP.APPLICATION_ID = CM.APPLICATION_ID
			INNER JOIN MUFPARAM.MST_MKT_PROGRAM MMP
				ON CLS.PROD_PROG_CODE = MMP.MKT_PROGRAM_CODE
			INNER JOIN MUFAM.CONTRACT_OBJECT O
				ON CM.CONTRACT_ID = O.CONTRACT_ID
			WHERE AG.AGING_DATE = TRUNC(SYSDATE - 2)
			AND AG.HARI BETWEEN 0 AND 60
			AND CM.BRANCH_CODE IN (%s)
			AND NOT EXISTS (SELECT 1
					FROM MUFAM.TEMP_DATA_JF DJ
					WHERE DJ.CONTRACT_NO = CM.CONTRACT_NO
					AND DJ.APPLICATION_TYPE IS NOT NULL
					AND AG.HARI <= 7)
			AND ROWNUM <= 100	
	`, inClause)

	log.Println("query : ", query)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("DB ERROR: %v", err)
		log.Printf("QUERY: %s", query)
		return nil, err
	}
	defer rows.Close()

	var result []model.CmsRptAgingMis

	for rows.Next() {
		var d model.CmsRptAgingMis

		err := rows.Scan(
			&d.AgingDate,
			&d.TglProsesAging,
			&d.TglPpd,
			&d.MonthPpdDate,
			&d.YearPpdDate,
			&d.TglPengajuanReschedule,
			&d.HariReschedule,
			&d.NamaArea,
			&d.KodeArea,
			&d.NamaCabang,
			&d.KodeCabang,
			&d.Mufnet,
			&d.KodeMufnet,
			&d.FinancialType,
			&d.ContractNo,
			&d.CustName,
			&d.Pekerjaan,
			&d.Pasangan,
			&d.Dealer,
			&d.Merk,
			&d.Jenis,
			&d.Model,
			&d.TahunKendaraan,
			&d.Nopol,
			&d.Rangka,
			&d.Mesin,
			&d.StatusJf,
			&d.Object,
			&d.Portfolio,
			&d.PenggolonganProduct,
			&d.SalesThrough,
			&d.MarketingInternal,
			&d.MarketingHead,
			&d.IntName,
			&d.CfoName,
			&d.NameCa,
			&d.Otr,
			&d.EffRate,
			&d.NettDp,
			&d.Ke,
			&d.Tenor,
			&d.JatuhTempo,
			&d.StatusKontrak,
			&d.Hari,
			&d.BucketAwal,
			&d.BucketAkhir,
			&d.PembayaranM1,
			&d.PembayaranM2,
			&d.PembayaranM3,
			&d.PembayaranTerakhir,
			&d.PaymentType,
			&d.ApplPrincipalAmt,
			&d.Angsuran,
			&d.OstDenda,
			&d.TotalTitipan,
			&d.BalIntr,
			&d.BalPrin,
			&d.NoOid,
			&d.OdOid,
			&d.TipeRestructure,
			&d.NoContractOld,
			&d.BankPendanaan,
			&d.FlagAging,
			&d.TglProses,
			&d.SettlementNo,
			&d.Channel,
			&d.NamaProduct,
			&d.NamaRefMitra,
			&d.LapanganUsaha,
			&d.DueInstall,
			&d.InternalNpk,
			&d.Posisi,
			&d.TambahanTenor,
			&d.TanggalTarik,
			&d.NamaPicTarik,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, d)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// func (r *HandlingSettingExcelRepository) BatchInsertRobo(ctx context.Context, data []model.CmsRptAgingMis) error {

// 	tx, err := r.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()

// 	query := `INSERT INTO MUFAM.CMS_RPT_AGING_MIS_ (
// 		AGING_DATE, TGL_PROSES_AGING, TGL_PPD, MONTH_PPD_DATE, YEAR_PPD_DATE,
// 		TGL_PENGAJUAN_RESCHEDULE, HARI_RESCHEDULE, NAMA_AREA, KODE_AREA, NAMA_CABANG,
// 		KODE_CABANG, MUFNET, KODE_MUFNET, FINANCIAL_TYPE, CONTRACTNO,
// 		CUSTNAME, PEKERJAAN, PASANGAN, DEALER, MERK,
// 		JENIS, MODEL, TAHUN_KENDARAAN, NOPOL, RANGKA,
// 		MESIN, STATUS_JF, OBJECT, PORTFOLIO, PENGGOLONGAN_PRODUCT,
// 		SALES_THROUGH, MARKETING_INTERNAL, MARKETING_HEAD, INT_NAME, CFO_NAME,
// 		NAME_CA, OTR, EFF_RATE, NETT_DP, KE,
// 		TENOR, JATUH_TEMPO, STATUS_KONTRAK, HARI, BUCKET_AWAL,
// 		BUCKET_AKHIR, PEMBAYARAN_M1, PEMBAYARAN_M2, PEMBAYARAN_M3, PEMBAYARAN_TERAKHIR,
// 		PAYMENT_TYPE, APPL_PRINCIPAL_AMT, ANGSURAN, OST_DENDA, TOTAL_TITIPAN,
// 		BAL_INTR, BAL_PRIN, NO_OID, OD_OID, TIPE_RESTRUCTURE,
// 		NO_CONTRACT_OLD, BANK_PENDANAAN, FLAG_AGING, TGL_PROSES, SETTLEMENT_NO,
// 		CHANNEL, NAMA_PRODUCT, NAMA_REF_MITRA, LAPANGAN_USAHA, DUEINSTALL,
// 		INTERNAL_NPK, POSISI, TAMBAHAN_TENOR, TANGGAL_TARIK, NAMA_PIC_TARIK
// 	) VALUES (
// 		:1,:2,:3,:4,:5,:6,:7,:8,:9,:10,
// 		:11,:12,:13,:14,:15,:16,:17,:18,:19,:20,
// 		:21,:22,:23,:24,:25,:26,:27,:28,:29,:30,
// 		:31,:32,:33,:34,:35,:36,:37,:38,:39,:40,
// 		:41,:42,:43,:44,:45,:46,:47,:48,:49,:50,
// 		:51,:52,:53,:54,:55,:56,:57,:58,:59,:60,
// 		:61,:62,:63,:64,:65,:66,:67,:68,:69,:70,
// 		:71,:72,:73,:74,:75
// 	)`

// 	for _, d := range data {
// 		_, err := tx.ExecContext(ctx, query,
// 			d.AgingDate,
// 			d.TglProsesAging,
// 			d.TglPpd,
// 			d.MonthPpdDate,
// 			d.YearPpdDate,
// 			d.TglPengajuanReschedule,
// 			d.HariReschedule,
// 			d.NamaArea,
// 			d.KodeArea,
// 			d.NamaCabang,
// 			d.KodeCabang,
// 			d.Mufnet,
// 			d.KodeMufnet,
// 			d.FinancialType,
// 			d.ContractNo,
// 			d.CustName,
// 			d.Pekerjaan,
// 			d.Pasangan,
// 			d.Dealer,
// 			d.Merk,
// 			d.Jenis,
// 			d.Model,
// 			d.TahunKendaraan,
// 			d.Nopol,
// 			d.Rangka,
// 			d.Mesin,
// 			d.StatusJf,
// 			d.Object,
// 			d.Portfolio,
// 			d.PenggolonganProduct,
// 			d.SalesThrough,
// 			d.MarketingInternal,
// 			d.MarketingHead,
// 			d.IntName,
// 			d.CfoName,
// 			d.NameCa,
// 			d.Otr,
// 			d.EffRate,
// 			d.NettDp,
// 			d.Ke,
// 			d.Tenor,
// 			d.JatuhTempo,
// 			d.StatusKontrak,
// 			d.Hari,
// 			d.BucketAwal,
// 			d.BucketAkhir,
// 			d.PembayaranM1,
// 			d.PembayaranM2,
// 			d.PembayaranM3,
// 			d.PembayaranTerakhir,
// 			d.PaymentType,
// 			d.ApplPrincipalAmt,
// 			d.Angsuran,
// 			d.OstDenda,
// 			d.TotalTitipan,
// 			d.BalIntr,
// 			d.BalPrin,
// 			d.NoOid,
// 			d.OdOid,
// 			d.TipeRestructure,
// 			d.NoContractOld,
// 			d.BankPendanaan,
// 			d.FlagAging,
// 			d.TglProses,
// 			d.SettlementNo,
// 			d.Channel,
// 			d.NamaProduct,
// 			d.NamaRefMitra,
// 			d.LapanganUsaha,
// 			d.DueInstall,
// 			d.InternalNpk,
// 			d.Posisi,
// 			d.TambahanTenor,
// 			d.TanggalTarik,
// 			d.NamaPicTarik,
// 		)

// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return tx.Commit()
// }

func (r *HandlingSettingExcelRepository) BatchInsertRobo(ctx context.Context, data []model.CmsRptAgingMis) error {

	query := `INSERT INTO MUFAM.CMS_RPT_AGING_MIS_ (
		AGING_DATE, TGL_PROSES_AGING, TGL_PPD, MONTH_PPD_DATE, YEAR_PPD_DATE,
		TGL_PENGAJUAN_RESCHEDULE, HARI_RESCHEDULE, NAMA_AREA, KODE_AREA, NAMA_CABANG,
		KODE_CABANG, MUFNET, KODE_MUFNET, FINANCIAL_TYPE, CONTRACTNO,
		CUSTNAME, PEKERJAAN, PASANGAN, DEALER, MERK,
		JENIS, MODEL, TAHUN_KENDARAAN, NOPOL, RANGKA,
		MESIN, STATUS_JF, OBJECT, PORTFOLIO, PENGGOLONGAN_PRODUCT,
		SALES_THROUGH, MARKETING_INTERNAL, MARKETING_HEAD, INT_NAME, CFO_NAME,
		NAME_CA, OTR, EFF_RATE, NETT_DP, KE,
		TENOR, JATUH_TEMPO, STATUS_KONTRAK, HARI, BUCKET_AWAL,
		BUCKET_AKHIR, PEMBAYARAN_M1, PEMBAYARAN_M2, PEMBAYARAN_M3, PEMBAYARAN_TERAKHIR,
		PAYMENT_TYPE, APPL_PRINCIPAL_AMT, ANGSURAN, OST_DENDA, TOTAL_TITIPAN,
		BAL_INTR, BAL_PRIN, NO_OID, OD_OID, TIPE_RESTRUCTURE,
		NO_CONTRACT_OLD, BANK_PENDANAAN, FLAG_AGING, TGL_PROSES, SETTLEMENT_NO,
		CHANNEL, NAMA_PRODUCT, NAMA_REF_MITRA, LAPANGAN_USAHA, DUEINSTALL,
		INTERNAL_NPK, POSISI, TAMBAHAN_TENOR, TANGGAL_TARIK, NAMA_PIC_TARIK
	) VALUES (
		:1,:2,:3,:4,:5,:6,:7,:8,:9,:10,
		:11,:12,:13,:14,:15,:16,:17,:18,:19,:20,
		:21,:22,:23,:24,:25,:26,:27,:28,:29,:30,
		:31,:32,:33,:34,:35,:36,:37,:38,:39,:40,
		:41,:42,:43,:44,:45,:46,:47,:48,:49,:50,
		:51,:52,:53,:54,:55,:56,:57,:58,:59,:60,
		:61,:62,:63,:64,:65,:66,:67,:68,:69,:70,
		:71,:72,:73,:74,:75
	)`

	var (
		successCount int
		failedCount  int
	)

	for _, d := range data {

		_, err := r.db.ExecContext(ctx, query,
			d.AgingDate,
			d.TglProsesAging,
			d.TglPpd,
			d.MonthPpdDate,
			d.YearPpdDate,
			d.TglPengajuanReschedule,
			d.HariReschedule,
			d.NamaArea,
			d.KodeArea,
			d.NamaCabang,
			d.KodeCabang,
			d.Mufnet,
			d.KodeMufnet,
			d.FinancialType,
			d.ContractNo,
			d.CustName,
			d.Pekerjaan,
			d.Pasangan,
			d.Dealer,
			d.Merk,
			d.Jenis,
			d.Model,
			d.TahunKendaraan,
			d.Nopol,
			d.Rangka,
			d.Mesin,
			d.StatusJf,
			d.Object,
			d.Portfolio,
			d.PenggolonganProduct,
			d.SalesThrough,
			d.MarketingInternal,
			d.MarketingHead,
			d.IntName,
			d.CfoName,
			d.NameCa,
			d.Otr,
			d.EffRate,
			d.NettDp,
			d.Ke,
			d.Tenor,
			d.JatuhTempo,
			d.StatusKontrak,
			d.Hari,
			d.BucketAwal,
			d.BucketAkhir,
			d.PembayaranM1,
			d.PembayaranM2,
			d.PembayaranM3,
			d.PembayaranTerakhir,
			d.PaymentType,
			d.ApplPrincipalAmt,
			d.Angsuran,
			d.OstDenda,
			d.TotalTitipan,
			d.BalIntr,
			d.BalPrin,
			d.NoOid,
			d.OdOid,
			d.TipeRestructure,
			d.NoContractOld,
			d.BankPendanaan,
			d.FlagAging,
			d.TglProses,
			d.SettlementNo,
			d.Channel,
			d.NamaProduct,
			d.NamaRefMitra,
			d.LapanganUsaha,
			d.DueInstall,
			d.InternalNpk,
			d.Posisi,
			d.TambahanTenor,
			d.TanggalTarik,
			d.NamaPicTarik,
		)

		if err != nil {
			failedCount++

			// IMPORTANT: safe logging (tanpa pointer bug)
			fmt.Printf(
				"[SKIP ROW] CONTRACTNO=%s | ERROR=%v\n",
				d.ContractNo,
				err,
			)

			continue
		}

		successCount++
	}

	fmt.Printf(
		"[BATCH DONE] SUCCESS=%d | FAILED=%d\n",
		successCount,
		failedCount,
	)

	return nil
}

func (r *HandlingSettingExcelRepository) GetExcludeContractProd(ctx context.Context) ([]model.CmsArExcludeContract, error) {

	query := fmt.Sprintf(`
			SELECT CM.CONTRACT_NO CONTRACT_NO,
			       CUM.CUSTOMER_NAME NAMA_NASABAH,
			       (SELECT MOG.OBJ_GROUP_DESC
			          FROM MUFPARAM.MST_OBJ_GROUP MOG
			         WHERE MOG.OBJ_GROUP_CODE = CASE
			                 WHEN CO.OBJECT_CODE IN ('001', '002') THEN
			                  '001'
			                 WHEN CO.OBJECT_CODE IN ('003', '004') THEN
			                  '002'
			               END) OBJECT_GROUP,
			       (SELECT CT.CUST_DESC
			          FROM MUFAM.CMS_CUST_TYPE_OBJ_PRICING COP
			         INNER JOIN MUFAM.MST_CUSTOMER_TYPE CT
			            ON COP.CUST_TYPE_ID = CT.CUST_TYPE_ID
			         WHERE COP.OBJ_CODE = CO.OBJECT_CODE
			           AND CO.OBJECT_PRICE >= COP.START_PRICE
			           AND (COP.END_PRICE IS NULL OR CO.OBJECT_PRICE <= COP.END_PRICE)) TIPE_NASABAH,
			       (SELECT SC.LAST_COLL_PRIORITY
			          FROM MUFCMS.AR_COLL_SCORING_CONTRACT@TO_MUFCMS SC
			         WHERE SC.CONT_NO = CM.CONTRACT_NO
			           AND SC.PROCESS_DATE =
			               (SELECT MAX(S.PROCESS_DATE)
			                  FROM MUFCMS.AR_COLL_SCORING_CONTRACT@TO_MUFCMS S
			                 WHERE S.CONT_NO = CM.CONTRACT_NO)) COLL_SCORING,
			       (SELECT PM.MST_PAY_METODE_DESC
			          FROM MUFPARAM.MST_PAYMENT_METHODE PM
			         WHERE CM.PAYMENT_METHOD = PM.MST_PAY_METHODE_CODE) PAYMENT_TYPE,
			       (SELECT MFT.FIN_TYPE_DESC
			          FROM MUFPARAM.MST_FINANCING_TYPE MFT
			         WHERE CM.FINANCE_TYPE = MFT.FIN_TYPE_CODE) TIPE_PEMBIAYAAN,
			       (SELECT MFP.FIN_PROD_DESC
			          FROM MUFPARAM.MST_FINANCING_PRODUCT MFP
			         WHERE MFP.FIN_PROD_CODE = CM.FIN_SCHEME_CODE) SKEMA_PEMBIAYAAN,
			       CASE
			         WHEN AP.CHANNEL_TYPE_CODE = '08' AND CM.CONTRACT_ID_REF IS NOT NULL THEN
			          NVL((SELECT (CASE MPM2.APPL_PROD_CODE
			                       WHEN '1' THEN
			                        'NON CAPTIVE'
			                       WHEN '2' THEN
			                        'CAPTIVE'
			                       WHEN '3' THEN
			                        'MULTIGUNA NON CAPTIVE'
			                       WHEN '4' THEN
			                        'MULTIGUNA CAPTIVE'
			                       ELSE
			                        'RESTRUCTURE'
			                     END)
			                FROM MUFAM.CONTRACT_OBJECT       COB,
			                     MUFAM.MST_PORTFOLIO_MAPPING MPM2
			               WHERE COB.CONTRACT_ID = CM.CONTRACT_ID_REF
			                 AND COB.APPLICATION_TAG2 = MPM2.PROGRAM_DETAIL_ID
			                 AND COB.OBJECT_CODE = MPM2.OBJECT_CODE
			                 AND ROWNUM = 1),
			              (CASE MPM.APPL_PROD_CODE
			                WHEN '1' THEN
			                 'NON CAPTIVE'
			                WHEN '2' THEN
			                 'CAPTIVE'
			                WHEN '3' THEN
			                 'MULTIGUNA NON CAPTIVE'
			                WHEN '4' THEN
			                 'MULTIGUNA CAPTIVE'
			                ELSE
			                 'RESTRUCTURE'
			              END))
			         ELSE
			          (SELECT CASE MPM.APPL_PROD_CODE
			                    WHEN '1' THEN
			                     'NON CAPTIVE'
			                    WHEN '2' THEN
			                     'CAPTIVE'
			                    WHEN '3' THEN
			                     'MULTIGUNA NON CAPTIVE'
			                    WHEN '4' THEN
			                     'MULTIGUNA CAPTIVE'
			                    WHEN '5' THEN
			                     'RESTRUCTURE'
			                    ELSE
			                     NULL
			                  END
			             FROM MUFAM.MST_PORTFOLIO_MAPPING MPM
			            WHERE MPM.OBJECT_CODE = CO.OBJECT_CODE
			              AND MPM.PORTFOLIO_CODE = CM.PORTFOLIO_CODE
			              AND ROWNUM = 1)
			       END PENGGOLONGAN_PRODUCT,
			       (SELECT MB.BANK_NAME
			          FROM MUFAM.MST_BANK MB
			         WHERE CM.FUND_BANK_ID = MB.BANK_CODE
			           AND CM.BRANCH_CODE = MB.BANK_BR_ID) BANK_PENDANAAN,
			       (SELECT MKT_PROGRAM_DESC
			          FROM MUFPARAM.MST_MKT_PROGRAM MMP
			         WHERE CLS.PROD_PROG_CODE = MMP.MKT_PROGRAM_CODE) MARKETING_PROGRAM,
			       EC.HANDLING_TYPES TIPE_HANDLING,
			       'SUDAH EXCLUDE' STATUS_EXCLUDE
			  FROM MUFAM.CONTRACT_MASTER CM
			  JOIN MUFAM.CUSTOMER_MASTER CUM
			    ON CM.CUSTOMER_ID = CUM.CUSTOMER_ID
			  JOIN MUFAM.CONTRACT_OBJECT CO
			    ON CM.CONTRACT_ID = CO.CONTRACT_ID
			  JOIN MUFAM.APPLICATIONS AP
			    ON CM.APPLICATION_ID = AP.APPLICATION_ID
			  JOIN MUFAM.CONTRACT_LOAN_STRUCTURE CLS
			    ON CM.CONTRACT_ID = CLS.CONTRACT_ID
			  JOIN MUFAM.MST_PORTFOLIO_MAPPING MPM
			    ON CO.APPLICATION_TAG1 = MPM.PROGRAM_ID
			   AND CO.APPLICATION_TAG2 = MPM.PROGRAM_DETAIL_ID
			   AND CO.OBJECT_CODE = MPM.OBJECT_CODE
			  JOIN (SELECT EC.CONTRACT_NO,
			               LISTAGG(CASE
			                         WHEN EC.HANDLING_TYPE = 'DESKCOLL - SIM' THEN
			                          'Deskcoll (SIM)'
			                         WHEN EC.HANDLING_TYPE = 'ROBO' THEN
			                          'Robo AI'
			                         ELSE
			                          EC.HANDLING_TYPE
			                       END,
			                       ', ') WITHIN GROUP(ORDER BY EC.HANDLING_TYPE) AS HANDLING_TYPES
			          FROM (SELECT DISTINCT CONTRACT_NO, HANDLING_TYPE
			                  FROM 	@TO_MUFCMS
			                 WHERE IS_ACTIVE = 1) EC
			         GROUP BY EC.CONTRACT_NO) EC
			    ON CM.CONTRACT_NO = EC.CONTRACT_NO
	`)

	log.Println("query GetExcludeContractProd : ", query)
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("DB ERROR: %v", err)
		log.Printf("QUERY: %s", query)
		return nil, err
	}
	defer rows.Close()

	var result []model.CmsArExcludeContract

	for rows.Next() {
		var d model.CmsArExcludeContract

		err := rows.Scan(
			&d.ContractNo,
			&d.NamaNasabah,
			&d.ObjectGroup,
			&d.TipeNasabah,
			&d.CollScoring,
			&d.PaymentType,
			&d.TipePembiayaan,
			&d.SkemaPembiayaan,
			&d.PenggolonganProduct,
			&d.BankPendanaan,
			&d.MarketingProgram,
			&d.TipeHandling,
			&d.StatusExclude,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, d)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
