package model

import (
	"bytes"
)

//PinWechatKey : pin we chat key
func PinWechatKey(openID, appID string) string {
	return pinTwoStr(openID, appID, '/')
}

//PinMobile : pin mobile
func PinMobile(mobile, area string) string {
	return pinTwoStr(mobile, area, '-')
}

//pin two words
func pinTwoStr(a, b string, key byte) string {
	var buf = bytes.NewBuffer(make([]byte, len(a)+len(b)+1))
	buf.WriteString(a)
	buf.WriteByte(key)
	buf.WriteString(b)
	return buf.String()
}
