package workflow

import (
	"encoding/json"
	"fmt"
	"hduhelp-tweet/hduhelp"
)

func initWorkFlow(userName string, password string) (*WorkFlow, error) {
	request, err := hduhelp.Login(userName, password)
	if err != nil {
		return nil, err
	}
	var workflow WorkFlow
	workflow.Init(request)
	return &workflow, nil
}

func APIGetIDCode(userName string, password string) string {
	wf, err := initWorkFlow(userName, password)
	if err != nil {
		fmt.Println(err)
	}
	idCode, err := wf.GetIDCode()
	if err != nil {
		fmt.Println(err)
	}
	rawData, err := json.Marshal(idCode)
	if err != nil {
		fmt.Println(err)
	}
	return string(rawData)
}
