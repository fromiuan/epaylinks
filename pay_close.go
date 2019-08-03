package epaylinks

import (
	"errors"
)

type PayClose struct {
	CustomerCode  string `json:"customerCode"`  // 商户号
	OutCancelNo   string `json:"outCancelNo"`   // 撤销单单号
	OutTradeNo    string `json:"outTradeNo"`    // 商户订单号
	TransactionNo string `json:"transactionNo"` // 易票联订单号
	NonceStr      string `json:"nonceStr"`      // 随机字符串
}

type PayCloseRsp struct {
	CustomerCode  string `json:"customerCode"`  // 商户号
	ReturnCode    string `json:"returnCode"`    // 返回码
	ReturnMsg     string `json:"returnMsg"`     // 返回信息
	OutTradeNo    string `json:"outTradeNo"`    // 商户订单号
	TransactionNo string `json:"transactionNo"` // 易票联订单号
	NonceStr      string `json:"nonceStr"`      // 随机字符串
}

// 关单接口
func (c *Client) PayClose(pc *PayClose) (rsp *PayCloseRsp, err error) {
	rsp = new(PayCloseRsp)

	err = pc.checkParms()
	if err != nil {
		return rsp, err
	}

	err = c.doPostReq(payClose, pc, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (pc *PayClose) checkParms() error {
	if pc.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if pc.OutCancelNo == "" {
		return errors.New("撤销单单号不能为空")
	}
	if pc.OutTradeNo == "" {
		return errors.New("商户订单号不能为空")
	}
	if pc.TransactionNo == "" {
		return errors.New("易票联订单号不能为空")
	}
	if pc.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	return nil
}
