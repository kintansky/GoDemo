package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	// 连接数据库
	dsn := "root:password@tcp(127.0.0.1:3306)/omni_agent"
	db, err = sql.Open("mysql", dsn) // OPEN不会校验DSN的参数，校验数据在ping进行
	if err != nil {
		fmt.Printf("Open Failed %v\n", err)
		return
	}
	// 校验数据
	err = db.Ping() // 尝试连接数据库
	if err != nil {
		fmt.Printf("Open %s failed, %v\n", dsn, err)
		return
	}
	db.SetMaxOpenConns(20) // 设置连接池最大连接数，0为不限制
	db.SetMaxIdleConns(10) // 限制可保持的最大空闲连接数，0为不保存空闲连接
	fmt.Println("Connect Success")
	return
}

func transcationDemo() {
	tx, err := db.Begin() // 开始事务
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin transcation failed, %v\n", err)
		return
	}

	sqlStr1 := `Update mr_rec_ip_record set device_ip = ? where id =1`
	_, err = tx.Exec(sqlStr1, "1.1.1.1") // 执行事务
	if err != nil {
		fmt.Printf("sql:%s, execute failed, %v\n", sqlStr1, err)
		tx.Rollback()
		return
	}

	sqlStr2 := `Update mr_rec_ip_record set device_ip = ? where id =3`
	_, err = tx.Exec(sqlStr2, "2.2.2.2") // 执行事务
	if err != nil {
		fmt.Printf("sql:%s, execute failed, %v\n", sqlStr2, err)
		tx.Rollback()
		return
	}
	err = tx.Commit() // 整体提交事务
	if err != nil {
		fmt.Printf("Commit failed, %v\n", err)
		tx.Rollback()
		return
	}
	fmt.Println("Execute success")
}

func main() {
	initDB()
	transcationDemo()

}
