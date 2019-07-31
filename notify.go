package epaylinks

import (
	"encoding/json"
)

type Notify struct {
	CustomerCode     string `json:"customerCode"`
	OutTradeNo       string `json:"outTradeNo"`
	TransactionNo    string `json:"transactionNo"`
	ChannelOrder     string `json:"channelOrder"`
	Amount           int    `json:"amount"`
	PayState         string `json:"payState"`
	PayTime          int64  `json:"payTime"`
	SettCycle        string `json:"settCycle"`
	SettCycleInterva string `json:"settCycleInterval"`
	ProcedureFee     int    `json:"procedureFee"`
	AttachData       string `json:"attachData"`
	NonceStr         string `json:"nonceStr"`
}

type NotifyRsp struct {
	ReturnCode string `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
}

// 通知验证
func (c *Client) Notify(sign, body string) (rsp *Notify, err error) {
	rsp = new(Notify)
	err = json.Unmarshal([]byte(body), rsp)
	if err != nil {
		return rsp, err
	}
	err = c.Verify(sign, body)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}
