package coolq

import (
	"strings"
)

// EnCodes array消息节点们转成CQ码
func EnCodes(data []MsgElem) string {
	output := ""
	for _, value := range data {
		output += EnCodeCQ(value)
	}
	return output
}

// EnCodeCQ 单个array节点转CQ码
func EnCodeCQ(Data MsgElem) string {
	output := "[CQ:" + Data.Type
	for key, value := range Data.Data {
		if value == "" {
			continue
		}
		output += "," + key + "=" + encodeValue(value)
	}
	return output + "]"
}

func encodeText(text string) string {
	text = strings.Replace(text, "&", "&amp;", -1)
	text = strings.Replace(text, "[", "&#91;", -1)
	text = strings.Replace(text, "]", "&#93;", -1)
	return text

}

func encodeValue(text string) string {
	text = encodeText(text)
	text = strings.Replace(text, ",", "&#44;", -1)
	return text
}

// EnXMLCode 生成xml的cq码
func EnXMLCode(message string, resid int) string {
	return EnCodeCQ(NewXML(message, resid)) + " "
}

// EnJSONCode 生成JSON的CQ码
func EnJSONCode(data string, redis int) string {
	return EnCodeCQ(NewJson(data, redis)) + " "
}

// EnImageCode 生成Image的CQ码
func EnImageCode(file string, cache int) string {
	return EnCodeCQ(NewImage(file, cache)) + " "
}

// EnFlashImageCode 闪照
func EnFlashImageCode(file string, cache int) string {
	return EnCodeCQ(NewFlashImage(file, cache)) + " "
}

// EnShowImageCode 秀图
// id	类型
// 40000	普通
// 40001	幻影
// 40002	抖动
// 40003	生日
// 40004	爱你
// 40005	征友
func EnShowImageCode(file string, id int, cache int) string {
	return EnCodeCQ(NewShowImage(file, id, cache)) + " "
}

// EnAtCode 生成At CQ码。 @all 的时候 qq填 `all`
func EnAtCode(qq string) string {
	return EnCodeCQ(NewAt(qq))
}

// EnFaceCode 生成表情CQ码
func EnFaceCode(id int) string {
	return EnCodeCQ(NewFace(id)) + " "
}

// EnEmojiCode 生成emoji 的CQ码
func EnEmojiCode(id int) string {
	return EnCodeCQ(NewEmoji(id)) + " "
}

// EnBfaceCode 生成Bface 的CQ码
func EnBfaceCode(id int) string {
	return EnCodeCQ(NewBface(id)) + " "
}

// EnSfaceCode 生成Sface 的CQ码
func EnSfaceCode(id int) string {
	return EnCodeCQ(NewSFace(id)) + " "
}

// EnRecordCode 生成路由CQ码 | file 语音文件名 | magic 发送时可选, 默认 0, 设置为 1 表示变声
func EnRecordCode(file string, magic int) string {
	return EnCodeCQ(NewRecord(file, magic)) + " "
}

// EnRpsCode 猜拳魔法表情
func EnRpsCode(rpstype string) string {
	return EnCodeCQ(NewRps(rpstype)) + " "
}

// EnDiceCode 掷骰子魔法表情
func EnDiceCode(dicType string) string {
	return EnCodeCQ(NewDice(dicType)) + " "
}

// EnShakeCode 窗口抖动（戳一戳） 发
func EnShakeCode() string {
	return EnCodeCQ(NewShake()) + " "
}

// EnAnonymousCode 匿名发消息 发
func EnAnonymousCode(ignore bool) string {
	return EnCodeCQ(NewAnonymous(ignore)) + " "
}

// EnMusicCode 生成音乐的CQ码
func EnMusicCode(mtype string, id string) string {
	return EnCodeCQ(NewMusic(mtype, id)) + " "
}

// EnDiyMusicCode 自定义音乐分享
func EnDiyMusicCode(url, audio, title, content, image string) string {
	return EnCodeCQ(NewDiyMusic(url, audio, title, content, image)) + " "
}

// EnShareCode 链接分享
func EnShareCode(url, title, content, image string) string {
	return EnCodeCQ(NewShare(url, title, content, image)) + " "
}

// EnVideoCode file 视频地址, 支持http和file发送 cover 视频封面, 支持http, file和base64发送, 格式必须为jpg
func EnVideoCode(file, cover string) string {
	return EnCodeCQ(NewVideo(file, cover)) + " "
}

// EnReplyCode 回复
func EnReplyCode(id int) string {
	return EnCodeCQ(NewReply(id)) + " "
}

// EnDiyRyplyCode 自定义回复
func EnDiyRyplyCode(text string, qq int64, time int64, seq int64) string {
	return EnCodeCQ(NewDiyReply(text, qq, time, seq)) + " "
}

// EnPokeCode 群 戳一戳
func EnPokeCode(qq int64) string {
	return EnCodeCQ(NewPoke(qq)) + " "
}

// EnGiftCode 礼物
// 目前支持的礼物 ID :
//
// id	类型
// 0	甜 Wink
// 1	快乐肥宅水
// 2	幸运手链
// 3	卡布奇诺
// 4	猫咪手表
// 5	绒绒手套
// 6	彩虹糖果
// 7	坚强
// 8	告白话筒
// 9	牵你的手
// 10	可爱猫咪
// 11	神秘面具
// 12	我超忙的
// 13	爱心口罩
func EnGiftCode(qq int64, id int) string {
	return EnCodeCQ(NewGift(qq, id)) + " "
}

// EnCardImageCode 装逼大图
func EnCardImageCode(file string, minwidth, minheight, maxwidth, maxheight int64, source string, icon string) string {
	return EnCodeCQ(NewCardImage(file, minwidth, minheight, maxwidth, maxheight, source, icon)) + " "
}

// EnTTSCode 文本转语音 发
func EnTTSCode(text string) string {
	return EnCodeCQ(NewTTS(text)) + " "
}
