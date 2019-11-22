# Epaylinks
 epaylinks适用于易票联efps商户接入的sdk以及支付网关,efps暂未实现分账功能api
 易票联支付网关(网页支付)接口可见[paycenter](https://github.com/fromiuan/epaylinks/tree/master/paycenter)

## Complete

### espf

- 统一下单接口(Payment)
- 支付结果异步通知接口(Notify)
- 支付结果查询接口(Query)
- 微信APP支付接口(AppWxPayment)
- 微信H5支付接口(H5WxPayment)
- 主扫支付接口(NativePayment)
- 微信公众号/小程序支付接口(WxJSAPIPayment)
- 支付宝服务窗支付接口(AliJSAPIPayment)
- 扫码（被扫）支付接口(MicroPayment)
- 申请快捷支付接口(QuickpayApply)
- 提交快捷支付接口(QuickpayCommit)
- 网银支付接口(UnionPayMent)
- 协议支付绑卡预交易(ProtocolBindCard)
- 协议支付绑卡确认(ProtocolBindCardConfirm)
- 协议支付预交易(ProtocolPayPre)
- 协议支付确认交易(ProtocolPayConfirm)
- 协议支付解绑(ProtocolUnBindCard)
- 订单撤销接口(PayCancel)
- 关单接口(PayClose)

### 支付网关

- 支付接口(paycenter/Payment)
- 订单查询接口(paycenter/Query)

## 接口文档材料参考
	
- 易票联支付网关(网页支付)标准接口规范
- EFPS商户接入文档

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