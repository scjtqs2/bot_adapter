package coolq

import (
	"strings"
)

// DeCode 对收到的数据进行解码 将string的cq码转码成 array格式
func DeCode(raw string) (r []MsgElem) {
	var t, key string
	d := map[string]string{}
	makeElem := func() {
		msg := MsgElem{
			Type: t,
			Data: MSG{},
		}
		for k, v := range d {
			msg.Data[k] = v
		}
		r = append(r, msg)
	}
	for raw != "" {
		i := 0
		for i < len(raw) && !(raw[i] == '[' && i+4 < len(raw) && raw[i:i+4] == "[CQ:") {
			i++
		}
		// 非CQ码，转成正常 text
		if i > 0 {
			for _, txt := range SplitURL(decodeText(raw[:i])) {
				r = append(r, NewText(txt))
			}
		}

		if i+4 > len(raw) {
			return
		}
		raw = raw[i+4:] // skip "[CQ:"
		i = 0
		for i < len(raw) && raw[i] != ',' && raw[i] != ']' {
			i++
		}
		if i+1 > len(raw) {
			return
		}
		t = raw[:i]
		for k := range d { // clear the map, reuse it
			delete(d, k)
		}
		raw = raw[i:]
		i = 0
		for {
			if raw[0] == ']' {
				makeElem()
				raw = raw[1:]
				break
			}
			raw = raw[1:]

			for i < len(raw) && raw[i] != '=' {
				i++
			}
			if i+1 > len(raw) {
				return
			}
			key = raw[:i]
			raw = raw[i+1:] // skip "="
			i = 0
			for i < len(raw) && raw[i] != ',' && raw[i] != ']' {
				i++
			}

			if i+1 > len(raw) {
				return
			}
			d[key] = decodeValue(raw[:i])
			raw = raw[i:]
			i = 0
		}
	}
	return
}

// decodeText 解码文字
func decodeText(text string) string {
	text = strings.ReplaceAll(text, "&amp;", "&")
	text = strings.ReplaceAll(text, "&#91;", "[")
	text = strings.ReplaceAll(text, "&#93;", "]")
	return text
}

// decodeValue 解码CQ里的值
func decodeValue(text string) string {
	text = decodeText(text)
	text = strings.ReplaceAll(text, "&#44;", ",")
	return text
}
