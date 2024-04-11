package main

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//db, err := sql.Open("mysql", "root:qIlRn)muX83d@tcp(82.156.56.237:3306)/telegram_bot")
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()

	// 创建 Squirrel 查询构建器
	//var Squirrel squirrel.Squirrel

	qb := squirrel.Select("*")
	//	From("users").
	//	Where("age = 25")
	//Where(squirrel.Eq{"age": 25})

	// 添加另一个 Where 条件
	//qb := squirrel.SelectBuilder{}
	qb = qb.Where("status = active")
	qb = qb.Where("or status1 = active")

	// 生成 SQL 查询和参数
	sql, _, err := qb.ToSql()

	fmt.Println(sql, err)

	//// 执行查询
	//rows, err := db.Query(sql, args...)
	//if err != nil {
	//	panic(err)
	//}
	//defer rows.Close()
	//
	//// 处理查询结果
	//for rows.Next() {
	//	var id int
	//	var name string
	//	var age int
	//	err := rows.Scan(&id, &name, &age)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	//}
	//
	//if err := rows.Err(); err != nil {
	//	panic(err)
	//}
}

type Option func(qb *squirrel.SelectBuilder) *squirrel.SelectBuilder

func WithName(name string) Option {
	return func(qb *squirrel.SelectBuilder) *squirrel.SelectBuilder {
		qb.Where("name = ?", name)
		return qb
	}
}

func WithAge(age int64) Option {
	return func(qb *squirrel.SelectBuilder) *squirrel.SelectBuilder {
		qb.Where("age = ?", age)
		return qb
	}
}
