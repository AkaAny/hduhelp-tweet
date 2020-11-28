package zf

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/url"
)

func toGBK(s []byte) ([]byte, error) {
	var reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	result, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func toGBKUrlEncoded(str string) (string, error) {
	gbkData, err := toGBK([]byte(str))
	if err != nil {
		return "", err
	}
	var result = url.QueryEscape(string(gbkData))
	fmt.Println(result)
	return result, nil
}
