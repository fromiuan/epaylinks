package epaylinks

import (
	"testing"
)

func TestQuery(t *testing.T) {
	test_query := &Query{
		CustomerCode:  "xxxxxx",
		OutTradeNo:    "xxxxxx",
		TransactionNo: "ZF201907309404540788007",
		NonceStr:      "123",
	}

	rsp, err := client.Query(test_query)
	if err != nil {
		t.Log(err)
	}
	t.Log(rsp)
}
