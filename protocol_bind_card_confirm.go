package epaylinks

import (
	"errors"
)

type BindCardConfirm struct {
	SmsNo        string `json:"smsNo"`        // 绑卡流水号
	CustomerCode string `json:"customerCode"` // 商户号
	MemberId     string `json:"memberId"`     // 会员编号
	NonceStr     string `json:"nonceStr"`     // 随机字符串
	SmsCode      string `json:"smsCode"`      // 短息验证码
}

type BindCardConfirmRsp struct {
	ReturnCode   string `json:"returnCode"`   // 返回状态码
	ReturnMsg    string `json:"returnMsg"`    // 返回信息
	SmsNo        string `json:"smsNo"`        // 绑卡流水号
	protocol     string `json:"protocol"`     // 协议号
	CustomerCode string `json:"customerCode"` // 商户号
	MemberId     string `json:"memberId"`     // 会员编号
	NonceStr     string `json:"nonceStr"`     // 随机字符串
}

func (c *Client) BindCardConfirm(bcc *BindCardConfirm) (rsp *BindCardConfirmRsp, err error) {
	rsp = new(BindCardConfirmRsp)

	err = bcc.checkParms()
	if err != nil {
		return rsp, err
	}

	req := newSetting(protocolBindCardConfirm, c)
	err = req.doPostReq(bcc, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (bcc *BindCardConfirm) checkParms() error {
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
