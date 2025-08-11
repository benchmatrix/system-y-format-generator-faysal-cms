package CMS

type MonitoringSummary struct {
	Takaful       interface{} `json:"takaful"`
	Stocks        interface{} `json:"stocks"`
	FixedAssets   interface{} `json:"fixedassets"`
	CurrentAssets interface{} `json:"currentassets"`
	Pledge        interface{} `json:"pledge"`
}
