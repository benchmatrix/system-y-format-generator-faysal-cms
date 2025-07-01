package AuthModels

type CustomError struct {
	ErrorCode   int    `json:"ErrorCode"`
	Status      string `json:"Status"`
	Description string `json:"Description"`
}
