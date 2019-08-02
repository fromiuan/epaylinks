package epaylinks

import (
	"errors"
)

type Query struct {
	CustomerCode  string `json:"customerCode"`
	OutTradeNo    string `json:"outTradeNo"`
	TransactionNo string `json:"transactionNo"`
	NonceStr      string `json:"nonceStr"`
}

type QueryRsp struct {
	ReturnCode        string `json:"returnCode"`
	ReturnMsg         string `json:"returnMsg"`
	CustomerCode      string `json:"customerCode"`
	OutTradeNo        string `json:"outTradeNo"`
	TransactionNo     string `json:"transactionNo"`
	ChannelOrder      string `json:"channelOrder"`
	Amount            int    `json:"amount"`
	PayState          string `json:"payState"`
	PayTime           string `json:"payTime"`
	SettCycle         string `json:"settCycle"`
	SettCycleInterval int    `json:"settCycleInterval"`
	ProcedureFee      int    `json:"procedureFee"`
	AttachData        string `json:"attachData"`
	NonceStr          string `json:"nonceStr"`
}

// 支付结果查询
func (c *Client) Query(q *Query) (rsp *QueryRsp, err error) {
	rsp = new(QueryRsp)

	err = q.checkParms()
	if err != nil {
		return rsp, err
	}

	req := newSetting(paymentQuery, c)
	err = req.doPostReq(q, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (q *Query) checkParms() error {
	if q.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	return nil
}
