// Package posttype  https://github.com/botuniverse/onebot-11/tree/master/event
package event

const (
	MESSAGE_TYPE_PRIVATE = "private" // 私聊消息
	MESSAGE_TYPE_GROUP   = "group"   // 群消息
)

// MessagePrivate 私聊消息
type MessagePrivate struct {
	Time        int64          `json:"time"`         // 时间戳
	SelfID      int64          `json:"self_id"`      // 收到事件的机器人 QQ 号
	PostType    string         `json:"post_type"`    // message 上报类型
	MessageType string         `json:"message_type"` // private 消息类型
	SubType     string         `json:"sub_type"`     // friend、group、other  消息子类型，如果是好友则是 friend，如果是群临时会话则是 group
	MessageID   int64          `json:"message_id"`   // 消息 ID
	UserID      int64          `json:"user_id"`      // 发送者 QQ 号
	Message     string         `json:"message"`      // 消息内容 string 或 array
	RawMessage  string         `json:"raw_message"`  // 原始消息内容
	Font        int64          `json:"font"`         // 0 字体
	Sender      *MessageSender `json:"sender"`       // 发送人信息
}

// MessageGroup 群消息
type MessageGroup struct {
	Time        int64             `json:"time"`         // 时间戳
	SelfID      int64             `json:"self_id"`      // 收到事件的机器人 QQ 号
	PostType    string            `json:"post_type"`    // message 上报类型
	MessageType string            `json:"message_type"` // group 消息类型
	SubType     string            `json:"sub_type"`     // normal、anonymous、notice  消息子类型，正常消息是 normal，匿名消息是 anonymous，系统提示（如「管理员已禁止群内匿名聊天」）是 notice
	MessageID   int64             `json:"message_id"`   // 消息 ID
	GroupID     int64             `json:"group_id"`     // 群号
	UserID      int64             `json:"user_id"`      // 发送者 QQ 号
	Anonymous   *MessageAnonymous `json:"anonymous"`    // 匿名信息，如果不是匿名消息则为 null
	Message     string            `json:"message"`      // 消息内容 string 或 array
	RawMessage  string            `json:"raw_message"`  // 原始消息内容
	Font        int64             `json:"font"`         // 0 字体
	Sender      *MessageSender    `json:"sender"`       // 发送人信息
}

// MessageAnonymous 群消息的 annoymous
type MessageAnonymous struct {
	Flag string `json:"flag"` // 匿名用户 flag，在调用禁言 API 时需要传入
	ID   int64  `json:"id"`   // 匿名用户 ID
	Name string `json:"name"` // 匿名用户名称
}

// MessageSender 消息发送者的信息
type MessageSender struct {
	Age      int32  `json:"age"`      // 年龄
	Area     string `json:"area"`     // (群消息) 地区
	Card     string `json:"card"`     // (群消息) 群名片／备注
	Level    string `json:"level"`    // (群消息) 成员等级
	NickName string `json:"nickname"` // 昵称
	Role     string `json:"role"`     // (群消息) 	角色，owner 或 admin 或 member
	Sex      string `json:"sex"`      // 性别，male 或 female 或 unknown
	Title    string `json:"title"`    // (群消息) 	专属头衔
	UserID   int64  `json:"user_id"`  // 发送者 QQ 号
}
