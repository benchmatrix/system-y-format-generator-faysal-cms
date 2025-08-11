package CMS

type TotalAssetsDetails struct {
	CurrentAssets              []TotalAssetsChild       `json:"currentassets"`
	FixedAssets                []TotalAssetsChild       `json:"fixedassets"`
	Mortgage                   []TotalAssetsChild       `json:"mortgage"`
	StocksHeldUnderPledge      []TotalAssetsChild       `json:"stocks"`
	CashCollateral             []TotalAssetsChild       `json:"cashcollateral"`
	Gop                        []TotalAssetsChild       `json:"gop"`
	PledgeOfShares             []TotalAssetsChild       `json:"pledge"`
	CollateralSummary          []CollateralSummary      `json:"collateralsummary"`
	TotalCurrentAssets         []TotalAssetsChildTotal  `json:"totalcurrentassets"`
	TotalFixedAssets           []TotalAssetsChildTotal  `json:"totalfixedassets"`
	TotalMortgage              []TotalAssetsChildTotal  `json:"totalmortgage"`
	TotalStocksHeldUnderPledge []TotalAssetsChildTotal  `json:"totalstocks"`
	TotalCashCollateral        []TotalAssetsChildTotal  `json:"totalcashcollateral"`
	TotalGop                   []TotalAssetsChildTotal  `json:"totalgop"`
	TotalPledgeOfShares        []TotalAssetsChildTotal  `json:"totalpledge"`
	TotalCollateralSummary     []TotalCollateralSummary `json:"totalcollateralsummary"`
}
