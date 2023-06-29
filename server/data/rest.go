package data

type RestRequest struct {
	Url         string
	Headers     map[string]interface{}
	Meta        map[string]interface{}
	UrlParams   map[string]interface{}
	QueryParams map[string]interface{}
	Metadata    map[string]interface{}
	Body        []byte
}

type RestResponse struct {
	Data   interface{}
	Status int
}

type AccountUpdateRequest struct {
	Name string `json:"name"`
}

type CategoryCreateRequest struct {
	Name  string `json:"name"`
	Color int    `json:"color"`
}

type CategoryUpdateRequest struct {
	CategoryCreateRequest
}

type TransactionCreateRequest struct {
	CategoryId  *string `json:"categoryId"`
	Description *string `json:"description"`
	Amount      float64 `json:"amount"`
	Month       int     `json:"month"`
	Day         int     `json:"day"`
	Year        int     `json:"year"`
}

type TransactionUpdateRequest struct {
	TransactionCreateRequest
}
