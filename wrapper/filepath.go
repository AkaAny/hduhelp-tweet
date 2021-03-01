package wrapper

import "path/filepath"

var PWD string = "." //默认使用系统PWD指定的环境变量

func SetPWD(path string) {
	PWD = path
}

func GetPath(path string) string {
	return filepath.Join(PWD, path)
}
