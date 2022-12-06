package ip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Info struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

// GetPublicIP 通过访问https://myexternalip.com/raw获取公网ip
func GetPublicIP() string {
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)
}

// OnlineIpInfo 通过ip-api.com接口查询IP信息
// 返回：IP地址的信息（格式：字符串的json）
func OnlineIpInfo(ip string) string {
	resp, err := http.Get(fmt.Sprintf("https://ip-api.com/json/%s?lang=zh-CN", ip))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	var result Info
	if err = json.Unmarshal(out, &result); err != nil {
		return ""
	}
	return result.Country + result.RegionName + result.City
}
