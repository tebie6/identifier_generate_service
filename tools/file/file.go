package file

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// GetFileContentToInt64 获取文件内容转Int64
func GetFileContentToInt64(filePath string) (int64, error) {

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	str := string(content)
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)

	if len(str) == 0 {
		return 0, nil
	}

	return strconv.ParseInt(str, 10, 64)
}

// GetFileContentToString 获取文件内容转String
func GetFileContentToString(filePath string) (string, error) {

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	str := string(content)
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)

	return str, nil
}
