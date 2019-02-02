package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"poi"
	"time"
)

var flag = 1

func main() {
	s := &poi.Status{}
	if flag == 1 {
		s.Reset1500()
	} else {
		s.Reset5000()
	}
	s.LastLongitudePosition = s.LeftLongitude
	s.LastLatitudePosition = s.UpperLatitude
	Scrape(s)
}

func Scrape(s *poi.Status) {
	log.Println("开始抓取 ...")

	f := &os.File{}
	err := errors.New("")

	if flag == 1 {
		f, err = os.OpenFile("D:\\poi\\1500poi.txt", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0600)
	} else {
		f, err = os.OpenFile("D:\\poi\\5000poi.txt", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0600)
	}

	if err != nil {
		log.Panic("打开文件出错。", err)
	}
	defer f.Close()

	for k, v := range poi.CategoryMap {
		fmt.Println(k + "--------------------------------------------")
		levelOne := k
		levelTwoArray := v

		for _, value := range levelTwoArray {
			for s.LastLongitudePosition < s.RightLongitude {
				upper := s.UpperLatitude
				lower := s.LowerLatitude
				left := s.LastLongitudePosition
				right := left + s.LastLongitudeLength

				if right > s.RightLongitude {
					right = s.RightLongitude
				}

			NEXT_PAGE:
				if s.ApiAvailableTimes < 1 {
					log.Println("今日API次数用完了")
					os.Exit(0)
				}
				log.Printf("开始抓取[%.6f,%.6f,%.6f,%.6f], %s, 页号：%d", lower, left, upper, right, value, s.LastPageIndex+1)
				time.Sleep(time.Second / 30)
				s.ApiAvailableTimes--
				url := fmt.Sprintf("http://api.map.baidu.com/place/v2/search?output=json&page_size=20&scope=1&query=%s&bounds=%.6f,%.6f,%.6f,%.6f&ak=%s&page_num=%d",
					value, lower, left, upper, right, s.ApiKey, s.LastPageIndex+1)

				resp, err := http.Get(url)
				if nil != err {
					log.Panic("网络出现错误", err)
				} else {
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						log.Panic("读取HTTP Body数据时出错", err)
					}
					poiResponse := &poi.PoiResponse{}
					err = json.Unmarshal(body, poiResponse)
					if nil != err {
						log.Println("解析服务器返回的数据出错，1秒后重新获取", err)
						time.Sleep(time.Second)
					} else if poiResponse.Status != 0 {
						log.Printf("未能成功获取POI数据，服务器返回：%s\n", poiResponse.Message)
						time.Sleep(time.Second)
					} else {
						if poiResponse.Total == 0 { //没有获取到数据
							//扩大矩形面积
							log.Printf("未能获取到数据，区域宽度（经度）由%f调整为%f\n", s.LastLongitudeLength, s.LastLongitudeLength*2)
							s.LastLongitudePosition += s.LastLongitudeLength
							//加个判断，不能让它无限膨胀
							if s.LastLongitudeLength*2 < s.RightLongitude-s.LeftLongitude {
								s.LastLongitudeLength *= 2
							}
						} else if poiResponse.Total == 400 { //获取的数据达到400上限，可能不完整
							log.Printf("获取的数据总数达到400条，可能不完整，区域宽度（经度）由%f调整为%f\n", s.LastLongitudeLength, s.LastLongitudeLength/2)
							//缩小矩形面积
							s.LastLongitudeLength /= 2
						} else { // 获取的数据在 0 到 400之间，有效
							log.Printf("获取的数据总条数为%d,有效！\n", poiResponse.Total)

							for _, p := range poiResponse.Results {
								record := fmt.Sprintf("%s#####%s#####%s#####%f#####%f#####%s#####%s#####%s#####%s#####%s#####%s\n", levelOne, value, p.Name, p.Location.Lgt, p.Location.Lat, p.Telephone, p.Province, p.City, p.Area, p.Address, p.Uid)
								//log.Print(record)
								if _, err = f.WriteString(record); err != nil {
									panic(err)
								}
							}
							size := len(poiResponse.Results)
							if size == 20 { //表示还有下一页
								log.Printf("当前页数据条数为20条，表示不是末尾页，页号+1")
								s.LastPageIndex++
								goto NEXT_PAGE
							} else {
								log.Printf("当前页数据为%d条，是末尾页，页号重置", size)
								s.LastPageIndex = -1
								s.LastLongitudePosition += s.LastLongitudeLength
							}
						}
					}
				}
			} //抓完一行了，经度回到原点
			s.LastLongitudePosition = s.LeftLongitude
		}
		//抓完一个类别了，回到原点
		s.LastCategoryIndex++
	}
	f.Close()
}
