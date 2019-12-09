package IO

import (
	"errors"
	"io_thon/ciam_thon/lrerror"
)

const domain = "https:///api.loginradius.io"

type IO struct {
	Context *Context
	Domain  string
}

type Config struct {
	ApiKey    string
	ApiSecret string
}

type Context struct {
	ApiKey    string
	ApiSecret string
	Token     string
}

func NewIO(cfg *Config, optionalArgs ...map[string]string) (*IO, error) {

	if cfg.ApiKey == "" || cfg.ApiSecret == "" {
		errMsg := "Must initialize IO client with ApiKey and ApiSecret"
		err := lrerror.New("IntializationError", errMsg, errors.New(errMsg))
		return nil, err
	}

	ctx := Context{
		ApiKey:    cfg.ApiKey,
		ApiSecret: cfg.ApiSecret,
	}

	// If an access token is passed on initiation, set it in Context
	for _, arg := range optionalArgs {
		if arg["token"] != "" {
			ctx.Token = arg["token"]
		} else {
			ctx.Token = ""
		}
	}

	return &IO{
		Context: &ctx,
		Domain:  domain,
	}, nil
}
