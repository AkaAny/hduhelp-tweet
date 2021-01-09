package workflow

import (
	"fmt"
	"hduhelp-tweet/hduhelp"
	"net/http"
)

func (workflow WorkFlow) GetIDCode() (*IDCode, error) {
	var request = workflow.GetRequest()
	var url = "https://api.hduhelp.com/workflow/id/code"
	resp, _, errs := request.Options(url).End()
	fmt.Println(resp.StatusCode)
	if errs != nil && resp.StatusCode != http.StatusNoContent {
		panic(errs[0])
	}
	resp, body, errs := request.Get(url).End()
	if errs != nil {
		panic(errs[0])
	}
	fmt.Println(resp.StatusCode, body)
	baseResponse, err := hduhelp.FromBody(body)
	if err != nil {
		return nil, err
	}
	err = baseResponse.AsError()
	if err != nil {
		return nil, err
	}
	var result IDCode
	err = baseResponse.UnmarshallData(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
