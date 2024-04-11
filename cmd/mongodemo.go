package main

import (
	"context"
	"fmt"
	"go-zero-example/svc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type plcData struct {
	Did     string `json:"did"`
	Utime   string `json:"utime"`
	Content []struct {
		Pid   string `json:"pid"`
		Type  string `json:"type"`
		Addr  string `json:"addr"`
		Addrv string `json:"addrv"`
		Ctime string `json:"ctime"`
	} `json:"content"`
}

func NewPidPipLine(ctx context.Context, pid, did string, pageNumber,
	pageSize int64) (pipeline mongo.Pipeline) {

	pipeline = mongo.Pipeline{
		{{"$match", bson.M{
			"did":         did,
			"content.pid": pid,
		}}},
		{{"$project", bson.M{
			"did":   1,
			"utime": 1,
			"content": bson.M{
				"$filter": bson.M{
					"input": "$content",
					"as":    "c",
					"cond":  bson.M{"$eq": []interface{}{"$$c.pid", pid}},
				},
			},
		}}},
		{{"$sort", bson.M{"utime": -1}}},
		//{{"$skip", skip}},
		//{{"$limit", pageSize}},
	}

	if pageNumber > 0 && pageSize > 0 {
		// 计算跳过的文档数量
		skip := (pageNumber - 1) * pageSize
		pipeline = append(pipeline, bson.D{
			{"$skip", skip},
		})

		pipeline = append(pipeline, bson.D{
			{"$limit", pageSize},
		})

		//{"$skip", skip},
		//{"$limit", pageSize},
	}
	return pipeline
}

func main() {
	svc.NewServiceContext()

	//logx.SetLevel(logx.InfoLevel)
	//
	//resetConf := rest.RestConf{}
	//server := rest.MustNewServer(resetConf)
	//
	svc.NewServiceContext()

	//data := `{"did":"","utime":"2024/04/09 10:38:10","content":[{"pid":"02","type":"0","addr":"电流","addrv":"113",
	//"ctime":"2024/04/09 10:38:10"}]}`
	var plc plcData
	ctx := context.Background()

	//	json.Unmarshal([]byte(data), &plc)
	//
	result, err := svc.Ctx.MongoConn.InsertOne(ctx, plc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.InsertedID)

	pipeline := NewPidPipLine(ctx, "01", "FHC230980939", 1, 1)

	var dataInfo []plcData

	err = svc.Ctx.MongoConn.Aggregate(ctx, &dataInfo, pipeline)

	session, err := svc.Ctx.MongoConn.StartSession()
	defer session.EndSession(ctx)

	err = session.StartTransaction()
	if err != nil {

	}

	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		_, err := svc.Ctx.MongoConn.InsertOne(sessionContext, "")
		if err != nil {
		}

		return err

	})

	if err != nil {
		session.AbortTransaction(context.Background())
	}

	session.CommitTransaction(ctx)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dataInfo)

	//server.Start()

}
