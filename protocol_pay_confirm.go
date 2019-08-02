package epaylinks

import (
	"errors"
)

type ProtocolPayConfirm struct {
	Token        string `json:"token"`        // 绑卡流水号
	CustomerCode string `json:"customerCode"` // 商户号
	MemberId     string `json:"memberId"`     // 会员编号
	Protocol     string `json:"protocol"`     // 协议号
	NonceStr     string `json:"nonceStr"`     // 随机字符串
	SmsCode      string `json:"smsCode"`      // 短息验证码
}

type ProtocolPayConfirmRsp struct {
	ReturnCode        string `json:"returnCode"`        // 返回状态码
	ReturnMsg         string `json:"returnMsg"`         // 返回信息
	CustomerCode      string `json:"customerCode"`      // 商户号
	MemberId          string `json:"memberId"`          // 会员编号
	NonceStr          string `json:"nonceStr"`          // 随机字符串
	transactionNo     string `json:"transactionNo"`     // 易票联订单号
	channelOrder      string `json:"channelOrder"`      // 上游订单号
	Amount            int64  `json:"amount"`            // 支付金额
	ProcedureFee      int64  `json:"procedureFee"`      // 手续费
	PayState          string `json:"payState"`          // 支付结果
	PayTime           string `json:"payTime"`           // 支付完成时间
	SettCycle         string `json:"settCycle"`         // 该支付交易所属的清算周期
	settCycleInterval int    `json:"settCycleInterval"` // 清算周期长度
}

// 协议支付确认交易
func (c *Client) ProtocolPayConfirm(ppc *ProtocolPayConfirm) (rsp *ProtocolPayConfirmRsp, err error) {
	rsp = new(ProtocolPayConfirmRsp)

	err = ppc.checkParms()
	if err != nil {
		return rsp, err
	}

	req := newSetting(protocolPayConfirm, c)
	err = req.doPostReq(ppc, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}
func (ppc *ProtocolPayConfirm) checkParms() error {
	if ppc.Token == "" {
		return errors.New("绑卡流水号不能为空")
	}
	if ppc.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if ppc.MemberId == "" {
		return errors.New("会员编号不能为空")
	}
	if ppc.Protocol == "" {
		return errors.New("协议号不能为空")
	}
	if ppc.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	if ppc.SmsCode == "" {
		return errors.New("短息验证码不能为空")
	}
	return nil
}
