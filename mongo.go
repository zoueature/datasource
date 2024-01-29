package datasource

import (
	"context"
	"github.com/zoueature/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCli struct {
	*mongo.Client
	defaultDb string
}

func (o MongoCli) Database(db ...string) *mongo.Database {
	selectDb := o.defaultDb
	if len(db) > 0 {
		selectDb = db[0]
	}
	return o.Client.Database(selectDb)
}

// OpenMongo 打开mogo链接
func OpenMongo(cfg config.DatabaseConfig) *MongoCli {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.Dsn()))
	if err != nil {
		panic("[Datasource] init mongo client error: " + err.Error())
	}
	return &MongoCli{
		Client:    client,
		defaultDb: cfg.Database,
	}
}
