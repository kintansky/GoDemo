package main

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql" // 只是用包的init()

	"database/sql"
)

var db *sqlx.DB

type Person struct {
	// 使用sqlx的时候注意需要使用反射设置tag，字段要和数据库一致，然后record_time这种日期格式只能作为string返回了
	ID    sql.NullInt64  `db:"id"`
	Name  sql.NullString `db:"name"`  // 对于可能为null的字段，只能使用这种方式，返回的是一个stuct{Srting string, valid bool},如果valid为false，则为null
	Birth sql.NullString `db:"birth"` // NullTime解析好像有问题
}

func initDB() (err error) {
	// 连接数据库
	dsn := "root:password@tcp(127.0.0.1:3306)/test"
	db, err = sqlx.Connect("mysql", dsn) // OPEN不会校验DSN的参数，校验数据在ping进行
	if err != nil {
		fmt.Printf("Open Failed %v\n", err)
		return
	}
	db.SetMaxOpenConns(20) // 设置连接池最大连接数，0为不限制
	db.SetMaxIdleConns(10) // 限制可保持的最大空闲连接数，0为不保存空闲连接
	fmt.Println("Connect Success")
	return
}

// SQLX 不在需要每个字段赋值
func queryOneRow() {
	sqlStr := `select * from table1 where id = ?`
	var p Person
	err := db.Get(&p, sqlStr, 3) // 这里必须传指针
	if err != nil {
		fmt.Printf("Get failed, %v\n", err)
		return
	}
	fmt.Printf("%v", p)
	v, err := p.Name.Value()
	if err != nil {
		fmt.Println(v)
		fmt.Println(err)
	}
	fmt.Println(v.(string))
}

func queryMultiRow() {
	sqlStr := `select * from table1 where id < ?`
	var ps []Person
	err := db.Select(&ps, sqlStr, 5) // 这里必须传指针
	if err != nil {
		fmt.Printf("Select failed %v\n", err)
		return
	}
	for _, p := range ps {
		fmt.Printf("%v", p)
	}

}

func insertData() {
	sqlStr := `insert into table1 (name, birth) values (?, ?)`
	// 构造null结构体用于插入，这个结构体内，只要Valid为false，String不管是什么值，都会被Null提带写入数据库
	data := make([]interface{}, 0)
	data = append(data, sql.NullString{String: "nullName", Valid: false}, time.Now().Format("2006-01-01 15:04:05"))
	_, err := db.Exec(sqlStr, data...)
	if err != nil {
		fmt.Println("insert error", err)
		return
	}
}

func main() {
	initDB()

	insertData()

	queryOneRow()
	queryMultiRow()

}
