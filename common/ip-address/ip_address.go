package ipaddress

import (
	"encoding/json"
	"net"
	"ruoyi-go/common/curl"
	"ruoyi-go/common/utils"

	"github.com/mileusna/useragent"
)

// ip地址
type IpAddress struct {
	Ip         string `json:"ip"`
	Pro        string `json:"pro"`
	ProCode    string `json:"proCode"`
	City       string `json:"city"`
	CityCode   string `json:"cityCode"`
	Region     string `json:"region"`
	RegionCode string `json:"regionCode"`
	Addr       string `json:"addr"`
	Browser    string `json:"browser"`
	Os         string `json:"os"`
}

// 根据ip获取地址
func GetAddress(ip string, userAgent string) *IpAddress {

	var ipAddress IpAddress

	// 解析userAgent
	userAgentData := useragent.Parse(userAgent)
	ipAddress.Browser = userAgentData.Name
	ipAddress.Os = userAgentData.OS

	var internalIp = "(((\\d)|([1-9]\\d)|(1\\d{2})|(2[0-4]\\d)|(25[0-5]))\\.){3}((\\d)|([1-9]\\d)|(1\\d{2})|(2[0-4]\\d)|(25[0-5]))$"

	if utils.CheckRegex(internalIp, ip) || ip == "127.0.0.1" || ip == "::1" {
		ipAddress.Ip = ip
		ipAddress.Addr = "内网地址"
		return &ipAddress
	}

	if netIp := net.ParseIP(ip); netIp == nil || netIp.IsLoopback() {
		ipAddress.Ip = ip
		ipAddress.Addr = "未知地址"
		return &ipAddress
	}

	body, err := curl.DefaultClient().Send(&curl.RequestParam{
		Url: "http://whois.pconline.com.cn/ipJson.jsp",
		Query: map[string]interface{}{
			"ip":   ip,
			"json": true,
		},
	})

	if err != nil {
		ipAddress.Ip = ip
		ipAddress.Addr = "未知地址"
		return &ipAddress
	}

	if err := json.Unmarshal([]byte(body), &ipAddress); err != nil {
		ipAddress.Ip = ip
		ipAddress.Addr = "未知地址"
		return &ipAddress
	}

	return &ipAddress
}
