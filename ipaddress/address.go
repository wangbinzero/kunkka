package ipaddress

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
)

func AnalysisIP(address string) {
	db, err := geoip2.Open("./GeoLite2-City.mmdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	ip := net.ParseIP(address)
	record, err := db.City(ip)
	if err != nil {
		fmt.Println("查询地址出错: ", err)
		return
	}

	fmt.Println("城市名称: ", record.City.Names["zh-CN"])
	fmt.Println("英文名称: ", record.Subdivisions[0].Names["zh-CN"])
	fmt.Println("国家名称: ", record.Country.Names["zh-CN"])
}

func GetAddress() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.interfaces failed,err: ", err.Error())
		return ""
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						fmt.Println(ipnet.IP.String())
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}
