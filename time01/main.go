package main

import (
	"fmt"
	"time"
)

func sleeper(loopInterval uint) {
	endTime := time.Now().Add(time.Duration(loopInterval) * time.Minute)
	for t := range time.Tick(1 * time.Second) {
		if t.After(endTime) {
			break
		}
		fmt.Printf("\r%s", t.Format("2006/01/02 15:04:05"))
	}
	fmt.Println()
}

func main() {
	//sleeper(1)
	now := time.Now()
	h := 7
	activationWeekday := time.Thursday
	days := int(activationWeekday)-int(now.Weekday())
	if days < 0 {
		days +=7
	}
	nextActivationDate := now.AddDate(0, 0, days) // 生成激活日期
	d := -1 * ((now.Hour()-h)*60*60 + now.Minute()*60 + now.Second())                  // 生成激活时间至设定时间
	nextActivationTime := nextActivationDate.Add(time.Duration(d) * time.Second)
	
	fmt.Println(nextActivationTime)
	
	//now := time.Now()
	//fmt.Printf("%v, %T\n", now, now)
	//fmt.Println(now.Hour())
	//year := now.Year()
	//month := now.Month()
	//day := now.Day()
	//hour := now.Hour()
	//minute := now.Minute()
	//second := now.Second()
	//fmt.Println(year, month, day, hour, minute, second)
	//fmt.Println("interval")
	//timeStamp1 := now.Unix() // 单位s
	//timeStampAdd1min := now.Add(1 * time.Minute).Unix()
	//fmt.Println(timeStampAdd1min - timeStamp1)

	//fmt.Println(timeStamp1)
	//timeStamp2 := now.UnixNano() // 单位ns
	//fmt.Println(timeStamp2)

	// 换算
	//t := time.Unix(0, 1580477422142595600)
	//fmt.Println(t)

	// 注意一点：如果提前定义了n传进去的话，因为类型不一样，所以需要类型转换
	//n := 0 // n 的类型是 int
	// time.Sleep(n) // cannot use n (type int) as type time.Duration in argument to time.Sleep传参的类型不匹配
	//time.Sleep(time.Duration(n) * time.Second) // 默认单位是ns，所以需要指定单位
	// 而这种情况则不需要类型转换
	//time.Sleep(2 * time.Second)

	// 常用操作
	//fmt.Println(now)
	//t2 := now.Add(-1 * time.Hour) // 增加或减少1h
	//fmt.Println(now.After(t2))
	//fmt.Println(t2)
	// 注意Sub方法是计算两个时间的差值，返回time.duration
	//d := t2.Sub(now)
	//fmt.Println(d)

	// 定时器
	//for tmp := range time.Tick(1 * time.Second) {
	//	fmt.Printf("\r%v", tmp)
	//}

	// 时间格式化
	//ret1 := now.Format("2006-01-02 15:04:05") // 注意Go语言中的时间格式，分隔符可以变化，但是这个标准时间不能错，可以是上下午后面加PM
	//fmt.Println(ret1)

	// 解析时间字符串
	//loc, err := time.LoadLocation("Asia/Shanghai") // 解析时间最好先加载时区，默认会使用UTC
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(loc)
	//timeStr := "2020/01/01 13:00:30"
	//t3, err := time.Parse("2006/01/02 15:04:05", timeStr)
	//if err != nil {
	//	fmt.Println("time parse failed")
	//	return
	//}
	//fmt.Printf("%v, %T\n", t3, t3) // 默认解析成UTC时区2020-01-01 13:00:30 +0000 UTC
	// 带上时区
	//t3, err = time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	//if err != nil {
	//	fmt.Println("time parse failed")
	//	return
	//}
	//fmt.Printf("%v, %T\n", t3, t3) //解析成了国内时间2020-01-01 13:00:30 +0800 CST

}
