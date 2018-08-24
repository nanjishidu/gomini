package gomini

import (
	"fmt"
	"strconv"
	"strings"
)

type IpNetmask struct {
	Ipaddr    string //IP地址
	Mask      string //IP掩码
	Subnet    string //网络地址
	Broadcast string //广播地址
	IpStart   string //起始IP
	IpEnd     string //结束IP
}

// 如要将一B类IP地址为168.195.0.0的网络划分成若干子网，要求每个子网内有主机数为700台，
// 则该子网掩码的计算方法如下（也是对应以上各基本步骤）：
// 第1步，首先将子网中要求容纳的主机数“700”转换成二进制，得到1010111100。
// 第2步，计算出该二进制的位数为10位，即n = 10
// 第3步，将255.255.255.255从后向前的10位全部置“0”,得到的二进制数为“11111111.11111111.11111100.00000000”，
// 转换成十进制后即为255.255.252.0，这就是该要划分成主机数为700的B类IP地址 168.195.0.0的子网掩码。

func GetIpMask(start, end string) string {
	sip, eip := Ip2long(start), Ip2long(end)
	s := strconv.FormatInt(int64(eip-sip), 2)
	return start + "/" + GetIntStr(32-len(s))
}

func GetIpAddrs(ipaddr, mask string, ok bool) (ipNetmask *IpNetmask, ipRangeAddrs []string) {
	var netmask uint32
	if strings.Contains(mask, ".") {
		netmask = Ip2long(mask)
	} else {
		umask, _ := GetStrUint(mask)
		netmask = 0xffffffff << (32 - umask)
	}
	ipNetmask = new(IpNetmask)
	if Long2ip(netmask) == "255.255.255.254" || Long2ip(netmask) == "255.255.255.250" {
		//如果mask是 31 ,30，退出
		return
	}
	if Long2ip(netmask) == "255.255.255.255" {
		//如果mask是 32，取ip地址
		if ok {
			ipRangeAddrs = append(ipRangeAddrs, ipaddr)
		}
		return
	}

	network := (Ip2long(ipaddr)) & netmask
	broadcast := network | (^netmask)
	ipNetmask.Ipaddr = ipaddr
	ipNetmask.Subnet = Long2ip(network)
	ipNetmask.Mask = Long2ip(netmask)
	ipNetmask.Broadcast = Long2ip(broadcast)
	ipNetmask.IpStart = Long2ip(network + 1)
	ipNetmask.IpEnd = Long2ip(broadcast - 1)
	if ok {
		for i := network + 1; i < broadcast; i++ {
			ipRangeAddrs = append(ipRangeAddrs, Long2ip(i))

		}
	}
	return

}

// 判断IP地址是否为预留ip
func IsPrivateIP(ip string) bool {
	ip2long := Ip2long(ip)
	net_a := Ip2long("10.255.255.255") >> 24
	net_b := Ip2long("172.31.255.255") >> 20
	net_c := Ip2long("192.168.255.255") >> 16
	return ip2long>>24 == net_a || ip2long>>20 == net_b || ip2long>>16 == net_c
}

func Ip2long(ipstr string) (ip uint32) {

	ips := strings.Split(ipstr, ".")
	if len(ips) != 4 {
		return
	}
	ip1, _ := strconv.Atoi(ips[0])
	ip2, _ := strconv.Atoi(ips[1])
	ip3, _ := strconv.Atoi(ips[2])
	ip4, _ := strconv.Atoi(ips[3])
	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return
	}
	ip += uint32(ip1 * 0x1000000)
	ip += uint32(ip2 * 0x10000)
	ip += uint32(ip3 * 0x100)
	ip += uint32(ip4)
	return
}

func Long2ip(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}
