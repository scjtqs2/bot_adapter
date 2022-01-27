package sha256

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// CalculateSignature 计算签名
func CalculateSignature(timestamp, nonce, encryptKey, bodystring string) string {
	var b strings.Builder
	b.WriteString(timestamp)
	b.WriteString(nonce)
	b.WriteString(encryptKey)
	b.WriteString(bodystring) // bodystring指整个请求体，不要在反序列化后再计算
	bs := []byte(b.String())
	h := sha256.New()
	h.Write(bs)
	bs = h.Sum(nil)
	sig := fmt.Sprintf("%x", bs)
	return sig
}
