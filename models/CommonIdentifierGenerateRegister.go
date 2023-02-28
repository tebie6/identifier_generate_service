package models

// 注册记录表
type CommonIdentifierGenerateRegister struct {
	Id            int64  `gorm:"primary_key"`
	Ip            string `gorm:"ip"`
	LastTimestamp int64  `gorm:"last_timestamp"`
}

// 自定义表名
func (*CommonIdentifierGenerateRegister) TableName() string {
	return "common_identifier_generate_register"
}
