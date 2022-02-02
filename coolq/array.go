// Package coolq cq码的转码和解码
package coolq

import "fmt"

type MSG map[string]string

// MsgElem 结构体，用来转json用
type MsgElem struct {
	Type string            `json:"type"` // text image at 等
	Data map[string]string `json:"data"` // 不同类型的data不一样
}

// NewText 文本消息
func NewText(text string) MsgElem {
	return MsgElem{
		Type: TEXT,
		Data: MSG{
			"text": text,
		},
	}
}

// NewXML XML消息
func NewXML(message string, resid int) MsgElem {
	return MsgElem{Type: XML, Data: MSG{"data": message, "resid": fmt.Sprintf("%d", resid)}}
}

// NewJson JSON消息
func NewJson(data string, redis int) MsgElem {
	return MsgElem{Type: JSON,
		Data: MSG{
			"data":  data,
			"resid": fmt.Sprintf("%d", redis),
		}}
}

// NewImage 生成Image的CQ码
func NewImage(file string, cache int) MsgElem {
	return MsgElem{Type: IMAGE, Data: MSG{"file": file, "cache": fmt.Sprintf("%d", cache)}}
}

// NewFlashImage 闪照
func NewFlashImage(file string, cache int) MsgElem {
	return MsgElem{Type: IMAGE, Data: MSG{
		"file":  file,
		"type":  "flash",
		"cache": fmt.Sprintf("%d", cache),
	}}
}

// NewShowImage 秀图
// id	类型
// 40000	普通
// 40001	幻影
// 40002	抖动
// 40003	生日
// 40004	爱你
// 40005	征友
func NewShowImage(file string, id int, cache int) MsgElem {
	return MsgElem{Type: IMAGE, Data: MSG{
		"file":  file,
		"type":  "show",
		"cache": fmt.Sprintf("%d", cache),
		"id":    fmt.Sprintf("%d", id),
	}}
}

// NewAt 生成At CQ码。 @all 的时候 qq填 `all`
func NewAt(qq string) MsgElem {
	return MsgElem{Type: AT, Data: MSG{"qq": qq}}
}

// NewFace 生成表情CQ码
func NewFace(id int) MsgElem {
	return MsgElem{Type: FACE, Data: MSG{"id": fmt.Sprintf("%d", id)}}
}

// NewEmoji 生成emoji 的CQ码
func NewEmoji(id int) MsgElem {
	return MsgElem{Type: EMOJI, Data: MSG{"id": fmt.Sprintf("%d", id)}}
}

// NewBface 生成Bface 的CQ码
func NewBface(id int) MsgElem {
	return MsgElem{Type: BFACE, Data: MSG{"id": fmt.Sprintf("%d", id)}}
}

// NewSFace 生成Sface 的CQ码
func NewSFace(id int) MsgElem {
	return MsgElem{Type: SFACE, Data: MSG{"id": fmt.Sprintf("%d", id)}}
}

// NewRecord 生成路由CQ码 | file 语音文件名 | magic 发送时可选, 默认 0, 设置为 1 表示变声
func NewRecord(file string, magic int) MsgElem {
	return MsgElem{Type: RECORD, Data: MSG{"file": file, "magic": fmt.Sprintf("%d", magic)}}
}

// NewRps 猜拳魔法表情
func NewRps(rpstype string) MsgElem {
	return MsgElem{Type: RPS, Data: MSG{"type": rpstype}}
}

// NewDice 掷骰子魔法表情
func NewDice(dicType string) MsgElem {
	return MsgElem{Type: DICE, Data: MSG{"type": dicType}}
}

// NewShake 窗口抖动（戳一戳） 发
func NewShake() MsgElem {
	return MsgElem{Type: SHAKE}
}

// NewAnonymous 匿名发消息 发
func NewAnonymous(ignore bool) MsgElem {
	if ignore {
		return MsgElem{Type: ANONYMOUS, Data: MSG{"ignore": "true"}}
	}
	return MsgElem{Type: ANONYMOUS}
}

// NewMusic 生成音乐的CQ码
func NewMusic(mtype string, id string) MsgElem {
	return MsgElem{Type: MUSIC, Data: MSG{"type": mtype, "id": id}}
}

// NewDiyMusic 自定义音乐分享
func NewDiyMusic(url, audio, title, content, image string) MsgElem {
	return MsgElem{
		Type: MUSIC,
		Data: MSG{
			"type":    "custom",
			"url":     url,
			"audio":   audio,
			"title":   title,
			"content": content,
			"image":   image,
		},
	}
}

// NewShare 链接分享
func NewShare(url, title, content, image string) MsgElem {
	return MsgElem{Type: SHARE,
		Data: MSG{
			"url":     url,
			"title":   title,
			"content": content,
			"image":   image,
		}}
}

// NewVideo file 视频地址, 支持http和file发送 cover 视频封面, 支持http, file和base64发送, 格式必须为jpg
func NewVideo(file, cover string) MsgElem {
	return MsgElem{Type: VIDEO,
		Data: MSG{
			"file":  file,
			"cover": cover,
		}}
}

// NewReply 回复
func NewReply(id int) MsgElem {
	return MsgElem{Type: REPLY, Data: MSG{
		"id": fmt.Sprintf("%d", id),
	}}
}

// NewDiyReply 自定义回复
func NewDiyReply(text string, qq int64, time int64, seq int64) MsgElem {
	return MsgElem{Type: REPLY, Data: MSG{
		"text": text,
		"qq":   fmt.Sprintf("%d", qq),
		"time": fmt.Sprintf("%d", time),
		"seq":  fmt.Sprintf("%d", seq),
	}}
}

// NewPoke 群 戳一戳
func NewPoke(qq int64) MsgElem {
	return MsgElem{Type: POKE, Data: MSG{
		"qq": fmt.Sprintf("%d", qq),
	}}
}

// NewGift 礼物
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
func NewGift(qq int64, id int) MsgElem {
	return MsgElem{Type: GIFT, Data: MSG{
		"qq": fmt.Sprintf("%d", qq),
		"id": fmt.Sprintf("%d", id),
	}}
}

// NewCardImage 装逼大图
func NewCardImage(file string, minwidth, minheight, maxwidth, maxheight int64, source string, icon string) MsgElem {
	return MsgElem{Type: CARDIMAGE, Data: MSG{
		"file":      file,
		"minwidth":  fmt.Sprintf("%d", minwidth),
		"minheight": fmt.Sprintf("%d", minheight),
		"maxwidth":  fmt.Sprintf("%d", maxwidth),
		"maxheight": fmt.Sprintf("%d", maxheight),
		"source":    source,
		"icon":      icon,
	}}
}

// NewTTS 文本转语音 发
func NewTTS(text string) MsgElem {
	return MsgElem{Type: TTS, Data: MSG{
		"text": text,
	}}
}

// NewNode 合并转发。
func NewNode(id string) MsgElem {
	return MsgElem{
		Type: "node",
		Data: MSG{
			"id": id,
		},
	}
}
