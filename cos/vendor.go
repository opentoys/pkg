package cos

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

func jsonCopy(dst, src interface{}) (e error) {
	buf, e := json.Marshal(dst)
	if e != nil {
		return
	}
	e = json.Unmarshal(buf, &src)
	return
}

func newxml(buf []byte) (data map[string]interface{}, e error) {
	data = make(map[string]interface{})
	e = xml.Unmarshal(buf, &data)
	return
}

func queryValues(v interface{}) (qs url.Values, e error) {
	return values(v)
}

func httpHeader(v interface{}) (qs http.Header, e error) {
	return values(v)
}

func values(v interface{}) (qs map[string][]string, e error) {
	qs = make(url.Values)
	if v == nil {
		return
	}

	if data, ok := v.(map[string]string); ok {
		for k, v := range data {
			qs[k] = []string{v}
		}
		return
	}

	if data, ok := v.(url.Values); ok {
		qs = data
		return
	}

	if data, ok := v.(map[string][]string); ok {
		qs = url.Values(data)
		return
	}

	var data = map[string]interface{}{}
	buf, e := json.Marshal(v)
	if e != nil {
		return
	}
	e = json.Unmarshal(buf, &data)
	if e != nil {
		return
	}

	for k, v := range data {
		qs[k] = []string{fmt.Sprintf("%v", v)}
	}
	return
}
