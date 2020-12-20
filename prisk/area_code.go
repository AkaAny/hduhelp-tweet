package prisk

import (
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
)

func (pr PRisk) GetAreaList() []Area {
	var request = pr.GetRequest()
	var url = "http://bmfw.www.gov.cn/myqfxdjcx/source/WAP/js/area.js"
	resp, body, errs := request.Get(url).End()
	fmt.Println(resp.StatusCode)
	if errs != nil {
		panic(errs[0])
	}
	err := ioutil.WriteFile("area.js", []byte(body), 0644)
	if err != nil {
		panic(err)
	}
	rawJson := mustParseArea(body)
	var areas []Area
	err = json.Unmarshal([]byte(rawJson), &areas)
	return areas
}

func mustParseArea(areaJS string) string {
	vm := otto.New()

	_, err := vm.Run(string(areaJS))
	if err != nil {
		panic(err)
	}

	value, err := vm.Eval("JSON.stringify(area)")
	if err != nil {
		panic(err)
	}
	var result = value.String()
	fmt.Println(result)
	return result
}
