package api

type APIConfig struct {
	AccountAPI     *AccountAPI
	BudgetAPI      *BudgetAPI
	TransactionAPI *TransactionAPI
	CategoryAPI    *CategoryAPI
}

func CreateConfig() *APIConfig {

}
