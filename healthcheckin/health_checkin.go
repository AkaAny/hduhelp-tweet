package healthcheckin

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"hduhelp-tweet/hduhelp"
	"net/http"
)

type HealthCheckin struct {
	request *gorequest.SuperAgent
}

func (checkin HealthCheckin) GetRequest() *gorequest.SuperAgent {
	return checkin.request
}

func (checkin *HealthCheckin) Init(request *gorequest.SuperAgent) {
	request.Header.Set("Referer", "https://healthcheckin.hduhelp.com/")
	//不加Origin头会403
	request.Header.Set("Origin", "https://healthcheckin.hduhelp.com")
	checkin.request = request
}

func (checkin HealthCheckin) PostCheckin(timeStamp int64) error {
	personInfo, err := checkin.GetPersonInfo()
	if err != nil {
		return err
	}
	fmt.Println(personInfo)
	var model = PostModel{
		Name:          personInfo.StaffName,
		TimeStamp:     timeStamp, //time.Now().Unix(),
		Province:      "330000",
		City:          "330100",
		Country:       "330104",
		AnswerJsonStr: ANSWER_JSON_STR,
	}
	sign, err := getSign(personInfo.StaffID, model)
	if err != nil {
		return err
	}
	var url = fmt.Sprintf("https://api.hduhelp.com/base/healthcheckin?sign=%s", sign)
	fmt.Println(url)
	var request = checkin.GetRequest()
	resp, _, errs := request.Options(url).End()
	fmt.Println(resp.StatusCode)
	if errs != nil && resp.StatusCode != http.StatusNoContent {
		return errs[0]
	}

	request.Post(url).Type("json")
	resp, body, errs := request.Send(model).End()
	fmt.Println(resp.StatusCode)
	if errs != nil {
		return errs[0]
	}
	fmt.Println(body)
	return nil
}

func (checkin HealthCheckin) GetPersonInfo() (*PersonInfo, error) {
	var url = "https://api.hduhelp.com/base/person/info"
	resp, body, errs := checkin.GetRequest().Get(url).End()
	if errs != nil {
		return nil, errs[0]
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
	var result PersonInfo
	err = baseResponse.UnmarshallData(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (checkin HealthCheckin) GetDailyStatus() (*DailyStatus, error) {
	var url = "https://api.hduhelp.com/base/healthcheckin/info/daily"
	resp, body, errs := checkin.GetRequest().Get(url).End()
	if errs != nil {
		return nil, errs[0]
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
	var result DailyStatus
	err = baseResponse.UnmarshallData(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
