package cas

import (
	"fmt"
	"regexp"
)

func getValue(line string) (string, error) {
	var valuePair = regexp.MustCompile("value=\".+\"").FindString(line)
	//不做检查，因为已经符合line表达式
	var value = regexp.MustCompile("(value=)|(\")").ReplaceAllString(valuePair, "")
	fmt.Printf("value:%s\n", value)
	return value, nil
}
