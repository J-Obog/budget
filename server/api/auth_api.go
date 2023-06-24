package api

import "github.com/J-Obog/paidoff/data"

type AuthAPI struct {
}

func (api *AuthAPI) Login(req *data.RestRequest, res *data.RestResponse) error {
	return nil
}

func (api *AuthAPI) Register(req *data.RestRequest, res *data.RestResponse) error {
	return nil
}

func (api *AuthAPI) Logout(req *data.RestRequest, res *data.RestResponse) error {
	return nil
}

func (api *AuthAPI) Refresh(req *data.RestRequest, res *data.RestResponse) error {
	return nil
}

func (api *AuthAPI) ConfirmEmail(req *data.RestRequest, res *data.RestResponse) error {
	return nil
}
