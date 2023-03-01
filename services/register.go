package services

import (
	"github.com/jinzhu/gorm"
	"github.com/tebie6/identifier_generate_service/models"
	"github.com/tebie6/identifier_generate_service/tools"
)

type RegisterService struct{}

// GetWorkerId 获取WorkerId
func (r *RegisterService) GetWorkerId() (int64, error) {

	db := models.GetDbInst()

	// 服务器IP
	ip := tools.GetOutboundIP()

	// 查询是否有注册
	register := models.CommonIdentifierGenerateRegister{}
	err := db.Where("ip = ?", ip).First(&register).Error

	// 不存在则注册
	if err == gorm.ErrRecordNotFound {
		register.Ip = ip
		err = db.Create(&register).Error
		if err != nil {
			return 0, err
		}
	}

	if err != nil {
		return 0, err
	}

	return register.Id, nil
}

// GetLastTimestamp 获取最后操作时间
func (r *RegisterService) GetLastTimestamp() (int64, error) {

	db := models.GetDbInst()

	// 服务器IP
	ip := tools.GetOutboundIP()

	// 查询是否有注册
	register := models.CommonIdentifierGenerateRegister{}
	err := db.Where("ip = ?", ip).First(&register).Error

	// 不存在则注册
	if err == gorm.ErrRecordNotFound {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return register.LastTimestamp, nil
}

// SaveLastTimestamp 更新最后操作时间
func (r *RegisterService) SaveLastTimestamp(ip string, lastTimestamp int64) error {

	db := models.GetDbInst()

	// 查询是否有注册
	register := models.CommonIdentifierGenerateRegister{}
	err := db.Where("ip = ?", ip).First(&register).Error
	if err != nil {
		return err
	}

	register.LastTimestamp = lastTimestamp
	err = db.Save(&register).Error
	if err != nil {
		return err
	}

	return nil
}
