package epaylinks

import (
	"testing"
	"time"
)

func TestPayment(t *testing.T) {
	test_unifieder := &Unifieder{
		OutTradeNo:    "2019072411380",
		CustomerCode:  "xxxxxx",
		TerminalNo:    "10001",
		ClientIp:      "127.0.0.1",
		NeedSplit:     false,
		NoCreditCards: true,
		NonceStr:      "bed33215a4d741bc8322acf17ee5dfaf",
		OrderInfo: &OrderInfo{
			Id:           "test",
			BusinessType: "test",
			GoodsList: []GoodsList{GoodsList{
				Name:   "红富士",
				Number: "1箱",
				Amount: 1,
			}, GoodsList{
				Name:   "82年的茅台",
				Number: "1瓶",
				Amount: 1,
			}},
		},
		PayAmount:            5,
		PayCurrency:          "CNY",
		ChannelType:          "01",
		NotifyUrl:            "http://xxxxxx.com/notify",
		RedirectUrl:          "http://xxxxxx.com/pay",
		AttachData:           "attachData",
		TransactionStartTime: time.Now().Format("20060102150405"),
		TransactionEndTime:   "",
	}
	rsp, err := client.Payment(test_unifieder)
	if err != nil {
		t.Log(err)
	}
	t.Log(rsp)
}
