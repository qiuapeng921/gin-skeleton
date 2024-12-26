package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	// 检查是否有文件作为参数传入
	if len(os.Args) < 2 {
		fmt.Println("请将 Excel 文件拖拽到该程序运行。")
		return
	}

	// 获取拖拽的文件路径
	excelFilePath := os.Args[1]

	// 检查文件是否存在
	if _, err := os.Stat(excelFilePath); os.IsNotExist(err) {
		log.Fatalf("文件不存在: %s\n", excelFilePath)
	}

	// 打开 Excel 文件
	f, err := excelize.OpenFile(excelFilePath)
	if err != nil {
		log.Fatalf("无法打开 Excel 文件: %v\n", err)
	}

	// 获取所有的 sheet 名称
	sheetNames := f.GetSheetList()

	// 打开或创建一个 txt 文件
	outputFilePath := "sheets.txt"
	file, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("无法创建文件: %v\n", err)
	}
	defer file.Close()

	// 写入 sheet 名称到 txt 文件
	writer := bufio.NewWriter(file)
	for _, sheetName := range sheetNames {
		_, err := writer.WriteString(sheetName + "\n")
		if err != nil {
			log.Fatalf("写入文件失败: %v\n", err)
		}
	}
	writer.Flush()

	fmt.Printf("工作表名称已成功保存到 %s\n", outputFilePath)
}
