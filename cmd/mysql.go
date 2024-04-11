package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type AdminList struct {
	Id         uint64    `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	Uid        int64     `gorm:"column:uid;type:bigint(20);default:0;comment:币虎用户id;NOT NULL" json:"uid"`
	ChatId     int64     `gorm:"column:chat_id;type:bigint(20);default:0;comment:tg_id;NOT NULL" json:"chat_id"`
	GroupId    int64     `gorm:"column:group_id;type:bigint(20);default:0;comment:群id;NOT NULL" json:"group_id"`
	GroupName  string    `gorm:"column:group_name;type:varchar(255);comment:群名称;NOT NULL" json:"group_name"`
	Permission string    `gorm:"column:permission;type:varchar(50);comment:权限;NOT NULL" json:"permission"`
	UserName   string    `gorm:"column:user_name;type:varchar(255);comment:tg user_name;NOT NULL" json:"user_name"`
	FirstName  string    `gorm:"column:first_name;type:varchar(255);comment:tg first_name;NOT NULL" json:"first_name"`
	LastName   string    `gorm:"column:last_name;type:varchar(255);comment:tg last_name;NOT NULL" json:"last_name"`
	IsDelete   int       `gorm:"column:is_delete;type:int(11);default:0;comment:删除标识 0 未删除 1 已删除;NOT NULL" json:"is_delete"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP;comment:修改时间;NOT NULL" json:"update_time"`
	Ctime      time.Time `gorm:"column:ctime;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"ctime"`
	Mtime      time.Time `gorm:"column:mtime;type:datetime;default:CURRENT_TIMESTAMP;comment:修改时间;NOT NULL" json:"mtime"`
}

func (AdminList) TableName() string {
	return "admin_list"
}

func main() {
	dsn := "root:qIlRn)muX83d@tcp(82.156.56." +
		"237:3306)/telegram_bot?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"

	//dsn := "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = db.Session(&gorm.Session{DryRun: true})

	var user AdminList
	var userarry []AdminList

	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		tx.Model(&user).Where("name = 2")
		return tx.Model(&user).Where("id = ?", 100).Limit(10).Order("age desc").Find(&userarry)
	})
	fmt.Println(sql)

	db.Where("SELECT name FROM users WHERE id = 1").Limit(5)

	stmt := db.Session(&gorm.Session{DryRun: true}).Statement
	a := stmt.SQL.String()
	fmt.Println("aaaaaaaaaa", a)

	var adminList AdminList

	condition := "id > 1"

	//query := db.Where(condition).Statement
	//sql1 := query.SQL.String()
	//
	//fmt.Println("Condition:", condition)
	//fmt.Println("Query Object:", query)
	//fmt.Println("Generated SQL:", sql)

	tx := db.Where(condition).Statement.SQL
	fmt.Println("tx=====", tx)
	fmt.Println(adminList)

	// 构建查询并返回原生 SQL
	//query := db.Where(condition).Statement
	//sql := query.SQL.String()
	//fmt.Println("sql ===== ", sql)

}
