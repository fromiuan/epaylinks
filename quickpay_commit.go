package epaylinks

import (
	"errors"
)

type QuickpayCommit struct {
	Token    string `json:"token"`    // 唯一代表该订单
	SmsCode  string `json:"smsCode"`  // 持卡人收到的短信验证码
	NonceStr string `json:"nonceStr"` // 随机字符串
}

type QuickpayCommitRsp struct {
	ReturnCode        string `json:"returnCode"`        // 返回状态码
	ReturnMsg         string `json:"returnMsg"`         // 返回信息
	OutTradeNo        string `json:"outTradeNo"`        // 商户订单号
	TransactionNo     string `json:"transactionNo"`     // 易票联订单号
	ChannelOrder      string `json:"channelOrder"`      // 上游订单号
	Amount            int64  `json:"amount"`            // 支付金额
	ProcedureFee      int64  `json:"procedureFee"`      // 手续费
	PayState          string `json:"payState"`          // 支付结果
	PayTime           string `json:"payTime"`           // 支付完成时间
	SettCycle         string `json:"settCycle"`         // 该支付交易所属的清算周期
	SettCycleInterval int    `json:"settCycleInterval"` // 清算周期长度
	NonceStr          string `json:"nonceStr"`          // 随机字符串
}

// 提交快捷支付接口
func (c *Client) QuickpayCommit(qc *QuickpayCommit) (rsp *QuickpayCommitRsp, err error) {
	rsp = new(QuickpayCommitRsp)

	err = qc.checkParms()
	if err != nil {
		return rsp, err
	}

	err = c.doPostReq(quickpayCommit, qc, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (qc *QuickpayCommit) checkParms() error {
	if qc.Token == "" {
		return errors.New("Token不能为空")
	}
	if qc.SmsCode == "" {
		return errors.New("请填写短信验证码")
	}
	if qc.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	return nil
}
