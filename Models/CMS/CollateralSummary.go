package CMS

type CollateralSummary struct {
	CurrentAssetAmount   interface{} `json:"currentassetamount"`
	CurrentAssetStatus   interface{} `json:"currentassetstatus"`
	FixedAssetAmount     interface{} `json:"fixedassetamount"`
	FixedAssetStatus     interface{} `json:"fixedassetstatus"`
	MortgageAmount       interface{} `json:"mortgageamount"`
	MortgageStatus       interface{} `json:"mortgagestatus"`
	StockAmount          interface{} `json:"stockamount"`
	StockStatus          interface{} `json:"stockstatus"`
	CashCollateralAmount interface{} `json:"cashcollateralamount"`
	CashCollateralStatus interface{} `json:"cashcollateralstatus"`
	GopAmount            interface{} `json:"gopamount"`
	GopStatus            interface{} `json:"gopstatus"`
	PledgeAmount         interface{} `json:"pledgeamount"`
	PledgeStatus         interface{} `json:"pledgestatus"`
}
