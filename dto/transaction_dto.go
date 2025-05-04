package dto

type TransactionQueryParam struct {
	SourceID   uint   `json:"source_id" form:"source_id"`
	CategoryID uint   `json:"category_id" form:"category_id"`
	Purpose    string `json:"purpose" form:"purpose"`
	Remark     string `json:"remark" form:"remark"`
	StartDate  string `json:"start_date" form:"start_date"`
	EndDate    string `json:"end_date" form:"end_date"`
	Type       string `json:"type" form:"type"`
}

type TransactionSummaryQueryParam struct {
	Period string `json:"period" form:"period"`
}

type TransactionSummaryData struct {
	Period        string  `json:"period"`
	IncomeAmount  float64 `json:"income_amount"`
	ExpenseAmount float64 `json:"expense_amount"`
}
