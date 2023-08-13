package rest

type RouteHandler func(*Request) *Response

type Request struct {
	Url    string
	Query  Query
	Params PathParams
	Body   *JSONBody
}
