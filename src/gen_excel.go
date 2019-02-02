package main

import (
	"bufio"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"os"
	"strconv"
	"strings"
)

var distance = 1500

func main() {
	lastCategory := ""
	uidMap := map[string]string{}

	excelPath := fmt.Sprintf("D:\\poi\\poi_%d.xlsx", distance)
	xlsx, err := excelize.OpenFile(excelPath)
	if _, ok := err.(*os.PathError); ok {
		xlsx = excelize.NewFile()
		xlsx.Path = excelPath
	}

	file, _ := os.Open(fmt.Sprintf("D:\\poi\\%dpoi.txt", distance))
	defer file.Close()
	reader := bufio.NewReader(file)
	excelRow := 1
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		s := string(line)
		splitArray := strings.Split(s, "#####")
		levelOne := splitArray[0]
		levelTwo := splitArray[1]
		name := splitArray[2]
		lng := splitArray[3] //经度
		lat := splitArray[4] //纬度
		address := splitArray[9]
		uid := splitArray[10]

		sheetIndex := xlsx.GetSheetIndex(levelOne)
		if sheetIndex == 0 {
			excelRow = 1
			xlsx.NewSheet(levelOne)
			xlsx.SetCellValue(levelOne, "A"+strconv.Itoa(excelRow), "二级行业分类")
			xlsx.SetCellValue(levelOne, "B"+strconv.Itoa(excelRow), "名称")
			xlsx.SetCellValue(levelOne, "C"+strconv.Itoa(excelRow), "经度")
			xlsx.SetCellValue(levelOne, "D"+strconv.Itoa(excelRow), "纬度")
			xlsx.SetCellValue(levelOne, "E"+strconv.Itoa(excelRow), "地址")
			excelRow++
		} else {
			if levelTwo != lastCategory {
				uidMap = map[string]string{}
			}
			if _, ok := uidMap[uid]; ok {
				fmt.Println(fmt.Sprintf("%s重复", uid))
				continue
			}
			uidMap[uid] = ""

			xlsx.SetCellValue(levelOne, "A"+strconv.Itoa(excelRow), levelTwo)
			xlsx.SetCellValue(levelOne, "B"+strconv.Itoa(excelRow), name)
			xlsx.SetCellValue(levelOne, "C"+strconv.Itoa(excelRow), lng)
			xlsx.SetCellValue(levelOne, "D"+strconv.Itoa(excelRow), lat)
			xlsx.SetCellValue(levelOne, "E"+strconv.Itoa(excelRow), address)
			excelRow++
		}
	}
	xlsx.DeleteSheet("Sheet1")
	xlsx.Save()
}
