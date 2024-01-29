package datasource

import (
	"github.com/zoueature/config"
	"testing"
)

func TestMongo(t *testing.T) {
	cli := OpenMongo(config.DatabaseConfig{
		Type:     "mongo",
		Username: "",
		Password: "",
		Database: "test",
		Nodes: []string{
			"192.168.1.202:27017",
		},
	})
	t.Log(cli.Database("test").Collection("test").())
}
