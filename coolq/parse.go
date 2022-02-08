package coolq

import (
	"math"
	"regexp"
	"strconv"
	"sync"
)

// TYPE STRING
const (
	TEXT      = "text"
	FACE      = "face"
	EMOJI     = "emoji"
	BFACE     = "bface"
	SFACE     = "sface"
	IMAGE     = "image"
	RECORD    = "record"
	AT        = "at"
	RPS       = "rps"
	DICE      = "dice"
	SHAKE     = "shake"
	ANONYMOUS = "anonymous"
	MUSIC     = "music"
	SHARE     = "share"
	XML       = "xml"
	JSON      = "json"
	VIDEO     = "video"
	REDBAG    = "redbag" // 只有收
	REPLY     = "reply"
	POKE      = "poke" // 群聊 戳一戳
	GIFT      = "gift"
	CARDIMAGE = "cardimage" // 装逼大图
	TTS       = "tts"       // tts 语音
)

var (
	// once lazy compile the reg
	once sync.Once
	// reg is splitURL regex pattern.
	reg *regexp.Regexp
)

// SplitURL 将给定URL字符串分割为两部分，用于URL预处理防止风控
func SplitURL(s string) []string {
	once.Do(func() { // lazy init.
		reg = regexp.MustCompile(`(?i)[a-z\d][-a-z\d]{0,62}(\.[a-z\d][-a-z\d]{0,62})+\.?`)
	})
	idx := reg.FindAllStringIndex(s, -1)
	if len(idx) == 0 {
		return []string{s}
	}
	var result []string
	last := 0
	for i := 0; i < len(idx); i++ {
		if len(idx[i]) != 2 {
			continue
		}
		m := int(math.Abs(float64(idx[i][0]-idx[i][1]))/1.5) + idx[i][0]
		result = append(result, s[last:m])
		last = m
	}
	result = append(result, s[last:])
	return result
}

// ParseAtCode 解析 CQ码的AT 中的QQ号
func ParseAtCode(message string) []string {
	patten := `\[CQ:at,qq=(\d+)\]`
	r, _ := regexp.Compile(patten)
	p := r.FindAllStringSubmatch(message, -1)
	var ret []string
	for _, strings := range p {
		ret = append(ret, strings[1])
	}
	return ret
}

// IsAtMe 判断是否at某人
func IsAtMe(message string, qq int64) (bool, error) {
	patten := `\[CQ:at,qq=(\d+)\]`
	r, err := regexp.Compile(patten)
	if err != nil {
		return false, err
	}
	p := r.FindAllStringSubmatch(message, -1)
	qqstr := strconv.FormatInt(qq, 10)
	for _, strings := range p {
		if strings[1] == qqstr {
			return true, nil
		}
	}
	return false, nil
}

// CleanCQCode 用正则清理掉所有的 CQ码
func CleanCQCode(message string) string {
	reg, _ := regexp.Compile(`\[CQ:\w+?.*?]`)
	return reg.ReplaceAllString(message, "")
}
