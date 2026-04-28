package enum

const InsertHandlingSpSptCms = `
		INSERT INTO MUFCMS.AR_HANDLING_SETTING_ (
			DESC_HANDLING, 
			TIPE_HANDLING,
			START_OD,
			END_OD,
			STATUS,
			IS_ACTIVE,
			INSERT_DATE,
			INSERT_BY,
			UPDATE_DATE,
			UPDATE_BY
		) VALUES (
			:1, :2, :3, :4, :5, :6, trunc(:7), :8, trunc(:9), :10
		)
		RETURNING HANDLING_SETTING_ID INTO :11
	`

const InsertSpSptBranchQueryCms = `
		INSERT INTO MUFCMS.AR_HANDLING_BRANCH (
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
const InsertSpSptObjectQueryCms = `
		INSERT INTO MUFCMS.AR_HANDLING_OBJECT_GROUP (
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
const InsertSpSptTipeNasabahQueryCms = `
INSERT INTO MUFCMS.AR_HANDLING_TIPE_NASABAH (
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

const InsertSpSptCollScoringQueryCms = `
INSERT INTO MUFCMS.AR_HANDLING_COLL_SCORING (
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
const InsertSpSptPaymentTypeQueryCms = `
INSERT INTO MUFCMS.AR_HANDLING_PAYMENT_TYPE (
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

const InsertSpSptTipePembiayaanQueryCms = `
INSERT INTO MUFCMS.AR_HANDLING_T_PEMBIAYAAN (
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

const InsertSpSptSkemaPembiayaanQueryCms = `
INSERT INTO MUFCMS.AR_HANDLING_S_PEMBIAYAAN (
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

const InsertSpSptGolonganProductQueryCms = `
INSERT INTO MUFCMS.AR_HANDLING_GOL_PRODUK (
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

const InsertSpSptBankPendanaanQueryCms = `
INSERT INTO MUFCMS.AR_HANDLING_BANK_PENDANAAN (
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
