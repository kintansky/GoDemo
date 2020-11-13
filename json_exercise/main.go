package main

import (
	"fmt"
	"os"
	// 同一个文件夹下的不同go文件可以直接使用，不需要导入，所以这里不用导入student
)

func showMenu() {
	fmt.Println(`
	欢迎使用
	1、添加学员
	2、编辑学员
	3、展示学员
	4、退出
	`)
}

func getInput() (string, int, int) {
	var (
		name  string
		age   int
		score int
	)
	fmt.Println("请输入学员信息")
	fmt.Println("姓名: ")
	fmt.Scanf("%s\n", &name)
	fmt.Println("年龄: ")
	fmt.Scanf("%d\n", &age)
	fmt.Println("得分: ")
	fmt.Scanf("%d\n", &score)
	return name, age, score
}

func main() {
	var input int = -1
	c1 := newClass("class01", 10)
	for {
		showMenu()
		fmt.Println("请选择序号")
		fmt.Scanf("%d\n", &input)
		fmt.Println("选择的是", input)
		switch input {
		case 1:
			name, age, score := getInput()
			c1.addStudent(name, age, score)
			fmt.Println("Success")
		case 2:
			fmt.Println("需要修改的学生编号: ")
			var studentID int
			fmt.Scanf("%d\n", &studentID)
			newName, newAge, newScore := getInput()
			c1.editStudent(studentID, newName, newAge, newScore)
			fmt.Println("Success")
		case 3:
			c1.showClass()
		case 4:
			os.Exit(0)
		}
	}

}
