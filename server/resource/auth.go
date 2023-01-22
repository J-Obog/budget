package resource

type AuthResource struct {
}

func (this *AuthResource) Login(req Request) *Response {
	return nil
}

func (this *AuthResource) Refresh(req Request) *Response {
	return nil
}

func (this *AuthResource) Revoke(req Request) *Response {
	return nil
}
