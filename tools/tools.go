/*
 @Title
 @Description
 @Author  Ted
 @Update  2020/7/8 11:36 上午
*/

package tools

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"
	"strings"
)

// 获取业务逻辑层面的真实IP地址
func GetClientIps(r *http.Request) string {
	headerNames := []string{
		"ali-cdn-real-ip", "cf-connecting-ip", "x-connecting-ip",
		"x_client_ip", "x-forwarded-for", "http_client_ip", "remote_addr",
	}

	for _, hName := range headerNames {
		ipAddress := r.Header.Get(hName)
		if ipAddress != "" {
			return ipAddress
		}
	}

	return r.RemoteAddr
}

// 获取业务逻辑层面的真实IP地址：切割成数组+去空
func GetClientIpArr(r *http.Request) []string {
	ips := GetClientIps(r)
	_result := strings.Split(ips, ",")

	if len(_result) == 0 {

		return _result
	}

	result := make([]string, 0)

	for _, v := range _result {
		_v := strings.TrimSpace(v)
		if _v == "" {
			continue
		}
		result = append(result, _v)
	}

	return result
}

// 获取业务逻辑层面的真实IP地址：单地址
func GetClientIp(r *http.Request) string {
	ipArr := GetClientIpArr(r)
	if len(ipArr) == 0 {
		return ""
	}

	return ipArr[0]
}

func GetCaller(skip int) string {
	_, file, line, _ := runtime.Caller(skip + 1)
	file = file[strings.LastIndex(file, "/")+1:]
	return fmt.Sprintf("%s:%d", file, line)
}

// 获取服务器IP
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println("Err: close net failed ", err)
		}
	}()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String())
	return localAddr.IP.String()
}
