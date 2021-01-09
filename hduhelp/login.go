package hduhelp

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
	"hduhelp-tweet/cas"
	"io/ioutil"
	url2 "net/url"
	"strings"
)

func Login(userName string, password string) (*gorequest.SuperAgent, error) {
	request, err := cas.Login(userName, password)
	if err != nil {
		panic(err)
	}
	request.Header.Set("User-Agent", MICRO_MESSAGE_USER_AGENT)
	request.RedirectPolicy(func(req gorequest.Request, via []gorequest.Request) error {
		var urlStr = req.URL.String()
		fmt.Println("redirect:", urlStr)
		if strings.Contains(urlStr, "auth") {
			var actualQuery = req.URL.RawFragment
			actualQuery = strings.Split(actualQuery, "?")[1]
			fmt.Println(actualQuery)
			queryMap, err := url2.ParseQuery(actualQuery)
			if err != nil {
				return err //取消继续跳转
			}
			var token = queryMap.Get("auth")
			fmt.Println("token:", token)
			request.Header.Add("Authorization", fmt.Sprintf("token %s", token))
			request.Header.Add("X-Hduhelp-Vfy", "")
			//恢复默认重定向策略
			//request.RedirectPolicy(func(req gorequest.Request, via []gorequest.Request) error {
			//	return nil
			//})
		}
		return nil
	})
	var url = "https://api.hduhelp.com/login/direct/cas?clientID=app&school=hdu&redirect=https%3A%2F%2Fapp.hduhelp.com%2F%23%2Flogin%3FredirectTo%3D%252F"
	resp, body, errs := request.Get(url).End()
	if errs != nil {
		return nil, errs[0]
	}
	fmt.Println(resp.StatusCode)
	err = ioutil.WriteFile("resp.html", []byte(body), 0644)
	if err != nil {
		return nil, err
	}

	validateResponse, err := Validate(request)
	if err != nil {
		return nil, err
	}
	if !validateResponse.Valid() {
		return nil, errors.Errorf("token is invalid (expire:%v)", validateResponse.GetExpireTime())
	}
	return request, nil
}

func Validate(request *gorequest.SuperAgent) (*ValidateResponse, error) {
	//这个validate验证响应有两个data，内容相同，大小写不同
	resp, body, errs := request.Get("https://api.hduhelp.com/token/validate").End()
	if errs != nil {
		return nil, errs[0]
	}
	fmt.Println(resp.StatusCode, body)
	baseResponse, err := FromBody(body)
	if err != nil {
		return nil, err
	}
	err = baseResponse.AsError()
	if err != nil {
		return nil, err
	}
	var validateResponse ValidateResponse
	err = baseResponse.UnmarshallData(&validateResponse)
	if err != nil {
		return nil, err
	}
	return &validateResponse, nil
}
