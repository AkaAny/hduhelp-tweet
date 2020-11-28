package zf

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
)

type Zf struct {
	request *gorequest.SuperAgent
	XH      string
	XM      string
}

func (zf Zf) GetRequest() *gorequest.SuperAgent {
	return zf.request
}

func Create(xh string, xm string) *Zf {
	return &Zf{
		XH: xh,
		XM: xm,
	}
}

func (zf *Zf) Init(request *gorequest.SuperAgent) error {
	var jwUrl = fmt.Sprintf("http://jxgl.hdu.edu.cn/default.aspx")
	resp, body, errs := request.Get(jwUrl).End()
	if errs != nil {
		return errs[0]
	}
	fmt.Println(resp.StatusCode)
	err := ioutil.WriteFile("resp_ticket.html", []byte(body), 0644)
	if err != nil {
		panic(err)
	}

	var mainPageUrl = fmt.Sprintf("http://jxgl.hdu.edu.cn/xs_main.aspx?xh=%s", zf.XH)
	resp, body, errs = request.Get(mainPageUrl).End()
	if errs != nil {
		return errs[0]
	}
	fmt.Println(resp.StatusCode)
	err = ioutil.WriteFile("main_page.html", []byte(body), 0644)
	if err != nil {
		panic(err)
	}
	//正方教务系统的坑，必须加Referer头：https://blog.csdn.net/qiangrenpu8881/article/details/82260774
	request.Header.Add("Referer", mainPageUrl)

	zf.request = request
	return nil
}

func (zf Zf) GetClassTable() error {
	xmEncoded, err := toGBKUrlEncoded(zf.XM)
	if err != nil {
		panic(err)
	}
	var tableUrl = fmt.Sprintf("http://jxgl.hdu.edu.cn/xskbcx.aspx?xh=%s&xm=%s&gnmkdm=N121603",
		zf.XH, xmEncoded)
	resp, body, errs := zf.GetRequest().Get(tableUrl).End()
	if errs != nil {
		return errs[0]
	}
	fmt.Println(resp.StatusCode)
	err = ioutil.WriteFile("table.html", []byte(body), 0644)
	if err != nil {
		panic(err)
	}
	return nil
}
