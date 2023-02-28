package rpc

import (
	"github.com/gin-gonic/gin"
	"xx_identifier_generate_service/tools/log"
	"xx_identifier_generate_service/tools/rpchelper"
	"xx_identifier_generate_service/worker"
)

func GetId(c *gin.Context) {

	Id, err := worker.GetId()
	if err != nil {
		// TODO 增加报警推送
		apiError(c, 10001, "Service internal error")
		return
	}

	apiSuccess(c, map[string]interface{}{
		"Id": Id,
	})
}

func GetIds(c *gin.Context) {

	count, err := rpchelper.RequestParameterInt(c, "c")
	if err != true {
		log.Error("error", "c empty: %v", err)
		apiError(c, ApiErrMsg, "c empty")
		return
	}

	if count < 1 || count > 100 {
		apiError(c, ApiErrMsg, "c limit 1 - 100")
		return
	}

	var IdList []int64
	for i := 0; i < int(count); i++ {
		Id, err := worker.GetId()
		if err != nil {
			// TODO 增加报警推送
			// 有一个错误则返回错误
			apiError(c, 10001, "Service internal error")
			return
		}

		IdList = append(IdList, Id)
	}

	apiSuccess(c, map[string]interface{}{
		"IdList": IdList,
	})
}
