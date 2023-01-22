package resource

import "github.com/J-Obog/paidoff/db"

type Response struct {
	Body   []byte
	Status int
}

type Request struct {
	Url         string
	UrlParams   map[string]interface{}
	QueryParams map[string]interface{}
	Meta        map[string]interface{}
	Body        []byte
}

//responses
type AuthLoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthRefreshResponse struct {
	AccessToken string `json:"accessToken"`
}

type AccountResponse struct {
	Id                   string `json:"id"`
	Email                string `json:"email"`
	NotificationsEnabled bool   `json:"notificationsEnabled"`
	CreatedAt            int64  `json:"createdAt"`
	UpdatedAt            int64  `json:"updatedAt"`
}

type CategoryResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ImageUrl  string `json:""`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type CategoryListResponse struct {
	Categories []CategoryResponse `json:"categories"`
}

type GoalResponse struct {
	Id            string      `json:"id"`
	AccountId     string      `json:"accountId"`
	CategoryId    string      `json:"categoryId"`
	Month         int64       `json:"month"`
	Year          int64       `json:"year"`
	Name          string      `json:"name"`
	CurrentAmount float64     `json:"currentAmount"`
	TargetAmount  float64     `json:"targetAmount"`
	GoalType      db.GoalType `json:"goalType"`
	CreatedAt     int64       `json:"createdAt"`
	UpdatedAt     int64       `json:"updatedAt"`
}

type GoalListResponse struct {
	Goals []GoalResponse `json:"goals"`
}

//requests
type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRevokeRequest struct {
	Token string `json:"token"`
}

type AccountUpdateRequest struct {
	Email                *string `json:"email"`
	Password             *string `json:"password"`
	NotificationsEnabled *bool   `json:"notificationsEnabled"`
}

type AccountCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GoalUpdateRequest struct {
	CategoryId    *string  `json:"categoryId"`
	Name          *string  `json:"name"`
	CurrentAmount *float64 `json:"currentAmount"`
	TargetAmount  *float64 `json:"targetAmount"`
}

type GoalCreateRequest struct {
	CategoryId    string      `json:"categoryId"`
	Month         int64       `json:"month"`
	Year          int64       `json:"year"`
	Name          string      `json:"name"`
	CurrentAmount float64     `json:"currentAmount"`
	TargetAmount  float64     `json:"targetAmount"`
	GoalType      db.GoalType `json:"goalType"`
}
