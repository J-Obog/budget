package rest

type Request struct {
	Url string
	//Query  map[string][]string
	//Params map[string]string
	Body map[string]any
}

type Response struct {
	Data   any
	Status int
}
