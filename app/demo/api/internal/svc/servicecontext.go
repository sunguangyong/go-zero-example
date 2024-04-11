package svc

import (
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go-zero-example/app/demo/api/internal/config"
	"go-zero-example/common"
)

type ServiceContext struct {
	Config    config.Config
	MongoConn *mon.Model
}

func NewServiceContext(c config.Config) *ServiceContext {
	mongoConn := mon.MustNewModel(common.Url, common.DbName, common.CollectionName)

	return &ServiceContext{
		Config:    c,
		MongoConn: mongoConn,
	}
}
