package poi

import (
	"fmt"
	"time"
)

// POI 位置
type PoiLocation struct {
	Lat float64 `json:"lat"` //纬度
	Lgt float64 `json:"lng"` //经度
}

// POI 信息
type Poi struct {
	Name      string      `json:"name"`
	Location  PoiLocation `json:"location"`
	Address   string      `json:"address"`
	Province  string      `json:"province"`
	City      string      `json:"city"`
	Area      string      `json:"area"`
	StreetId  string      `json:"street_id"`
	Telephone string      `json:"telephone"`
	//Detail    string      `json:"detail"`
	Uid string `json:"uid"`
}

type PoiResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Total   int    `json:"total"`
	Results []Poi  `json:"results"`
}

var CategoryMap = map[string][]string{
	"美食":   {"中餐厅", "外国餐厅", "小吃快餐店", "蛋糕甜品店", "咖啡厅", "茶座", "酒吧"},
	"酒店":   {"星级酒店", "快捷酒店", "公寓式酒店"},
	"购物":   {"购物中心", "百货商场", "超市", "便利店", "家居建材", "家电数码", "商铺", "集市"},
	"生活服务": {"通讯营业厅", "邮局", "物流公司", "售票处", "洗衣店", "图文快印店", "照相馆", "房产中介机构", "公用事业", "维修点", "家政服务", "殡葬服务", "彩票销售点", "宠物服务", "报刊亭", "公共厕所"},
	"丽人":   {"美容", "美发", "美甲", "美体"},
	"旅游景点": {"公园", "动物园", "植物园", "游乐园", "博物馆", "水族馆", "海滨浴场", "文物古迹", "教堂", "风景区"},
	"休闲娱乐": {"度假村", "农家院", "电影院", "KTV", "剧院", "歌舞厅", "网吧", "游戏场所", "洗浴按摩", "休闲广场"},
	"运动健身": {"体育场馆", "极限运动场所", "健身中心"},
	"教育培训": {"高等院校", "中学", "小学", "幼儿园", "成人教育", "亲子教育", "特殊教育学校", "留学中介机构", "科研机构", "培训机构", "图书馆", "科技馆"},
	"文化传媒": {"新闻出版", "广播电视", "艺术团体", "美术馆", "展览馆", "文化宫"},
	"医疗":   {"综合医院", "专科医院", "诊所", "药店", "体检机构", "疗养院", "急救中心", "疾控中心"},
	"汽车服务": {"汽车销售", "汽车维修", "汽车美容", "汽车配件", "汽车租赁", "汽车检测场"},
	"交通设施": {"地铁站", "地铁线路", "长途汽车站", "公交车站", "公交线路", "港口", "停车场", "加油加气站", "服务区", "收费站", "桥", "充电站", "路侧停车位"},
	"金融":   {"银行", "ATM", "信用社", "投资理财", "典当行"},
	"房地产":  {"写字楼", "住宅区", "宿舍"},
	"公司企业": {"公司", "园区", "农林园艺", "厂矿"},
	"政府机构": {"中央机构", "各级政府", "行政单位", "公检法机构", "涉外机构", "党派团体", "福利机构", "政治教育机构"},
	"出入口":  {"高速公路出口", "高速公路入口", "机场出口", "机场入口", "车站出口", "车站入口", "停车场出入口"},
	"自然地物": {"岛屿", "山峰", "水系"},

	//"金融": {"银行"},
}

var CategoryMapBak = map[string][]string{
	"美食":   {"中餐厅", "外国餐厅", "小吃快餐店", "蛋糕甜品店", "咖啡厅", "茶座", "酒吧"},
	"酒店":   {"星级酒店", "快捷酒店", "公寓式酒店"},
	"购物":   {"购物中心", "百货商场", "超市", "便利店", "家居建材", "家电数码", "商铺", "集市"},
	"生活服务": {"通讯营业厅", "邮局", "物流公司", "售票处", "洗衣店", "图文快印店", "照相馆", "房产中介机构", "公用事业", "维修点", "家政服务", "殡葬服务", "彩票销售点", "宠物服务", "报刊亭", "公共厕所"},
	"丽人":   {"美容", "美发", "美甲", "美体"},
	"旅游景点": {"公园", "动物园", "植物园", "游乐园", "博物馆", "水族馆", "海滨浴场", "文物古迹", "教堂", "风景区"},
	"休闲娱乐": {"度假村", "农家院", "电影院", "KTV", "剧院", "歌舞厅", "网吧", "游戏场所", "洗浴按摩", "休闲广场"},
	"运动健身": {"体育场馆", "极限运动场所", "健身中心"},
	"教育培训": {"高等院校", "中学", "小学", "幼儿园", "成人教育", "亲子教育", "特殊教育学校", "留学中介机构", "科研机构", "培训机构", "图书馆", "科技馆"},
	"文化传媒": {"新闻出版", "广播电视", "艺术团体", "美术馆", "展览馆", "文化宫"},
	"医疗":   {"综合医院", "专科医院", "诊所", "药店", "体检机构", "疗养院", "急救中心", "疾控中心"},
	"汽车服务": {"汽车销售", "汽车维修", "汽车美容", "汽车配件", "汽车租赁", "汽车检测场"},
	"交通设施": {"飞机场", "火车站", "地铁站", "地铁线路", "长途汽车站", "公交车站", "公交线路", "港口", "停车场", "加油加气站", "服务区", "收费站", "桥", "充电站", "路侧停车位"},
	"金融":   {"银行", "ATM", "信用社", "投资理财", "典当行"},
	"房地产":  {"写字楼", "住宅区", "宿舍"},
	"公司企业": {"公司", "园区", "农林园艺", "厂矿"},
	"政府机构": {"中央机构", "各级政府", "行政单位", "公检法机构", "涉外机构", "党派团体", "福利机构", "政治教育机构"},
	"出入口":  {"高速公路出口", "高速公路入口", "机场出口", "机场入口", "车站出口", "车站入口", "门", "停车场出入口"},
	"自然地物": {"岛屿", "山峰", "水系"},

	//"金融": {"银行"},
}

type Status struct {
	UpperLatitude         float64 `json:"upperLatitude"`         //上纬度
	LowerLatitude         float64 `json:"lowerLatitude"`         //下纬度
	LeftLongitude         float64 `json:"leftLongitude"`         //左经度
	RightLongitude        float64 `json:"rightLongitude"`        //右经度
	ApiAvailableTimes     int     `json:"apiAvailableTimes"`     //API可用次数
	LastUseDate           string  `json:"lastUseDate"`           //上次使用日期
	ApiKey                string  `json:"apiKey"`                //API Key
	LastLongitudePosition float64 `json:"lastLongitudePosition"` //上次抓取经度位置
	LastLatitudePosition  float64 `json:"lastLatitudePosition"`  //上次抓取纬度位置
	LastCategoryIndex     int     `json:"lastCategoryIndex"`     //上次抓取类别索引
	LastLongitudeLength   float64 `json:"lastLongitudeLength"`   //上次抓取矩形区域横向宽度（经度位移）
	LastLatitudeLength    float64 `json:"lastLatitudeLength"`    //上次抓取矩形区域纵向高度（纬度位移）
	LastPageIndex         int     `json:"lastPageIndex"`         //上次抓取页索引
}

/**
经度：114.490825  纬度：38.016821
1.5公里  38.030296   114.507929   38.003346  114.473721
5公里    38.061737   114.547837   37.971905  114.433813
*/
func (s *Status) Reset1500() {
	s.UpperLatitude = 38.030296
	s.LowerLatitude = 38.003346
	s.LeftLongitude = 114.473721
	s.RightLongitude = 114.507929
	s.ApiAvailableTimes = 30000
	s.LastUseDate = today()
	s.ApiKey = "7393214b3b391ac3d4679b4f4b8c698b"
	s.LastLongitudePosition = 0
	s.LastLatitudePosition = 0
	s.LastCategoryIndex = 0
	s.LastLongitudeLength = 0.001
	s.LastLatitudeLength = 0.002
	s.LastPageIndex = -1
}

func (s *Status) Reset5000() {
	s.UpperLatitude = 38.061737
	s.LowerLatitude = 37.971905
	s.LeftLongitude = 114.433813
	s.RightLongitude = 114.547837
	s.ApiAvailableTimes = 30000
	s.LastUseDate = today()
	s.ApiKey = "7393214b3b391ac3d4679b4f4b8c698b"
	s.LastLongitudePosition = 0
	s.LastLatitudePosition = 0
	s.LastCategoryIndex = 0
	s.LastLongitudeLength = 0.001
	s.LastLatitudeLength = 0.002
	s.LastPageIndex = -1
}

func today() string {
	now := time.Now()
	return fmt.Sprintf("%4d%2d%2d", now.Year(), now.Month(), now.Day())
}
