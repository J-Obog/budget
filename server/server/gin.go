package server

type GinServer struct {
	BaseServer
}

func (gin *GinServer) Start(port int) error {
	return nil
}

func (gin *GinServer) Stop() error {
	return nil
}
