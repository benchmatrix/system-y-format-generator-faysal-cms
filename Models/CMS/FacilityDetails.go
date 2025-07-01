package CMS

type FacilityDetails struct {
	Line                    interface{} `json:"line"`
	FacilityDetails         interface{} `json:"facilitydetails"`
	ControlUnit             interface{} `json:"controlunit"`
	Request                 interface{} `json:"request"`
	Classification          interface{} `json:"regulatoryclassification"`
	DateofClassification    interface{} `json:"dateofdefaultclassification"`
	Secured                 interface{} `json:"secured"`
	MaxTenor                interface{} `json:"maxtenor"`
	Maturity                interface{} `json:"maturity"`
	AvailabilityPeriod      interface{} `json:"availabilityperiod"`
	CommFees                interface{} `json:"commfees"`
	Grace                   interface{} `json:"grace"`
	ProposedLimit           interface{} `json:"proposedlimit"`
	ExistingLimit           interface{} `json:"existinglimit"`
	ProposedFRR             interface{} `json:"proposedfrr"`
	ExistingFRR             interface{} `json:"existingfrr"`
	FacilityChange          interface{} `json:"changefacility"`
	ReasonforChangeFacility interface{} `json:"reasonforchangefacility"`
	RepaymentTerms          interface{} `json:"repaymentTerms"`
	Pricing                 interface{} `json:"pricing"`
	BenchMarkRate           interface{} `json:"benchmarkrate"`
	Purpose                 interface{} `json:"purpose"`
	ProposedSecurities      interface{} `json:"proposedsecurities"`
	ExistingSecurities      interface{} `json:"existingsecurities"`
	Condition               interface{} `json:"condition"`
	RepaymentHistory        interface{} `json:"repaymenthistory"`
	CollateralChange        interface{} `json:"collateralchange"`
	ParentId                interface{} `json:"parentid"`
	Commission              interface{} `json:"commission"`
	FacilityType            interface{} `json:"facilitytype"`
	ExistingParentId        interface{} `json:"exctparentid"`
	FacilityExpiry          interface{} `json:"expirydate"`
	//SecurityDetails         []SecurityDetails `json:"securitydetails"`
}
