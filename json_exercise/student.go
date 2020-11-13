package main

import "fmt"

// Student 结构体
type Student struct {
	ID    int
	Name  string
	Age   int
	Score int
}

// Student构造函数
func newStudent(id int, name string, age int, score int) *Student {
	return &Student{
		ID:    id,
		Name:  name,
		Age:   age,
		Score: score,
	}
}

// Class 结构体
type Class struct {
	Title    string
	Students []*Student
}

// 空的Class构造函数
func newClass(title string, studentCnt int) *Class {
	return &Class{
		Title:    title,
		Students: make([]*Student, 0, studentCnt),
	}
}

// 添加Student方法
func (c *Class) addStudent(name string, age int, score int) {
	s := newStudent(len(c.Students)+1, name, age, score)
	c.Students = append(c.Students, s)
}

// 编辑Student方法
func (c *Class) editStudent(studentID int, newName string, newAge int, newScore int) {
	for _, stdnt := range c.Students {
		if stdnt.ID == studentID {
			if newName != "" {
				stdnt.Name = newName
			}
			if newAge > 0 {
				stdnt.Age = newAge
			}
			if newScore > 0 {
				stdnt.Score = newScore
			}
			break
		}
	}
}

// 显示Student
func (c *Class) showClass() {
	fmt.Printf("Class Name: %s\n", c.Title)
	for _, stdnt := range c.Students {
		fmt.Printf("id: %02d, name: %s, age: %d, score: %d\n", stdnt.ID, stdnt.Name, stdnt.Age, stdnt.Score)
	}
}

// 删除Student
func (c *Class) deleteStudent(studentID int) {
	tempIndex := 0
	for index, stdnt := range c.Students {
		if stdnt.ID == studentID {
			tempIndex = index
			break
		}
	}
	c.Students = append(c.Students[0:tempIndex], c.Students[tempIndex+1:]...)
}

// func main() {
// 	c1 := newClass("class01", 10)
// 	for i := 0; i < 10; i++ {
// 		c1.addStudent(fmt.Sprintf("student%02d", i), 20+i, 100-i)
// 	}
// 	c1.showClass()
// 	c1.editStudent(9, "", 0, 100)
// 	c1.showClass()
// 	c1.deleteStudent(2)
// 	c1.showClass()
// }
