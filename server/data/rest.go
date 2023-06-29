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
