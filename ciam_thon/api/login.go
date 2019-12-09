package api

import "io_thon/ciam_thon/httprutils"

func (lr IO) PostLogin(body interface{}, verificationurl string, emailtemplate string) (*httprutils.Response, error) {
	allowedQueries := map[string]string{
		"verificationurl": verificationurl, "emailtemplate": emailtemplate,
	}
	request, err := lr.Client.NewPostReq("/v1/email/check", body, allowedQueries)
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
