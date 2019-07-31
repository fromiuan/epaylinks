package epaylinks

import (
	"io/ioutil"
	"testing"
)

const (
	pfxFile  = "./certs/efps.pfx"
	password = "xxxxxx"
	serial   = "xxxxxx"
	merchant = "xxxxxx" // 商户编号
)

var client *Client
var body = `{"amount":5,"procedureFee":1,"payTime":1564461847641,"settCycle":"","settCycleInterval":"","outTradeNo":"xxxx","transactionNo":"xxxx","customerCode":"xxxx","payState":"00","attachData":"attachData","channelOrder":"4200000345201907304406403464","nonceStr":"58cb4f64eccd47708add64908aaa7246"}`
var sign = "V9lsNYksELBnBtBzGI0PJ9dPntNM0zLGvWkEvroZl4kiotOEC1Jk6lqz/8pkhQCUZpbcVWrV0/ivapDYIRLEs3/AuPe9qlphArZNnCnEYeJ9IXbUAbOO480bxZooWgf0hOWjYL9ojCo6nD4ErOIxjeG0Z8JInXGPR2N6j5CLb2Kbqz3+Gb4o3YXJCOYMmJuB/IIRhb/EiK3db0hvfbGgrLfZxkvZ50U7uXuhQiC2Dyw8x+gQAV9pouVDcQU7LZCf5rhyUaNlDAWcCAmBh55XRe/rcNEZd0pLmEzAI7dqFdfQ0D9Ko0lcxPP9WCKLumaUybpEVg6PlorFETH/g8pGTg=="

func init() {
	pfxData, err := ioutil.ReadFile(pfxFile)
	if err != nil {
		panic(err)
	}
	client = NewClient(pfxData, password, serial, merchant)
	client.SetDev(true)   // 开启测试模式
	client.SetDebug(true) // 开启调试模式
}

func TestVerify(t *testing.T) {
	err := client.Verify(sign, body)
	t.Log(err)
}
