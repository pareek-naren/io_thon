package api

import "io_thon/ciam_thon/httprutils"

func (lr IO) PostForgotPassword(body interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPostReq("/v1/password/forgot", body)
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
