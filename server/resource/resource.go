package resource

type AuthResource interface {
	Login(req Request) Response
	Refresh(req Request) Response
	Revoke(req Request) Response
}

type CategoryResource interface {
	GetCategory(req Request) Response
	GetCategories(req Request) Response
}

type GoalResource interface {
	GetGoal(req Request) Response
	GetGoals(req Request) Response
	UpdateGoal(req Request) Response
	CreateGoal(req Request) Response
	DeleteGoal(req Request) Response
}

type Account interface {
	GetAccount(req Request) Response
	UpdateAccount(req Request) Response
	CreateAccount(req Request) Response
	DeleteAccount(req Request) Response
}
