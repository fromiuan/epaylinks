package epaylinks

import (
	"errors"
)

type PayCancel struct {
	CustomerCode  string `json:"customerCode"`  // 商户号
	OutCancelNo   string `json:"outCancelNo"`   // 撤销单单号
	OutTradeNo    string `json:"outTradeNo"`    // 商户订单号
	TransactionNo string `json:"transactionNo"` // 易票联订单号
	NonceStr      string `json:"nonceStr"`      // 随机字符串
}

type PayCancelRsp struct {
	ReturnCode    string `json:"returnCode"`    // 返回状态码
	ReturnMsg     string `json:"returnMsg"`     // 返回信息
	OutCancelNo   string `json:"outCancelNo"`   // 商户退款单号
	OutTradeNo    string `json:"outTradeNo"`    // 商户订单号
	TransactionNo string `json:"transactionNo"` // 易票联订单号
	NonceStr      string `json:"nonceStr"`      // 随机字符串
}

// 订单撤销接口
func (c *Client) PayCancel(pubc *PayCancel) (rsp *PayCancelRsp, err error) {
	rsp = new(PayCancelRsp)

	err = pubc.checkParms()
	if err != nil {
		return rsp, err
	}

	err = c.doPostReq(payCancel, pubc, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (pubc *PayCancel) checkParms() error {
	if pubc.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if pubc.OutCancelNo == "" {
		return errors.New("撤销单单号不能为空")
	}
	if pubc.OutTradeNo == "" {
		return errors.New("商户订单号不能为空")
	}
	if pubc.TransactionNo == "" {
		return errors.New("易票联订单号不能为空")
	}
	if pubc.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	return nil
}
