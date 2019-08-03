package epaylinks

import (
	"errors"
)

type ProtocolBindCardConfirm struct {
	SmsNo        string `json:"smsNo"`        // 绑卡流水号
	CustomerCode string `json:"customerCode"` // 商户号
	MemberId     string `json:"memberId"`     // 会员编号
	NonceStr     string `json:"nonceStr"`     // 随机字符串
	SmsCode      string `json:"smsCode"`      // 短息验证码
}

type ProtocolBindCardConfirmRsp struct {
	ReturnCode   string `json:"returnCode"`   // 返回状态码
	ReturnMsg    string `json:"returnMsg"`    // 返回信息
	SmsNo        string `json:"smsNo"`        // 绑卡流水号
	protocol     string `json:"protocol"`     // 协议号
	CustomerCode string `json:"customerCode"` // 商户号
	MemberId     string `json:"memberId"`     // 会员编号
	NonceStr     string `json:"nonceStr"`     // 随机字符串
}

// 协议支付绑卡确认
func (c *Client) ProtocolBindCardConfirm(bcc *ProtocolBindCardConfirm) (rsp *ProtocolBindCardConfirmRsp, err error) {
	rsp = new(ProtocolBindCardConfirmRsp)

	err = bcc.checkParms()
	if err != nil {
		return rsp, err
	}

	err = c.doPostReq(protocolBindCardConfirm, bcc, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (bcc *ProtocolBindCardConfirm) checkParms() error {
	if bcc.SmsNo == "" {
		return errors.New("绑卡流水号不能为空")
	}
	if bcc.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if bcc.MemberId == "" {
		return errors.New("会员编号不能为空")
	}
	if bcc.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	if bcc.SmsCode == "" {
		return errors.New("短息验证码不能为空")
	}
	return nil
}
