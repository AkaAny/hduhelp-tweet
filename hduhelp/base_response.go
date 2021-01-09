package hduhelp

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type BaseResponse struct {
	Error   int64                  `json:"error"`
	Message string                 `json:"msg"`
	Data    map[string]interface{} `json:"data"` //这里有泛型就好了，可以对应不同API的data
	Cache   bool                   `json:"cache"`
}

func (br BaseResponse) UnmarshallData(obj interface{}) error {
	rawData, err := json.Marshal(br.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(rawData, obj)
}

func (br BaseResponse) AsError() error {
	if br.Error == 0 {
		return nil
	}
	return errors.Errorf("error:%d msg:%s", br.Error, br.Message)
}

func FromBody(body string) (*BaseResponse, error) {
	var br BaseResponse
	var err = json.Unmarshal([]byte(body), &br)
	if err != nil {
		return nil, err
	}
	return &br, err
}
