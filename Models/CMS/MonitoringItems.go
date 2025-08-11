package CMS

type MonitoringItems struct {
	Takaful           []MonitoringItemsChild `json:"takaful"`
	Stocks            []MonitoringItemsChild `json:"stocks"`
	FixedAssets       []MonitoringItemsChild `json:"fixedassets"`
	CurrentAssets     []MonitoringItemsChild `json:"currentassets"`
	Pledge            []MonitoringItemsChild `json:"pledge"`
	MonitoringSummary []MonitoringSummary    `json:"monitoringsummary"`
}
