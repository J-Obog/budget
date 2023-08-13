package store

import (
	"github.com/J-Obog/paidoff/config"
)

const (
	storeImpl = "postgres"
)

type StoreService struct {
	AccountStore     AccountStore
	BudgetStore      BudgetStore
	CategoryStore    CategoryStore
	TransactionStore TransactionStore
}

func NewStoreService(app *config.AppConfig) *StoreService {
	/*switch storeImpl {

	case "postgres":
		pgUrl := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			app.PostgresUser,
			app.PostgresPassword,
			app.PostgresHost,
			app.PostgresPort,
			app.PostgresDb)

		pgDb, err := gorm.Open(postgres.Open(pgUrl), &gorm.Config{
			AllowGlobalUpdate: true,
			Logger:            logger.Default.LogMode(logger.Silent),
		})

		if err != nil {
			log.Fatal(err)
		}
		return &StoreService{
			AccountStore:     &PostgresAccountStore{pgDb},
			CategoryStore:    &PostgresCategoryStore{pgDb},
			TransactionStore: &PostgresTransactionStore{pgDb},
			BudgetStore:      &PostgresBudgetStore{pgDb},
		}
	default:
		log.Fatal("Not a supported impl for store")
	}*/

	return nil
}
