package CMS

type FacilityDetails struct {
	Line               interface{} `json:"line"`
	FacilityDetails    interface{} `json:"facilitydetails"`
	ControlUnit        interface{} `json:"controlunit"`
	Request            interface{} `json:"request"`
	Secured            interface{} `json:"secured"`
	MaxTenor           interface{} `json:"maxtenor"`
	AvailabilityPeriod interface{} `json:"availabilityperiod"`
	ProposedLimit      interface{} `json:"proposedlimit"`
	ExistingLimit      interface{} `json:"existinglimit"`
	ProposedFRR        interface{} `json:"proposedfrr"`
	ExistingFRR        interface{} `json:"existingfrr"`
	RepaymentTerms     interface{} `json:"repaymentTerms"`
	Pricing            interface{} `json:"pricing"`
	Purpose            interface{} `json:"purpose"`
	ProposedSecurities interface{} `json:"proposedsecurities"`
	ExistingSecurities interface{} `json:"existingsecurities"`
	Condition          interface{} `json:"condition"`
	RepaymentHistory   interface{} `json:"repaymenthistory"`
	ParentId           interface{} `json:"parentid"`
	ExistingParentId   interface{} `json:"exctparentid"`
	FacilityExpiry     interface{} `json:"expirydate"`
}
