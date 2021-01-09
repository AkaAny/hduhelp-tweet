package workflow

import "github.com/parnurzeal/gorequest"

func PutHeader(request *gorequest.SuperAgent, key string, value string) {
	_, ok := request.Header[key]
	if ok {
		request.Header.Set(key, value)
		return
	}
	request.Header.Add(key, value)
}
