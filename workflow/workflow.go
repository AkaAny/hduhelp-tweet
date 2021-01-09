package workflow

import "github.com/parnurzeal/gorequest"

type WorkFlow struct {
	request *gorequest.SuperAgent
}

func (workflow WorkFlow) GetRequest() *gorequest.SuperAgent {
	return workflow.request
}

func (workflow *WorkFlow) Init(request *gorequest.SuperAgent) {
	request.Header.Set("User-Agent", USER_AGENT)
	request.Header.Set("Referer", "https://salmon.hduhelp.com/idCode/")
	request.Header.Set("Origin", "https://salmon.hduhelp.com")
	request.Header.Set("X-Requested-With", "com.alibaba.android.rimet")

	workflow.request = request
}
