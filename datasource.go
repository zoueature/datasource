package datasource

import (
	"gitlab.jiebu.com/base/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Orm struct {
	*gorm.DB
}

// Open 打开数据库连接
func Open(cfg config.DatabaseConfig) *Orm {
	orm, err := gorm.Open(mysql.Open(cfg.Dsn()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.Prefix,
			SingularTable: true,
		},
	})
	if err != nil {
		panic("[Datasource] open database err : " + err.Error())
	}
	return &Orm{orm}
}
