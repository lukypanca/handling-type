package enum

const InsertHandlingSpSptCms = `
		INSERT INTO MUFCMS.AR_H_SP (
			ID,
			DESC_HANDLING, 
			TIPE_HANDLING,
			TIPE_HANDLING_ID,
			START_OD,
			END_OD,
			FLAG_ROD,
			STATUS,
			INSERT_DATE,
			INSERT_BY,
			UPDATE_DATE,
			UPDATE_BY
		) VALUES (
		 	MUFCMS.AR_HANDLING_SEQ.NEXTVAL,
			:1, :2, :3, :4, :5, :6, :7, trunc(:8), :9, trunc(:10), :11
		)
		RETURNING ID INTO :12
	`

const InsertSpSptBranchQueryCms = `
		INSERT INTO MUFCMS.AR_H_SP_BRANCH  (
    		HANDLING_ID,
    		KODE_CABANG,
    		NAMA_CABANG,
    		KODE_AREA,
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
    		trunc(:5),
    		:6,
    		trunc(:7),
    		:8
		)
`
const InsertSpSptObjectQueryCms = `
		INSERT INTO MUFCMS.AR_H_SP_OBJ_GROUP (
    		HANDLING_ID,
    		OBJECT_GROUP,
    		INSERT_DATE,
    		INSERT_BY,
    		UPDATE_DATE,
    		UPDATE_BY
		)
		VALUES (
    		:1,
    		:2,
    		trunc(:3),
    		:4,
    		trunc(:5),
    		:6
		)
`
const InsertSpSptTipeNasabahQueryCms = `
INSERT INTO MUFCMS.AR_H_SP_TIPE_NASABAH (
    		HANDLING_ID,
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
INSERT INTO MUFCMS.AR_H_SP_COLL_SCORING (
    		HANDLING_ID,
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
INSERT INTO MUFCMS.AR_H_SP_PAYMENT_TYPE (
    		HANDLING_ID,
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
INSERT INTO MUFCMS.AR_H_SP_T_PEMBIAYAAN (
    		HANDLING_ID,
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
INSERT INTO MUFCMS.AR_H_SP_S_PEMBIAYAAN (
    		HANDLING_ID,
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
INSERT INTO MUFCMS.AR_H_SP_GOL_PRODUK (
    		HANDLING_ID,
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
INSERT INTO MUFCMS.AR_H_SP_BANK_PENDANAAN (
    		HANDLING_ID,
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
