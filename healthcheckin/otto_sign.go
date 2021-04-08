package healthcheckin

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func getSign(id string, model PostModel) (string, error) {
	encodedStaffID, err := getEncodeValue(id)
	if err != nil {
		return "", err
	}
	fmt.Println(encodedStaffID)
	encodedProvince, err := getEncodeValue(model.Province)
	if err != nil {
		return "", err
	}
	fmt.Println(encodedProvince)
	var rawSign = fmt.Sprintf("%s%s%d%s%s%s",
		model.Name, encodedStaffID,
		model.TimeStamp,
		encodedProvince, model.City, model.Country)
	fmt.Println(rawSign)
	var sign = sha1Hex(rawSign)
	fmt.Println(sign)
	return sign, nil
}

func getEncodeValue(str string) (string, error) {
	var encoded = url.QueryEscape(str)
	decoded, err := url.QueryUnescape(encoded)
	if err != nil {
		return "", err
	}
	resultBuf := bytes.NewBuffer(nil)
	var writer = base64.NewEncoder(base64.URLEncoding, resultBuf)
	_, err = writer.Write([]byte(decoded))
	if err != nil {
		return "", err
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}
	return string(resultBuf.Bytes()), nil
}

func sha1Hex(str string) string {
	var sha1Data = sha1.Sum([]byte(str))
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
