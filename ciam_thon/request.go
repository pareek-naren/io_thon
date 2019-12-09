package IO

import "io_thon/ciam_thon/httprutils"

// NewPostReq takes a uri, body, and optional queries to construct a POST request for a IO POST API endpoint
func (lr IO) NewPostReq(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {
	encodedBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := &httprutils.Request{
		Method: httprutils.Post,
		URL:    lr.Domain + path,
		Headers: map[string]string{
			"content-Type": "application/json",
		},
		QueryParams: map[string]string{
			"apiKey": lr.Context.ApiKey,
		},
		Body: encodedBody,
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}
	return request, nil
}
