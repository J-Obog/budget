package rest

type RouteHandler func(*Request) *Response

type Request struct {
	Url    string
	Query  map[string][]string
	Params map[string]string
	Body   []byte
}
