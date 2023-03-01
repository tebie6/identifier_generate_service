package worker

import (
	"errors"
	"fmt"
	"github.com/tebie6/identifier_generate_service/services"
	"github.com/tebie6/identifier_generate_service/tools"
	"github.com/tebie6/snowflake"
	"time"
)

var (
	// 工作节点
	Worker *snowflake.Worker

	// 注册表service
	registerService = services.RegisterService{}
)

// InitWorker 初始化节点
func InitWorker() {

	// 获取workerId
	workerId := getWorkerId()

	// 实例化节点
	worker, err := snowflake.NewWorker(workerId)
	if err != nil {
		panic(errors.New(fmt.Sprintf("NewWorker error:%s", err)))
	}

	Worker = worker

	// 检验时间 - 防止服务当次时间小于上次服务最后操作时间
	if err = checkTime1(); err != nil {
		panic(errors.New(fmt.Sprintf("checkTime error:%s", err)))
	}

	// defer语句中的变量，在defer声明时就决定了。所以没有办法用defer保存最后操作时间
	// 保存最后操作时间
	go saveLastTimestamp()

}

// saveLastTimestamp 保存最后操作时间
func saveLastTimestamp() {

	// 服务器IP
	ip := tools.GetOutboundIP()

	for {
		if Worker.GetLastTimestamp() == 0 {
			continue
		}

		err := registerService.SaveLastTimestamp(ip, Worker.GetLastTimestamp())
		if err != nil {
			// TODO 报警推送
			fmt.Println("SaveLastTimestamp error:", err)
		}

		time.Sleep(10 * time.Second)
	}
}

// GetId 获取ID
func GetId() (int64, error) {

	// 检验时间 - 防止服务运行中时间回拨
	if err := checkTime(); err != nil {
		return 0, err
	}

	Id := Worker.GetId()

	//fmt.Println(Worker.GetWorkerId(), Worker.GetLastTimestamp())

	return Id, nil
}

// getWorkerId 获取workerId
func getWorkerId() int64 {

	workerId, err := registerService.GetWorkerId()
	if err != nil {
		panic(errors.New(fmt.Sprintf("getWorkerId error:%s", err)))
	}

	return workerId
}

// checkTime 验证时间-防止时间回拨
func checkTime() error {

	now := time.Now().UnixMicro()
	if now >= Worker.GetLastTimestamp() {
		return nil
	}

	// 如果当前时间和最后操作时间相差大于3秒则报警推送
	if (Worker.GetLastTimestamp() - now) > 3000 {
		// TODO 增加报警推送
		fmt.Println("报警推送 system time exception")
		return errors.New("system time exception")
	}

	// 等待当前时间的误差
	for now <= Worker.GetLastTimestamp() {
		now = time.Now().UnixMicro()
	}

	return nil
}

// checkTime1 验证时间-启动服务时调用
func checkTime1() error {

	now := time.Now().UnixMicro()
	lastTimestamp, err := registerService.GetLastTimestamp()
	if err != nil {
		return err
	}

	// 在原有基础上加30秒用来抵消同步延迟问题
	lastTimestamp = lastTimestamp + 30*1000
	if now < lastTimestamp {
		return errors.New("system time exception")
	}

	return nil
}
