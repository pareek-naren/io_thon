package api

import (
	"io_thon/ciam_thon/httprutils"
)

func (lr IO) RegisterUser(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPostReq("/v1/register", body)
	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"verificationUrl": true, "emailTemplate": true,
		}
		validatedQueries, err := Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			request.QueryParams[k] = v
		}
	}
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
