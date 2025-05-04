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

type CreateTransactionRequest struct {
	TrxDate    string  `json:"trx_date"`
	Type       string  `json:"type"`
	SourceID   uint    `json:"source_id"`
	CategoryID uint    `json:"category_id"`
	Amount     float64 `json:"amount"`
	Purpose    string  `json:"purpose"`
	Remark     string  `json:"remark"`
}

type TransactionCategoryQueryParam struct {
	Type string `json:"type" form:"type"`
}

type TransactionSourceQueryParam struct {
	Type string `json:"type" form:"type"`
}

type TransactionSummaryData struct {
	Period        string  `json:"period"`
	IncomeAmount  float64 `json:"income_amount"`
	ExpenseAmount float64 `json:"expense_amount"`
}
