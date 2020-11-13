package main

import (
	"fmt"

	"github.com/k-sone/snmpgo"
)

func main() {
	snmp, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version:   snmpgo.V2c,
		Address:   "192.168.0.192:161",
		Retries:   1,
		Community: "public",
	})
	if err != nil {
		fmt.Println("Connect Server Failed. ", err)
		return
	}

	oids, err := snmpgo.NewOids([]string{
		"1.3.6.1.2.1.1.1.0", // 单独OID节点，下面不含子节点
		"1.3.6.1.2.1.1.1.3",
		// "1.3.6.1.2.1.1", // 整个取出，含子节点，只能用于bulk
	})
	if err != nil {
		// Failed to parse Oids
		fmt.Println("OID Parsed Failed. ", err)
		return
	}

	if err = snmp.Open(); err != nil {
		fmt.Println("Open Failed. ", err)
		return
	}
	defer snmp.Close() // 关闭连接

	pdu, err := snmp.GetRequest(oids)
	if err != nil {
		// Failed to request
		fmt.Println(err)
		return
	}
	if pdu.ErrorStatus() != snmpgo.NoError {
		// Received an error from the agent
		fmt.Println(pdu.ErrorStatus(), pdu.ErrorIndex())
	}

	// 返回[]*Varbind切片
	// fmt.Printf("%T, %v\n", pdu.VarBinds(), pdu.VarBinds())
	// VarBinds 嵌套的Variable是一个interface的结构体，而且是非匿名嵌套
	for _, varBindPtr := range pdu.VarBinds() {
		fmt.Println(varBindPtr.Oid, varBindPtr.Variable.Type(), varBindPtr.Variable.String())
	}

	// 如果调用matchoid会只返回非空
	// fmt.Println(pdu.VarBinds().MatchOid(oids[0]))

	// // bulk walk 方法
	// pdu, err := snmp.GetBulkWalk(oids, 0, 1)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if pdu.ErrorStatus() != snmpgo.NoError {
	// 	fmt.Println(pdu.ErrorStatus(), pdu.ErrorIndex())
	// 	return
	// }

	// for _, val := range pdu.VarBinds() {
	// 	fmt.Printf("%s = %s: %s\n", val.Oid, val.Variable.Type(), val.Variable)
	// }

}
