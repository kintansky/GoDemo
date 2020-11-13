package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type IPRecord struct {
	ID            int
	DeviceIP      string
	IPMask        int
	Gateway       string
	DeviceName    string
	LogicPort     string
	LogicPortNum  string
	Svlan         string
	Cvlan         string
	IPDescription string
	IPType        string
	RecordTime    time.Time // 如果字段含有datetime字段，只能通过类型转换赋值，不能直接赋值到结构体，如果要直接赋值给结构体，需要定义成string
	IPFunc        string
}

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

func prepareQueryDemo() {
	sqlStr := "select id, device_ip from mr_rec_ip_record where id = ?"
	stmt, err := db.Prepare(sqlStr) // 将语句像发送给服务器解析
	if err != nil {
		fmt.Printf("Prepare failed %v\n", err)
		return
	}
	defer stmt.Close() // 需要关闭

	// 如果有多句需要执行的，需要分别准备
	sqlStr2 := "update mr_rec_ip_record set device_ip = ? where id = ?"
	stmt2, err := db.Prepare(sqlStr2) // 将语句像发送给服务器解析
	if err != nil {
		fmt.Printf("Prepare failed %v\n", err)
		return
	}
	defer stmt2.Close()

	_, err = stmt2.Exec("3.3.3.3", 1)
	if err != nil {
		fmt.Printf("Execute failed %v\n", err)
		return
	}

	rows, err := stmt.Query(1)
	if err != nil {
		fmt.Printf("query failed %v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var IP IPRecord
		err := rows.Scan(&IP.ID, &IP.DeviceIP)
		if err != nil {
			fmt.Printf("Scan err, %v\n", err)
			return
		}
		fmt.Println(IP.ID, IP.DeviceIP)
	}
}

func main() {
	initDB()
	prepareQueryDemo()
}
