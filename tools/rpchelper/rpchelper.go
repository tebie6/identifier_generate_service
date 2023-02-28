package rpchelper

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"xx_identifier_generate_service/tools/log"
)

// 获取字符串值
func RequestParameterString(c *gin.Context, key string) string {
	v := c.Request.FormValue(key)
	return v
}

// 获取整数值
func RequestParameterInt(c *gin.Context, key string) (int64, bool) {
	v := RequestParameterString(c, key)

	if v == "" {
		return 0, false
	}

	i, e := strconv.ParseInt(v, 10, 64)
	if e != nil {
		return 0, false
	}

	return i, true
}

// 获取浮点值
func RequestParameterFloat(c *gin.Context, key string) (float64, bool) {
	v := RequestParameterString(c, key)

	if v == "" {
		return 0, false
	}

	i, e := strconv.ParseFloat(v, 64)
	if e != nil {
		return 0, false
	}

	return i, true
}

// 获取整数值列表
func RequestParameterIntList(c *gin.Context, key string) []int64 {

	// 声明一个map 用来存储返回数据
	data := *new([]int64)

	// 遍历表单数据
	for i := range c.Request.Form {

		// 判断 i 是否有前缀字符串prefix。
		if strings.HasPrefix(i, key+"[") {

			val, err := strconv.ParseInt(c.Request.Form.Get(i), 10, 64)
			if err != nil {
				data = append(data, 0)
				log.Error("error", "RequestParameterIntList error: %v", err)
				continue
			}
			data = append(data, val)
		}
	}

	return data
}

// 获取字符串值列表
func RequestParameterStringList(c *gin.Context, key string) map[string][]string {

	// 声明一个map 用来存储返回数据
	data := make(map[string][]string)

	// 遍历表单数据
	for i := range c.Request.Form {

		// 判断 i 是否有前缀字符串prefix。
		if strings.HasPrefix(i, key+"[") {

			// 将前缀 和 后缀 中括号[] 替换为空
			rp := strings.NewReplacer(key+"[", "", "[]", "", "]", "")

			// Replace返回 i 的所有替换进行完后的拷贝。
			data[rp.Replace(i)] = c.Request.Form[i]
		}
	}

	return data
}
