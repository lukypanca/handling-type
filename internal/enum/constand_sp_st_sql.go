package enum

const InsertHandlingSpSpt = `
		INSERT INTO MUFAM.CMS_AR_H_SP_SPT_ (
			DESC_HANDLING, 
			TIPE_HANDLING,
			START_OD,
			END_OD,
			FLAG_ROD,
			STATUS,
			IS_ACTIVE,
			INSERT_DATE,
			INSERT_BY,
			UPDATE_DATE,
			UPDATE_BY
		) VALUES (
			:1, :2, :3, :4, :5, :6, :7, trunc(:8), :9, trunc(:10), :11
		)
		RETURNING HANDLING_SETTING_ID INTO :11
	`

const InsertSpSptBranchQuery = `
		INSERT INTO MUFAM.CMS_AR_H_SP_SPT_BRANCH (
    		HANDLING_SETTING_ID,
    		KODE_CABANG,
    		NAMA_CABANG,
    		KODE_AREA,
    		NAMA_AREA,
    		INSERT_DATE,
    		INSERT_BY,
    		UPDATE_DATE,
    		UPDATE_BY
		)
		VALUES (
    		:1,
    		:2,
    		:3,
    		:4,
    		:5,
    		trunc(:6),
    		:7,
    		trunc(:8),
    		:9
		)
`
const InsertSpSptObjectQuery = `
		INSERT INTO MUFAM.CMS_AR_H_SP_SPT_OBJ_GROUP (
    		HANDLING_SETTING_ID,
    		OBJECT_CODE,
    		OBJECT_GROUP,
    		INSERT_DATE,
    		INSERT_BY,
    		UPDATE_DATE,
    		UPDATE_BY
		)
		VALUES (
    		:1,
    		:2,
    		:3,
    		trunc(:4),
    		:5,
    		trunc(:6),
    		:7
		)
`
const InsertSpSptTipeNasabahQuery = `
INSERT INTO MUFAM.CMS_AR_H_SP_SPT_TIPE_NASABAH (
    		HANDLING_SETTING_ID,
    		TIPE_NASABAH_CODE,
    		TIPE_NASABAH_DESC,
    		INSERT_DATE,
    		INSERT_BY,
    		UPDATE_DATE,
    		UPDATE_BY
		)
		VALUES (
    		:1,
    		:2,
    		:3,
    		trunc(:4),
    		:5,
    		trunc(:6),
    		:7
		)
`

const InsertSpSptCollScoringQuery = `
INSERT INTO MUFAM.CMS_AR_H_SP_SPT_COLL_SCORING (
    		HANDLING_SETTING_ID,
    		COLL_SCORING_CODE,
    		COLL_SCORING_DESC,
    		INSERT_DATE,
    		INSERT_BY,
    		UPDATE_DATE,
    		UPDATE_BY
		)
		VALUES (
    		:1,
    		:2,
    		:3,
    		trunc(:4),
    		:5,
    		trunc(:6),
    		:7
		)
`
const InsertSpSptPaymentTypeQuery = `
INSERT INTO MUFAM.CMS_AR_H_SP_SPT_PAYMENT_TYPE (
    		HANDLING_SETTING_ID,
    		PAYMENT_TYPE_CODE,
    		PAYMENT_TYPE_DESC,
    		INSERT_DATE,
    		INSERT_BY,
    		UPDATE_DATE,
    		UPDATE_BY
		)
		VALUES (
    		:1,
    		:2,
    		:3,
    		trunc(:4),
    		:5,
    		trunc(:6),
    		:7
		)
`

const InsertSpSptTipePembiayaanQuery = `
INSERT INTO MUFAM.CMS_AR_H_SP_SPT_T_PEMBIAYAAN (
    		HANDLING_SETTING_ID,
    		TIPE_PEMBIAYAAN_CODE,
    		TIPE_PEMBIAYAAN_DESC,
    		INSERT_DATE,
    		INSERT_BY,
    		UPDATE_DATE,
    		UPDATE_BY
		)
		VALUES (
    		:1,
    		:2,
    		:3,
    		trunc(:4),
    		:5,
    		trunc(:6),
    		:7
		)
`

const InsertSpSptSkemaPembiayaanQuery = `
INSERT INTO MUFAM.CMS_AR_H_SP_SPT_S_PEMBIAYAAN (
    		HANDLING_SETTING_ID,
    		SKEMA_PEMBIAYAAN_CODE,
    		SKEMA_PEMBIAYAAN_DESC,
    		INSERT_DATE,
    		INSERT_BY,
    		UPDATE_DATE,
    		UPDATE_BY
		)
		VALUES (
    		:1,
    		:2,
    		:3,
    		trunc(:4),
    		:5,
    		trunc(:6),
    		:7
		)
`

const InsertSpSptGolonganProductQuery = `
INSERT INTO MUFAM.CMS_AR_H_SP_SPT_GOL_PRODUK (
    		HANDLING_SETTING_ID,
    		PENGGOLONGAN_PRODUCT_CODE,
    		PENGGOLONGAN_PRODUCT_DESC,
    		INSERT_DATE,
    		INSERT_BY,
    		UPDATE_DATE,
    		UPDATE_BY
		)
		VALUES (
    		:1,
    		:2,
    		:3,
    		trunc(:4),
    		:5,
    		trunc(:6),
    		:7
		)
`

const InsertSpSptBankPendanaanQuery = `
INSERT INTO MUFAM.CMS_AR_H_SP_SPT_BANK_PENDANAAN (
    		HANDLING_SETTING_ID,
    		BANK_CODE,
    		BANK_DESC,
    		INSERT_DATE,
    		INSERT_BY,
    		UPDATE_DATE,
    		UPDATE_BY
		)
		VALUES (
    		:1,
    		:2,
    		:3,
    		trunc(:4),
    		:5,
    		trunc(:6),
    		:7
		)
`
