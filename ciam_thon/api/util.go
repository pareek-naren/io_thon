package api

import (
	"errors"
	"fmt"
	"io_thon/ciam_thon/lrerror"
)

func Validate(allowed map[string]bool, params interface{}) (map[string]string, error) {
	asserted, ok := params.(map[string]string)

	if !ok {
		errMsg := fmt.Sprintf("Error validating params: %+v:", params)
		err := lrerror.New("ValidationError", "Error validating params - params type error", errors.New(errMsg))
		return nil, err
	}

	for k, _ := range asserted {
		if !allowed[k] {
			err := lrerror.New("ValidationError", "Error validating params - invalid params submitted, please double check", errors.New("Error validating params"))
			return nil, err
		}
	}
	return asserted, nil
}
