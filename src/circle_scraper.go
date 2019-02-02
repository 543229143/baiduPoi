package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"poi"
	"strconv"
	"time"
)

func main() {
	radius := 1500
	loop := 1

	for k, v := range poi.CategoryMap {
		levelOne := k
		levelTwoArray := v

		excelPath := fmt.Sprintf("D:\\poi\\poi_%s.xlsx", levelOne)
		xlsx, err := excelize.OpenFile(excelPath)
		fmt.Println(err)
		if _, ok := err.(*os.PathError); ok {
			xlsx = excelize.NewFile()
			xlsx.Path = excelPath
		}

		for _, value := range levelTwoArray {
			sheetIndex := xlsx.GetSheetIndex(value)
			if sheetIndex == 0 {
				xlsx.NewSheet(value)
			}
			excelRow := len(xlsx.GetRows(value)) + 1

			xlsx.SetCellValue(value, "A"+strconv.Itoa(excelRow), "二级行业分类")
			xlsx.SetCellValue(value, "B"+strconv.Itoa(excelRow), "经度")
			xlsx.SetCellValue(value, "C"+strconv.Itoa(excelRow), "纬度")
			xlsx.SetCellValue(value, "D"+strconv.Itoa(excelRow), "地址")

			for i := 1; i <= loop; i++ {
				log.Printf("开始抓取 %s %s, 页号：%d", levelOne, value, i)
				time.Sleep(time.Second / 2)
				url := fmt.Sprintf("http://api.map.baidu.com/place/v2/search?query=%s&location=38.016977,114.490695&radius=%d&output=json&ak=7393214b3b391ac3d4679b4f4b8c698b&page_size=20&page_num=%d",
					value, radius, i)
				resp, err := http.Get(url)
				if err != nil {
					log.Panic("网络出现错误", err)
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Panic("读取HTTP Body数据时出错", err)
				}
				poiResponse := &poi.PoiResponse{}
				err = json.Unmarshal(body, poiResponse)
				if nil != err {
					log.Println("解析服务器返回的数据出错，5秒后重新获取", err)
					time.Sleep(time.Second * 5)
				}
				if poiResponse.Total == 400 {
					log.Printf("%s %s total 超过400，结束", levelOne, value)
					break
				} else {
					log.Printf("%s %s total %d", levelOne, value, poiResponse.Total)
					for _, p := range poiResponse.Results {
						xlsx.SetCellValue(value, "A"+strconv.Itoa(excelRow), p.Name)
						xlsx.SetCellValue(value, "B"+strconv.Itoa(excelRow), p.Location.Lgt)
						xlsx.SetCellValue(value, "C"+strconv.Itoa(excelRow), p.Location.Lat)
						xlsx.SetCellValue(value, "D"+strconv.Itoa(excelRow), p.Address)
					}
					size := len(poiResponse.Results)
					if size < 20 {
						break
					}
				}
			}
			xlsx.Save()
		}
	}
}
