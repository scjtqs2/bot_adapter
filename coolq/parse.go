package coolq

import (
	"math"
	"regexp"
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
	TTS       = "tts" // tts 语音
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


