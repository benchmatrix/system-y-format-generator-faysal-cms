package CMS

type DeferalWaiversOpenExceptionDetails struct {
	Ref            interface{} `json:"ref"`
	Description    interface{} `json:"description"`
	ExistingExpiry interface{} `json:"existingexpiry"`
	ProposedExpiry interface{} `json:"proposedexpiry"`
	Status         interface{} `json:"status"`
	Comments       interface{} `json:"comments"`
	ExpiryBool     interface{} `json:"expirybool"`

	TITLE                   interface{} `json:"title"`
	EXISTING_DEFFERAL_DAYS  interface{} `json:"existing_deferal_days"`
	EXISTING_DEFERRALCOUNT  interface{} `json:"existing_deferal_count"`
	EXISTING_DEFERRALEXPIRY interface{} `json:"existing_deferal_expiry"`
	PROPOSED_DEFFERAL_DAYS  interface{} `json:"proposed_deferal_days"`
	PROPOSED_DEFERRALCOUNT  interface{} `json:"proposed_deferal_count"`
	PROPOSED_DEFERRALEXPIRY interface{} `json:"proposed_deferal_expiry"`
	PROPOSED_STATUS         interface{} `json:"proposed_status"`
	TOTAL_DAYS              interface{} `json:"total_days"`
}
