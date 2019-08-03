 epaylinks/paycenter适用于易票联支付网关

## Complete

- 支付接口(Payment)
- 订单查询接口(Query)

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

	"github.com/fromiuan/epaylinks/paycenter"
)

var (
	pfxFilePath = "./yzt.pfx" // 证书位置
	password    = "xxxxxx"    // 证书密码
	signNo      = ""          // 证书序列号
	merchant    = "xxxxxx"    // 商户编号

	client *paycenter.Client
)

func init() {
	pfxData, err := ioutil.ReadFile(pfxFilePath)
	if err != nil {
		panic("not find pfx file")
	}
	client = paycenter.NewClient(pfxData, password, signNo, merchant)
	client.SetDev(true)   // 开启开发test模式
	client.SetDebug(true) // 开启debug模式
}

func main() {
	goi := &paycenter.GetOI{
		Version:         "4.0",
		SignType:        "SHA256withRSA",
		CertId:          pclicent.GetSignNo(),
		Sign:            "",
		Partner:         "130",
		OutTradeNo:      "123",
		TotalFee:        "10",
		CurrencyType:    "CNY",
		ReturnUrl:       "http://xxx.com/pay",
		NotifyUrl:       "http://xxx.com/notify",
		OrderCreateIp:   "127.0.0.1",
		PayId:           "",
		Base64Memo:      "",
		StoreOiType:     "",
		TimeOut:         "",
		MerchantAccount: "",
		CardType:        "",
	}

	rsps, err := pclicent.GetOI(goi)
	if err != nil {
		fmt.Println("===err", err)
	}
	fmt.Println("===rsps", rsps)
}
```