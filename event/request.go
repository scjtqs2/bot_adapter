package event

// 请求事件 加好友请求 加群请求/邀请

const (
	REQUEST_TYPE_FRIEND = "friend" // 加好友请求
	REQUEST_TYPE_GROUP  = "group"  // 加群请求/邀请
)

// RequestFriend 加好友请求
type RequestFriend struct {
	Time        int64  `json:"time"`         // 事件发生的时间戳
	SelfID      int64  `json:"self_id"`      // 收到事件的机器人QQ号
	PostType    string `json:"post_type"`    // 上报类型 request
	RequestType string `json:"request_type"` // friend 请求类型
	UserID      int64  `json:"user_id"`      // 请发送请求的QQ号
	Comment     string `json:"comment"`      // 验证信息
	Flag        string `json:"flag"`         // 请求 flag，在调用处理请求的 API 时需要传入
}

// RequestGroup 加群请求/邀请
type RequestGroup struct {
	Time        int64  `json:"time"`         // 事件发生的时间戳
	SelfID      int64  `json:"self_id"`      // 收到事件的机器人QQ号
	PostType    string `json:"post_type"`    // 上报类型 request
	RequestType string `json:"request_type"` // group 请求类型
	UserID      int64  `json:"user_id"`      // 请发送请求的QQ号
	Comment     string `json:"comment"`      // 验证信息
	Flag        string `json:"flag"`         // 请求 flag，在调用处理请求的 API 时需要传入
	GroupID     string `json:"group_id"`     // 群号
	SubType     string `json:"sub_type"`     // add invite 请求子类型，分别表示加群请求、邀请登录号入群
}
