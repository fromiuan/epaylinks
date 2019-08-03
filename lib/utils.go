package lib

import (
	"bytes"
	"encoding/json"
	"sort"
	"strconv"
)

func ParmsSign(parms map[string]interface{}) string {
	keys := make([]string, 0, len(parms))
	for key, _ := range parms {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	var buf bytes.Buffer
	for _, k := range keys {
		if k == "sign" {
			continue
		}
		value := ToString(parms[k])
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(value)
		buf.WriteByte('&')
	}
	return buf.String()
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
