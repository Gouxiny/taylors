package param

type AllListParam struct {
	Page             int
	PageSize         int
	Name             string
	Code             string
	MarketCapitalMax int64
	MarketCapitalMin int64
	PercentMax       float64
	PercentMin       float64
	VolumeRatioMax   float64
	VolumeRatioMin   float64
	CurrentMax       float64
	CurrentMin       float64
	CreateTime       int64
}

type AnalysisListParam struct {
	Page             int
	PageSize         int
	MarketCapitalMax int64
	MarketCapitalMin int64
	PercentMax       float64
	PercentMin       float64
	VolumeRatioMax   float64
	VolumeRatioMin   float64
	CurrentMax       float64
	CurrentMin       float64
	StartTime        int64
	EndTime          int64
	DayMin           int
}
