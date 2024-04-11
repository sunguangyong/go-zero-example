package svc

import (
	"github.com/zeromicro/go-zero/core/stores/mon"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-example/model"
)

type ServiceContext struct {
	MongoConn *mon.Model
	RedisConn *redis.Redis
	AdminConn model.AdminListModel
}

var Ctx ServiceContext

func NewServiceContext() {

	dataSource := "root:qIlRn)muX83d@tcp(82.156.56." +
		"237:3306)/telegram_bot?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"

	Ctx = ServiceContext{
		AdminConn: model.NewAdminListModel(sqlx.NewMysql(dataSource)),
	}

	//opts := mopt.ClientOptions{}
	//Username: "root",
	//Password: "123456",
	//ctx := context.Background()

	//model, err := mon.NewModel(common.Url, common.DbName, common.CollectionName)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(Ctx)
	//
	//Ctx.MongoConn = model

	//mongoConf := mongo.Config{
	//	Url:            common.Url,
	//	DbName:         common.DbName,
	//	CollectionName: common.CollectionName,
	//}
	//
	//mongoConn, err := mongo.NewMongo(mongoConf)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//return &ServiceContext{
	//	MongoConn: mongoConn,
	//}
}
