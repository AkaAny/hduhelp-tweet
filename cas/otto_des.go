package cas

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
)

func GetRSAValue(userName string, password string, ltValue string) string {
	rawJs, err := ioutil.ReadFile("assets/des.js")
	if err != nil {
		panic(err)
	}

	vm := otto.New()

	_, err = vm.Run(string(rawJs))
	if err != nil {
		panic(err)
	}

	var inputStr = userName + password + ltValue
	//encodeInp是JS函数的函数名
	value, err := vm.Call("strEnc", nil, inputStr, "1", "2", "3")
	if err != nil {
		panic(err)
	}
	var result = value.String()
	fmt.Println(result)
	return result
}
