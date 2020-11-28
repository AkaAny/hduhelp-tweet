package cas

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"regexp"
)

func Login(userName string, password string) (*gorequest.SuperAgent, error) {
	var request = gorequest.New()
	request.DoNotClearSuperAgent = true
	request.Header.Add("User-Agent", USER_AGENT)
	var url = "http://cas.hdu.edu.cn/cas/login?service=https%3A%2F%2Fi.hdu.edu.cn%2Ftp_up%2F"

	resp, body, errs := request.Get(url).End()
	if errs != nil {
		return nil, errs[0]
	}
	err := ioutil.WriteFile("cas.html", []byte(body), 0644)
	if err != nil {
		panic(err)
	}
	ltValue, err := getLTValue(body)
	if err != nil {
		panic(err)
	}
	execValue, err := getExecutionValue(body)
	if err != nil {
		panic(err)
	}
	var rsaValue = GetRSAValue(userName, password, ltValue)
	request.Post(url).Type("form")
	resp, body, errs = request.Send(map[string]interface{}{
		"rsa":       rsaValue,
		"ul":        len(userName),
		"pl":        len(password),
		"lt":        ltValue,
		"execution": execValue,
		"_eventId":  "submit",
	}).End()
	if errs != nil {
		return nil, errs[0]
	}
	fmt.Println(resp.StatusCode)
	err = ioutil.WriteFile("resp.html", []byte(body), 0644)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func getLTValue(body string) (string, error) {
	var lineExpr = regexp.MustCompile("<input type=\"hidden\" id=\"lt\" name=\"lt\" value=\".+\" />")
	var ltLine = lineExpr.FindString(body)
	if ltLine == "" {
		return "", errors.New("fail to find lt line")
	}
	fmt.Printf("lt line:%s\n", ltLine)
	return getValue(ltLine)
}

func getExecutionValue(body string) (string, error) {
	var lineExpr = regexp.MustCompile("<input type=\"hidden\" name=\"execution\" value=\".+\" />")
	var execLine = lineExpr.FindString(body)
	if execLine == "" {
		return "", errors.New("fail to find lt line")
	}
	fmt.Printf("execution line:%s\n", execLine)
	return getValue(execLine)
}
