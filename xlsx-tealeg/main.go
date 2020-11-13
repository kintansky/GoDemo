package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

func writeXls() {
	// var file *xlsx.File
	// var sheet *xlsx.Sheet
	// var row *xlsx.Row
	// var cell *xlsx.Cell
	// var err error
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 如果不新增row会一直使用这个row写入，cell同理
	for i := 0; i < 20; i++ {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = "test_" + strconv.Itoa(i)
		// cell.SetDateTime(time.Now())	// 这样能写入时间，但使用默认UTC时区
		cell = row.AddCell()
		loc, _ := time.LoadLocation("Asia/Shanghai")
		dateOption := xlsx.DateTimeOptions{
			Location:        loc,
			ExcelTimeFormat: xlsx.DefaultDateTimeFormat,
		}
		cell.SetDateWithOptions(time.Now(), dateOption)

	}
	// 继续写入第二个sheet
	sheet, err = file.AddSheet("sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = "sheet_" + strconv.Itoa(i)
		// cell.SetDateTime(time.Now())	// 这样能写入时间，但使用默认UTC时区
		cell = row.AddCell()
		loc, _ := time.LoadLocation("Asia/Shanghai")
		dateOption := xlsx.DateTimeOptions{
			Location:        loc,
			ExcelTimeFormat: xlsx.DefaultDateTimeFormat,
		}
		cell.SetDateWithOptions(time.Now(), dateOption)

	}

	if err := file.Save("book.xlsx"); err != nil {
		fmt.Println(err)
	}

}

func readXls() {
	file, err := xlsx.OpenFile("book.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// for _, sheet := range file.Sheets{

	// }
	s1 := file.Sheets[0]
	for _, row := range s1.Rows {
		for _, cell := range row.Cells {
			s, err := cell.FormattedValue() // 这样会使用cell的格式来解析数据，特别是时间数据
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%s\t", s)
		}
		fmt.Println()
	}

}

func main() {
	writeXls()
	readXls()

	// 可以直接获取真个book的数据
	myslice, _ := xlsx.FileToSlice("book.xlsx")
	fmt.Printf("%v\n", myslice)
}
