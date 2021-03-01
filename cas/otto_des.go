package cas

import (
	"errors"
	"fmt"
	"github.com/dop251/goja"
	"github.com/parnurzeal/gorequest"
	"hduhelp-tweet/wrapper"
	"io/ioutil"
)

func GetRSAValue(request *gorequest.SuperAgent, userName string, password string, ltValue string) (string, error) {
	rawJs, err := getDESJs(request)
	if err != nil {
		return "", err
	}
	vm := goja.New()
	_, err = vm.RunString(string(rawJs))
	if err != nil {
		panic(err)
	}
	var inputStr = userName + password + ltValue
	strEnc, valid := goja.AssertFunction(vm.Get("strEnc"))
	if !valid {
		return "", errors.New("invalid js")
	}
	value, err := strEnc(nil, vm.ToValue(inputStr), vm.ToValue("1"), vm.ToValue("2"), vm.ToValue("3"))
	if err != nil {
		panic(err)
	}
	var result = value.String()
	fmt.Println(result)
	return result, nil
}

func getDESJs(request *gorequest.SuperAgent) (string, error) {
	rawJs, err := ioutil.ReadFile(wrapper.GetPath("assets/des.js"))
	if err == nil {
		return string(rawJs), nil
	}
	var url = "https://cas.hdu.edu.cn/cas/comm/js/des.js"
	resp, body, errs := request.Get(url).End()
	fmt.Println(resp.StatusCode)
	if errs != nil {
		return "", errs[0]
	}
	return body, nil
}
