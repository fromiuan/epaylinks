package epaylinks

import (
	"encoding/json"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/fromiuan/epaylinks/lib"
)

var (
	// 测试地址
	testHost = "http://test-efps.epaylinks.cn"
	// 正式地址
	proHost = "https://efps.epaylinks.cn"

	// 统一下单
	unified = "/api/txs/pay/UnifiedPayment"
	// 支付结果查询接口
	paymentQuery = "/api/txs/pay/PaymentQuery"
	// 微信APP支付接口
	appWxPayMent = "/api/txs/pay/appWxPayMent"
	// 微信H5支付接口
	wxH5Payment = "/api/txs/pay/WxH5Payment"
	// 主扫支付接口
	nativePayment = "/api/txs/pay/NativePayment"
	// 微信公众号/小程序支付接口
	wxJSAPIPayment = "/api/txs/pay/WxJSAPIPayment"
	// 支付宝服务窗支付接口
	aliJSAPIPayment = "/api/txs/pay/AliJSAPIPayment"
	// 扫码（被扫）支付接口
	microPayment = "/api/txs/pay/MicroPayment"
	// 申请快捷支付接口
	quickpayApply = "/api/txs/quickpay/apply"
	// 提交快捷支付接口
	quickpayCommit = "/api/txs/quickpay/commit"
	// 网银支付接口
	unionPayMent = "/api/txs/pay/unionPayMent"
	// 协议支付绑卡预交易
	protocolBindCard = "/api/txs/protocol/bindCard"
	// 协议支付绑卡确认
	protocolBindCardConfirm = "/api/txs/protocol/bindCardConfirm"
	// 协议支付预交易
	protocolPayPre = "/api/txs/protocol/protocolPayPre"
	// 协议支付确认交易
	protocolPayConfirm = "/api/txs/protocol/protocolPayConfirm"
)

type OrderInfo struct {
	Id           string      `json:"id"`
	BusinessType string      `json:"businessType"`
	GoodsList    []GoodsList `json:"goodsList"`
}

type GoodsList struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	Amount int64  `json:"amount"`
}

type setting struct {
	reqUrl   string
	client   *Client
	encoding *lib.Encoding
}

func newSetting(reqUrl string, c *Client) *setting {
	if c.dev {
		reqUrl = testHost + reqUrl
	} else {
		reqUrl = proHost + reqUrl
	}
	return &setting{
		reqUrl:   reqUrl,
		client:   c,
		encoding: lib.NewEncding(c.pfxData, c.password),
	}
}

func (rs *setting) doPostReq(params interface{}, rsp interface{}) error {
	req := lib.Post(rs.reqUrl)
	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return err
	}
	signStr, err := rs.encoding.Sign(paramsBytes)
	if err != nil {
		return err
	}
	req.Header("x-efps-sign-no", rs.client.signNo)
	req.Header("x-efps-sign-type", "SHA256withRSA")
	req.Header("x-efps-sign", signStr)
	req.Header("x-efps-timestamp", time.Now().Format("20060102150405"))
	req.Header("Content-Type", "application/json")
	req.Body(paramsBytes)
	err = req.ToJSON(rsp)
	if rs.client.debug {
		str, err := req.String()
		log.Println("req str info:", str, err)
	}
	if err != nil {
		return err
	}
	return err
}

// sort map
func (rs *setting) sort(params map[string]interface{}) {
	keys := make([]string, 0, len(params))
	if len(params) > 0 {
		for k, _ := range params {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var newMp = make(map[string]interface{})
		for _, k := range keys {
			newMp[k] = params[k]
		}
		params = newMp
	}
}

func ToMap(data interface{}) (map[string]interface{}, error) {
	var mp map[string]interface{}
	byts, err := json.Marshal(data)
	if err != nil {
		return mp, err
	}
	err = json.Unmarshal(byts, &mp)
	return mp, err
}

func ToString(arg interface{}) (result string) {
	switch val := arg.(type) {
	case int:
		result = strconv.Itoa(val)
	case string:
		result = val
	}
	return result
}
