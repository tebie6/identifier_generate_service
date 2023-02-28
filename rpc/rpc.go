package rpc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xx_identifier_generate_service/conf"
	"xx_identifier_generate_service/tools/log"
)

const (
	ApiSuccess = 0
	ApiErrMsg  = 500
)

type ApiRet struct {
	Code int32                  `json:"code"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"msg"`
}

func InitRpc() {

	r := gin.New()

	r = routerRegister(r)
	rpcListenAddress := fmt.Sprintf("%s:%s",
		conf.GetConfigString("rpc", "host"),
		conf.GetConfigString("rpc", "port"),
	)
	err := r.Run(rpcListenAddress)
	if err != nil {
		log.Error("", "start server failed %s", err.Error())
	}
}

func routerRegister(r *gin.Engine) *gin.Engine {

	r.GET("/", test)
	r.GET("/get_id", GetId)
	r.GET("/get_ids", GetIds)
	return r
}

func apiSuccess(c *gin.Context, data map[string]interface{}) {
	ret := &ApiRet{
		Code: ApiSuccess,
		Data: data,
		Msg:  "success",
	}

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, ret)
	return
}

func apiError(c *gin.Context, code int32, msg string) {
	ret := &ApiRet{
		Code: code,
		Msg:  msg,
	}

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, ret)
	return
}
