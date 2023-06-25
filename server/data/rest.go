package data

type RestRequest struct {
	Url         string
	Headers     map[string]interface{}
	UrlParams   map[string]interface{}
	QueryParams map[string]interface{}
	Metadata    map[string]interface{}
	Body        []byte
}

type RestResponse struct {
	Data   interface{}
	Status int64
}
