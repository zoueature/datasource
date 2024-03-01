package datasource

import (
	"context"
	"github.com/qiniu/qmgo"
	"github.com/zoueature/config"
)

type MongoCli struct {
	*qmgo.QmgoClient
	defaultDb string
}

func (o MongoCli) Database(db ...string) *qmgo.Database {
	selectDb := o.defaultDb
	if len(db) > 0 {
		selectDb = db[0]
	}
	return o.Client.Database(selectDb)
}

// OpenMongo 打开mogo链接
func OpenMongo(cfg config.DatabaseConfig) *MongoCli {
	client, err := qmgo.Open(context.Background(), &qmgo.Config{
		Uri: cfg.Dsn(),
	})
	if err != nil {
		panic("[Datasource] init mongo client error: " + err.Error())
	}
	return &MongoCli{
		QmgoClient: client,
		defaultDb:  cfg.Database,
	}
}
