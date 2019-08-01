package epaylinks

import (
	"errors"
)

type ProtocolBindCard struct {
	MchtOrderNo      string `json:"mchtOrderNo"`      // 请求流水号
	CustomerCode     string `json:"customerCode"`     // 商户号 String(32) 是
	MemberId         string `json:"memberId"`         // 会员编号 String(32) 是
	UserName         string `json:"userName"`         // 会员姓名 String(50) 是 ”
	PhoneNum         string `json:"phoneNum"`         // 手机号码 String(11) 是
	BankCardNo       string `json:"bankCardNo"`       // 银行卡号 String(20) 是
	BankCardType     string `json:"bankCardType"`     // 银行卡类型 String(16) 是 debit：借记卡 credit：贷记卡
	CertificatesType string `json:"certificatesType"` // 证件类型 String(128) 是 固定传 01-身份证
	CertificatesNo   string `json:"certificatesNo"`   // 证件号码 String(32) 是
	Expired          string `json:"expired"`          // 有效日期 String(4) 否 信用卡必填 yymm，例如 2018 年 2 月，则为 1802
	Cvn              string `json:"cvn"`              //  CVN 码 String(10) 否 信用卡必填
	NonceStr         string `json:"nonceStr"`         // 随机字符串 String(32) 是
}

type ProtocolBindCardRsp struct {
	ReturnCode   string `json:"returnCode"`   // 返回状态码
	ReturnMsg    string `json:"returnMsg"`    // 返回信息 String(32) 否 处理失败的原因
	SmsNo        string `json:"smsNo"`        // 绑卡流水号 String 是 用于绑卡确认时传递
	CustomerCode string `json:"customerCode"` // 商户号 String(32) 是
	MemberId     string `json:"memberId"`     // 会员编号 String(32) 是
	NonceStr     string `json:"nonceStr"`     // 随机字符串 String(32) 是
}

func (c *Client) ProtocolBindCard(pbc *ProtocolBindCard) (rsp *ProtocolBindCardRsp, err error) {
	rsp = new(ProtocolBindCardRsp)

	err = pbc.checkParms()
	if err != nil {
		return rsp, err
	}

	req := newSetting(protocolBindCard, c)
	err = req.doPostReq(pbc, rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, err
}

func (pbc *ProtocolBindCard) checkParms() error {
	if pbc.CustomerCode == "" {
		return errors.New("商户号不能为空")
	}
	if pbc.MemberId == "" {
		return errors.New("会员编号不能为空")
	}
	if pbc.UserName == "" {
		return errors.New("会员姓名不能为空")
	}
	if pbc.PhoneNum == "" {
		return errors.New("手机号码不能为空")
	}
	if pbc.BankCardNo == "" {
		return errors.New("银行卡号不能为空")
	}
	if pbc.BankCardType == "" {
		return errors.New("银行卡类型不能为空")
	}
	if pbc.CertificatesType == "" {
		return errors.New("证件类型不能为空")
	}
	if pbc.CertificatesNo == "" {
		return errors.New("证件号码不能为空")
	}
	if pbc.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	return nil
}
