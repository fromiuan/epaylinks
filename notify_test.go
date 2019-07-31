package epaylinks

import (
	"testing"
)

func TestNotify(t *testing.T) {
	body := `{"amount":5,"procedureFee":1,"payTime":1564461847641,"settCycle":"","settCycleInterval":"","outTradeNo":"20190724113802346","transactionNo":"ZF201907309404540788007","customerCode":"5651300000000952","payState":"00","attachData":"attachData","channelOrder":"4200000345201907304406403464","nonceStr":"58cb4f64eccd47708add64908aaa7246"}`
	sign := "V9lsNYksELBnBtBzGI0PJ9dPntNM0zLGvWkEvroZl4kiotOEC1Jk6lqz/8pkhQCUZpbcVWrV0/ivapDYIRLEs3/AuPe9qlphArZNnCnEYeJ9IXbUAbOO480bxZooWgf0hOWjYL9ojCo6nD4ErOIxjeG0Z8JInXGPR2N6j5CLb2Kbqz3+Gb4o3YXJCOYMmJuB/IIRhb/EiK3db0hvfbGgrLfZxkvZ50U7uXuhQiC2Dyw8x+gQAV9pouVDcQU7LZCf5rhyUaNlDAWcCAmBh55XRe/rcNEZd0pLmEzAI7dqFdfQ0D9Ko0lcxPP9WCKLumaUybpEVg6PlorFETH/g8pGTg=="

	notify, err := client.Notify(sign, body)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(notify)
}
