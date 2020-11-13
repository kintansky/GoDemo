package main

import (
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func readXlsx() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 获取工作表中指定单元格的值
	cell := f.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)

	// 逐行逐列读取
	rows, err := f.Rows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		row := rows.Columns() // 返回一个[]string,所以返回的都是string类型
		for _, colCell := range row {
			fmt.Printf("(%T)%v\t", colCell, colCell)
		}
		fmt.Println()
	}
}

func writeXlsx() {
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet2")
	// 写一个单元格，sheet必须存在 ,注意这里可接收的数据格式
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)

	f.SetCellValue("sheet1", "A2", time.Now()) // 如果是时间格式，会默认使用date format is m/d/yy h:mm
	// 也可以通过数字来设置时间,但需要设置端元个格式为22
	f.SetCellValue("Sheet1", "B3", 42920.5)
	f.SetCellValue("Sheet1", "C3", 42920.5)
	style, err := f.NewStyle(`{"number_format": 22}`)
	if err != nil {
		fmt.Println(err)
	}
	f.SetCellStyle("Sheet1", "B3", "D3", style) // 设置一行的修改B3 D3范围即可
	f.SetCellValue("Sheet1", "D3", 42920.5)
	// 设置一列样式使用 f.SetColStyle("Sheet1", "H", style)
	// 如果设置多列err = f.SetColStyle("Sheet1", "C:F", style)

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err.Error())
	}

}

// 批量按行写入
func streamWriteXls() {
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err := file.SaveAs("Book2.xlsx"); err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	writeXlsx()
	readXlsx()
}
