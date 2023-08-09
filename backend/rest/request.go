package rest

type RouteHandler func(*Request) *Response

type Request struct {
	Url    string
	Query  map[string][]string
	Params PathParams
	Body   []byte
}
