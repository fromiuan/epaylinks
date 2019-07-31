# Epaylinks
 epaylinks适用于易票联efps商户接入的sdk

## Install

```bash
go get github.com/fromiuan/epaylinks
```

## Usage

```go
package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/fromiuan/epaylinks"
)

var (
	pfxFilePath = "./efps.pfx" // 证书位置
	password    = "xxxxxx"     // 证书密码
	signNo      = "xxxxxx"     // 证书序列号
	merchant    = "xxxxxx"     // 商户编号

	client *epaylinks.Client
)

func init() {
	pfxData, err := ioutil.ReadFile(pfxFilePath)
	if err != nil {
		panic("not find pfx file")
	}
	client = epaylinks.NewClient(pfxData, password, signNo, merchant)
	client.SetDev(true)   // 开启开发test模式
	client.SetDebug(true) // 开启debug模式
}

func main() {
	unifieder := &epaylinks.Unifieder{
		OutTradeNo:    "20190724113802347",
		CustomerCode:  "xxxxxx",
		TerminalNo:    "10001",
		ClientIp:      "127.0.0.1",
		NeedSplit:     false,
		NoCreditCards: true,
		NonceStr:      "bed33215a4d741bc8322acf17ee5dfaf",
		OrderInfo: &epaylinks.OrderInfo{
			Id:           "test",
			BusinessType: "test",
			GoodsList: []epaylinks.GoodsList{epaylinks.GoodsList{
				Name:   "红富士",
				Number: "1箱",
				Amount: 1,
			}, epaylinks.GoodsList{
				Name:   "82年的茅台",
				Number: "1瓶",
				Amount: 1,
			}},
		},
		PayAmount:            5,
		PayCurrency:          "CNY",
		ChannelType:          "01",
		NotifyUrl:            "http://xxx.com/notify",
		RedirectUrl:          "http://xxx.com/pay",
		AttachData:           "attachData",
		TransactionStartTime: time.Now().Format("20060102150405"),
		TransactionEndTime:   "",
	}

	rsp, err := client.Payment(unifieder)
	if err != nil {
		fmt.Println("payment error", err)
	}
	fmt.Println(rsp)
}
```