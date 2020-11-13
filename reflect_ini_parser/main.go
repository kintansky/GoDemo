package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// mysql conf 文件解析器

// MysqlConfig 结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig 结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

// Config 因为文件包含两个配置，所以可以再做一个结构体嵌套
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0、由于还必须修改结构体的值，所以data只能传指针，
	t := reflect.TypeOf(data)
	fmt.Printf("Name %v, Kind %v\n", t.Name(), t.Kind())
	if t.Kind() != reflect.Ptr {
		// err = fmt.Errorf("data should be pointer") // 格式化输出后返回error类型
		err = errors.New("data should be pointer")
		return
	}
	// 0.1 传进来的data必须是结构体的指针
	if t.Elem().Kind() != reflect.Struct { // t是一个指针，所以要用Elem
		err = errors.New("data should be struct")
		return
	}
	// 1、读取文本
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Printf("%#v\n", lineSlice)
	// 2、每行读入
	var structName string // 结构体
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		if len(line) == 0 { // 跳过空行
			continue
		}
		// 2.1、注释跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 []开头的表示session
		if strings.HasPrefix(line, "[") { // 进入节点
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line %d syntax error", idx+1) // +1是因为人看是从第1行开始的
				return
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 { // 节点名空
				err = fmt.Errorf("line %d syntax error", idx+1)
				return
			}
			// 反射遍历结构体的字段，对比ini配置setion名，获取对应的struct
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i) // 取出字段
				if sectionName == field.Tag.Get("ini") {
					// 匹配到tag对应名
					structName = field.Name
					fmt.Printf("found %s Struct %s\n", sectionName, structName)
				}
			}
		} else { // 找到section对应的struct后，继续循环，进入section的主体内容，获取键值对
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line %d has syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到结构体的值信息
			sType := sValue.Type()                     // 拿到结构体的类型信息
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data %s should be a struct", structName)
				return
			}
			// 通过反射拿结构体的字段
			var fieldName string
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i)          // tag 信息需要从类型信息中取
				if field.Tag.Get("ini") == key { // 配置文件的key= 结构体的tag，就取出字段并准备给这个字段赋值
					// 找到字段
					fieldName = field.Name
					// break
				}
			}
			// 对for循环获取到field进行赋值
			if len(fieldName) == 0 { // 如果结构体没有对应字段就跳过
				continue
			}
			fieldObj := sValue.FieldByName(fieldName)
			fmt.Println(fieldName, fieldObj.Type().Kind())
			switch fieldObj.Type().Kind() {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line %d type error", idx+1)
					return
				}
				fieldObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line %d value type error", idx+1)
					return
				}
				fieldObj.SetBool(valueBool)
			}
		}
	}
	// 2.3 如果不是[]开头，键值对
	return
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, %v\n", err)
	}
	fmt.Printf("%#v", cfg)
}
