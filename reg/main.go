package main

import (
	"fmt"
	"regexp"
)

func main() {
	// str := "qewrqwer123123"
	// re := regexp.MustCompile(`\d{3}`) // 使用mustcompile比cmopile安全
	// if re == nil {
	// 	fmt.Println("re compile err")
	// 	return
	// }
	// result := re.FindAllString(str, -1) // 如果第二个参数n<0,返回所有结果，n=1返回前1个结果
	// fmt.Println(result)
	// result2 := re.FindAllStringSubmatch(str, -1)
	// fmt.Println(result2)
	// result3 := re.FindAllStringIndex(str, -1) // 返回的是每个匹配到的子串的范围index
	// fmt.Println(result3)
	// fmt.Printf("%c, %T\n", []byte(str), []byte(str))

	// matched, _ := regexp.MatchString(`(\d+/){2}\d+`, "10/1/ll")
	// if matched {
	// 	fmt.Println(matched)
	// }

	// 子串
	re := regexp.MustCompile(`show router bgp neighbor \| match "Peer\|Description\|State" expression([\s\S]*?)GDFOS-MS-IPMAN-BNG01-BJ-AL`)
	
	s := `*A:GDFOS-MS-IPMAN-BNG01-BJ-AL#                               show router bgp neighbor | match "Peer|Description|State" expression 
Peer                 : 183.233.6.1
Description          : GDFOS-MC-IPMAN-RT01-DEJL-HW
Peer AS              : 65277            Peer Port            : 57176
Peer Address         : 183.233.6.1
Peer Type            : Internal         Dynamic Peer         : No
State                : Established      Last State           : Established
Advertise Inactive   : Disabled         Peer Tracking        : Disabled
Damp Peer Oscillatio*: Disabled         Update Errors        : 0    
Peer                 : 183.233.6.2
Description          : GDFOS-MC-IPMAN-RT01-DS-HW
Peer AS              : 65277            Peer Port            : 51120
Peer Address         : 183.233.6.2
Peer Type            : Internal         Dynamic Peer         : No
State                : Established      Last State           : Established
Advertise Inactive   : Disabled         Peer Tracking        : Disabled
Damp Peer Oscillatio*: Disabled         Update Errors        : 0    
Peer                 : 2409:8054:20::1
Description          : GDFOS-MC-IPMAN-RT01-DEJL-HW
Peer AS              : 65277            Peer Port            : 50732
Peer Address         : 2409:8054:20::1
Peer Type            : Internal         Dynamic Peer         : No
State                : Established      Last State           : Established
Advertise Inactive   : Disabled         Peer Tracking        : Disabled
Damp Peer Oscillatio*: Disabled         Update Errors        : 0    
Peer                 : 2409:8054:20::2
Description          : GDFOS-MC-IPMAN-RT01-DS-HW
Peer AS              : 65277            Peer Port            : 64879
Peer Address         : 2409:8054:20::2
Peer Type            : Internal         Dynamic Peer         : No
State                : Established      Last State           : Established
Advertise Inactive   : Disabled         Peer Tracking        : Disabled
Damp Peer Oscillatio*: Disabled         Update Errors        : 0    
*A:GDFOS-MS-IPMAN-BNG01-BJ-AL#`
	ret := re.FindAllStringSubmatch(s, -1) // 返回一个切片[整个表达式的匹配情况, 子串1，子串2...]
	//ret := re.FindStringSubmatch(s)
	for _, r := range ret {
		fmt.Printf("%q\n", r)
		fmt.Println("===============")
	}
	// s = `PING 192.168.1.120 (192.168.1.120) 56(84) bytes of data.
	// From 192.168.1.19 icmp_seq=1 Destination Host Unreachable
	// From 192.168.1.19 icmp_seq=2 Destination Host Unreachable
	// From 192.168.1.19 icmp_seq=3 Destination Host Unreachable
	// From 192.168.1.19 icmp_seq=4 Destination Host Unreachable
	// From 192.168.1.19 icmp_seq=5 Destination Host Unreachable

	// --- 192.168.1.120 ping statistics ---
	// 5 packets transmitted, 0 received, +5 errors, 100% packet loss, time 4013ms
	// pipe 4`

	// ret2 := re.FindAllStringSubmatch(s, -1)
	// fmt.Printf("%q\n", ret2)
}
