package resource

type AuthResource interface {
	Login()
	Refresh()
	Revoke()
}

type CategoryResource interface {
	GetCategory()
	GetCategories()
}

type GoalResource interface {
	GetGoal()
	GetGoals()
	UpdateGoal()
	CreateGoal()
	DeleteGoal()
}

type Account interface {
	GetAccount()
	UpdateAccount()
	CreateAccount()
	DeleteAccount()
}
