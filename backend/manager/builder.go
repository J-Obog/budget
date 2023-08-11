package manager

type ManagerService struct {
	AccountManager     *AccountManager
	CategoryManager    *CategoryManager
	TransactionManager *TransactionManager
	BudgetManager      *BudgetManager
}

/*
func NewManagerService(cfg *config.AppConfig) *ManagerService {
	storeSvc := store.NewStoreService(cfg)
	clock := clock.NewClock(cfg)
	uuidProvider := uuid.NewUuidProvider(cfg)

	return &ManagerService{
		AccountManager:     NewAccountManager(storeSvc.AccountStore, clock),
		BudgetManager:      NewBudgetManager(storeSvc.BudgetStore, uuidProvider, clock),
		TransactionManager: NewTransactionManager(storeSvc.TransactionStore, uuidProvider, clock),
		CategoryManager:    NewCategoryManager(storeSvc.CategoryStore, uuidProvider, clock),
	}
}*/
