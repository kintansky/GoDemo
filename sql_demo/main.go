package main

import (
	"database/sql" // 所有数据库必须引入
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // 只是用包的init()
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

// 查询一行
func queryRow() {
	sqlStr := `select id, device_ip, ip_mask from mr_rec_ip_record where id = ?`
	var IP IPRecord
	// 必须调用Scan，因为Scan后会释放连接回连接池，否则连接一直被占用，在连接池满后无法进行下一次查询
	err := db.QueryRow(sqlStr, 1).Scan(&IP.ID, &IP.DeviceIP, &IP.IPMask)
	if err != nil {
		fmt.Printf("Scan Failed %v\n", err) // 这部分或获取执行SQL过程中的错误
		return
	}
	fmt.Printf("%#v", IP)
}

// 查询多行
func queryMultiRow() {
	sqlStr := `select id, record_time from mr_rec_ip_record where id < ?`
	// 如果SQL含有datetime字段，可以通过两种办法取出，
	// 1、结构体对应字段当成string，可以直接取出赋值即可
	// 2、结构体需要time.Time类型的，先赋值给一个临时string变量，再转换
	rows, err := db.Query(sqlStr, 5) // 返回的是rows
	if err != nil {
		fmt.Printf("Query Failed, %v\n", err)
		return
	}
	defer rows.Close() // 一定要释放连接

	// // 循环读取
	// for rows.Next() {
	// 	var IP IPRecord
	// 	err := rows.Scan(&IP.ID, &IP.RecordTime)
	// 	if err != nil {
	// 		fmt.Printf("Scan Failed, %v\n", err)
	// 		return
	// 	}
	// 	fmt.Printf("ID:%d, record_time: %v\n", IP.id, IP.record_time)
	// }

	// 循环读取
	var recordTimeStr string
	loc, err := time.LoadLocation("Asia/Shanghai")
	for rows.Next() {
		var IP IPRecord
		err := rows.Scan(&IP.ID, &recordTimeStr)
		IP.RecordTime, err = time.ParseInLocation("2006-01-02 15:04:05", recordTimeStr, loc)
		if err != nil {
			fmt.Printf("DataTime Parse Failed, %v\n", err)
		}
		if err != nil {
			fmt.Printf("Scan Failed, %v\n", err)
			return
		}
		fmt.Printf("ID:%d, record_time: %v\n", IP.ID, IP.RecordTime)
	}
}

// 执行
func executeSQL() {
	sqlStr := `update mr_rec_ip_record set device_ip = "101.127.255.0" where id = ?` // 其他操作都使用exec函数进行
	ret, err := db.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("insert failed ,%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 可以拿出受影响的行数
	if err != nil {
		fmt.Printf("get row affetced failed, %v\n", err)
	}
	fmt.Printf("%d row affected\n", n)
}

func main() {
	initDB()
	defer db.Close()
	// queryRow()
	// queryMultiRow()
	executeSQL()

}
