package main

import (
	"context"
	"fmt"
	"go-zero-example/model"
	"go-zero-example/svc"
)

func init() {
	svc.NewServiceContext()
}

func main() {
	mysqlDemo()
}

func mysqlDemo() {
	ctx := context.Background()
	var adminListopts []model.AdminUserOption

	chatId := 5171150121
	userName := "sunguangyong"

	if chatId > 0 {
		adminListopts = append(adminListopts, svc.Ctx.AdminConn.WithChatId(int64(chatId)))
	}

	if userName != "" {
		adminListopts = append(adminListopts, svc.Ctx.AdminConn.WithName(userName))
	}

	adminListopts = append(adminListopts, svc.Ctx.AdminConn.WithOrder())

	adminList, err := svc.Ctx.AdminConn.CommonOptionFind(ctx, adminListopts...)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(adminList)

}
