package resource

import "github.com/J-Obog/paidoff/db"

type Response struct {
	Body    interface{}
	Success bool
}

type Request struct {
	Url         string
	UrlParams   map[string]interface{}
	QueryParams map[string]interface{}
	Metadata    map[string]interface{}
	Body        []byte
}

type AuthLoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRefreshResponse struct {
	AccessToken string `json:"accessToken"`
}

type AuthRevokeRequest struct {
	Token string `json:"token"`
}

type AccountResponse struct {
	Id                   string `json:"id"`
	Email                string `json:"email"`
	NotificationsEnabled bool   `json:"notificationsEnabled"`
	CreatedAt            int64  `json:"createdAt"`
	UpdatedAt            int64  `json:"updatedAt"`
}

type AccountUpdateRequest struct {
	Email                string `json:"email"`
	NotificationsEnabled bool   `json:"notificationsEnabled"`
}

type AccountCreateRequest struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CategoryResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ImageUrl  string `json:""`
	CreatedAt int64
	UpdatedAt int64
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

type GoalUpdateRequest struct {
	CategoryId    *string  `json:"categoryId"`
	Month         *int64   `json:"month"`
	Year          *int64   `json:"year"`
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
