package rpc

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {

	msg := "hello ted"
	fmt.Println(msg)

	apiSuccess(c, map[string]interface{}{
		"msg": msg,
	})
}
