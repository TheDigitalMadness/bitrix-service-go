package bitrixModels

type StageID string

const (
	REGISTRATION StageID = "C2:NEW"
	PROVIDER     StageID = "C2:EXECUTING"
	FIRST_BUY    StageID = "C2:PREPARATION"
	REPEAT_BUY   StageID = "C2:PREPAYMENT_INVOICE"
)

type AddDealFields struct {
	CategoryID int
	StageID    StageID
	Title      string
}

type UpdateDealFields struct {
	CategoryID int
	StageID    StageID
}
