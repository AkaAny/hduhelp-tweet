package main

import (
	"hduhelp-tweet/cas"
	"hduhelp-tweet/zf"
)

func main() {
	request, err := cas.Login(CAS_USER_NAME, CAS_PASSWORD)
	if err != nil {
		panic(err)
	}
	var zfObj = zf.Create(ZF_XH, ZF_XM)
	err = zfObj.Init(request)
	if err != nil {
		panic(err)
	}
	err = zfObj.GetClassTable()
	if err != nil {
		panic(err)
	}
}
