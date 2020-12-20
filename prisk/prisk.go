package prisk

import (
	"crypto/sha256"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strconv"
	"strings"
	"time"
)

const (
	KEY_MOBILE = "2CA32596474B4077834CCC191D351839" //查询指定地区
	KEY_PC     = "3C502C97ABDA40D0A60FBEE50FAAD1DA" //全部风险地区
)

type H map[string]string

type PRisk struct {
	request *gorequest.SuperAgent
}

func (pr PRisk) GetRequest() *gorequest.SuperAgent {
	return pr.request
}

func CreatePRisk() *PRisk {
	var request = gorequest.New()
	request.DoNotClearSuperAgent = true
	request.Header.Set("User-Agent", USER_AGENT)
	request.Header.Set("Referer", "http://bmfw.www.gov.cn/")
	request.Header.Set("Origin", "http://bmfw.www.gov.cn/")
	return &PRisk{request: request}
}

func getParamSign(key string, timeStamp string, areaCode string) H { //timestamp
	var i = "23y0ufFl5YxIyGrI8hWRUZmKkvtSjLQA"
	var nonce = "123456789abcdefg" //nonce
	var rawSign = timeStamp + i + nonce + timeStamp
	fmt.Println(rawSign)
	var sign = strings.ToUpper(sha256Hex(rawSign))
	fmt.Println(sign)
	var paramMap = H{
		"appId":           "NcApplication",
		"paasHeader":      "zdww",
		"timestampHeader": timeStamp,
		"nonceHeader":     nonce,
		"signatureHeader": sign,
		//确定接口类型，返回指定地区还是全部风险地区
		"key":       key,
		"area_code": areaCode,
	}
	return paramMap
}

func (pr PRisk) get(key string, areaCode string) { //timestamp
	var request = pr.GetRequest()
	var timeStamp = fmt.Sprintf("%d", time.Now().Unix()) //s
	var rawSign = timeStamp + "fTN2pfuisxTavbTuYVSsNJHetwq5bJvCQkjjtiLM2dCratiA" + timeStamp
	fmt.Println(rawSign)
	var sign = strings.ToUpper(sha256Hex(rawSign))
	fmt.Println(sign)
	var headerMap = H{
		"x-wif-nonce":     "QkjjtiLM2dCratiA",
		"x-wif-paasid":    "smt-application",
		"x-wif-signature": sign,
		"x-wif-timestamp": timeStamp,
	}
	putHeaderMap(request, headerMap)

	var paramMap = getParamSign(key, timeStamp, areaCode)
	request.Post("http://103.66.32.242:8005/zwfwMovePortal/interface/interfaceJson").Type("json")
	resp, body, errs := request.Send(paramMap).End()
	fmt.Println(resp.StatusCode)
	if errs != nil {
		fmt.Println(errs)
		return
	}
	fmt.Println(body)
}

func (pr PRisk) GetArea(areaCode string) { //timestamp
	pr.get(KEY_MOBILE, areaCode)
}

func (pr PRisk) GetAll() {
	pr.get(KEY_PC, "110101") //这里的areaCode会被忽略
}

func putHeaderMap(request *gorequest.SuperAgent, headerMap H) {
	for k, v := range headerMap {
		request.Header.Set(k, v)
	}
}

func sha256Hex(str string) string {
	var sha1Data = sha256.Sum256([]byte(str))
	var builder strings.Builder
	for _, bit := range sha1Data {
		var bitHex = strconv.FormatUint(uint64(bit), 16)
		if len(bitHex) == 1 {
			bitHex = "0" + bitHex
		}
		builder.WriteString(bitHex)
	}
	return builder.String()
}
