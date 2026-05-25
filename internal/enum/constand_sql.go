package enum

const HandlingCreated = "HANDLING_CREATED"
const ExcludeContractCreated = "EXCLUDE_CONTRACT_CREATED"

const InsertHandlingSetting = `
		INSERT INTO MUFAM.CMS_AR_HANDLING_SETTING (
			HANDLING_SETTING_ID,
			DESC_HANDLING, 
			TIPE_HANDLING,
			TIPE_HANDLING_ID,
			START_OD,
			END_OD,
			STATUS,
			INSERT_DATE,
			INSERT_BY,
			UPDATE_DATE,
			UPDATE_BY
		) VALUES (
		 	MUFAM.CMS_AR_HANDLING_SEQ.NEXTVAL,
			:1, :2, :3, :4, :5, :6, trunc(:7), :8, trunc(:9), :10
		)
		RETURNING HANDLING_SETTING_ID INTO :11
	`

const InsertBranchQuery = `
		INSERT INTO MUFAM.CMS_AR_HANDLING_BRANCH (
    		HANDLING_SETTING_ID,
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
const InsertObjectQuery = `
		INSERT INTO MUFAM.CMS_AR_HANDLING_OBJECT_GROUP (
    		HANDLING_SETTING_ID,
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
const InsertTipeNasabahQuery = `
INSERT INTO MUFAM.CMS_AR_HANDLING_TIPE_NASABAH (
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

const InsertCollScoringQuery = `
INSERT INTO MUFAM.CMS_AR_HANDLING_COLL_SCORING (
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
const InsertPaymentTypeQuery = `
INSERT INTO MUFAM.CMS_AR_HANDLING_PAYMENT_TYPE (
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

const InsertTipePembiayaanQuery = `
INSERT INTO MUFAM.CMS_AR_HANDLING_T_PEMBIAYAAN (
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

const InsertSkemaPembiayaanQuery = `
INSERT INTO MUFAM.CMS_AR_HANDLING_S_PEMBIAYAAN (
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

const InsertGolonganProductQuery = `
INSERT INTO MUFAM.CMS_AR_HANDLING_GOL_PRODUK (
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

const InsertBankPendanaanQuery = `
INSERT INTO MUFAM.CMS_AR_HANDLING_BANK_PENDANAAN (
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
const InsertCmsArExcludeContract = `
	INSERT INTO MUFAM.CMS_AR_EXCLUDE_CONTRACT (
		CONTRACT_NO,
		NAMA_NASABAH,
		OBJECT_GROUP,
		TIPE_NASABAH,
		COLL_SCORING,
		PAYMENT_TYPE,
		TIPE_PEMBIAYAAN,
		SKEMA_PEMBIAYAAN,
		PENGGOLONGAN_PRODUCT,
		BANK_PENDANAAN,
		MARKETING_PROGRAM,
		TIPE_HANDLING,
		STATUS_EXCLUDE,
		IS_ACTIVE,
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
		:6,
		:7,
		:8,
		:9,
		:10,
		:11,
		:12,
		:13,
		:14,
		TRUNC(:15),
		:16,
		TRUNC(:17),
		:18
	)
`