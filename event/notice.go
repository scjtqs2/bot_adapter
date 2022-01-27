package event

const (
	NOTICE_TYPE_GROUP_UPLOAD   = "group_upload"   // 群文件上传
	NOTICE_TYPE_GROUP_ADMIN    = "group_admin"    // 群管理员变动
	NOTICE_TYPE_GROUP_DECREASE = "group_decrease" // 群成员减少
	NOTICE_TYPE_GROUP_INCREASE = "group_increase" // 群成员增加
	NOTICE_TYPE_GROUP_BAN      = "group_ban"      // 群禁言
	NOTICE_TYPE_FRIEND_ADD     = "friend_add"     // 好友添加
	NOTICE_TYPE_GROUP_RECALL   = "group_recall"   // 群消息撤回
	NOTICE_TYPE_FRIEND_RECALL  = "friend_recall"  // 好友消息撤回
	NOTICE_TYPE_POKE           = "poke"           // 群内戳一戳
	NOTICE_TYPE_LUCKY_KING     = "lucky_king"     // 群红包运气王
	NOTICE_TYPE_HONOR          = "honor"          // 群成员荣誉变更

	// go-cqhttp实现的扩展类型
	// https://docs.go-cqhttp.org/event/#%E7%BE%A4%E6%88%90%E5%91%98%E5%90%8D%E7%89%87%E6%9B%B4%E6%96%B0
	CUSTOM_NOTICE_TYPE_GROUP_CARD = "group_card" // 群成员名片更新 ps 此事件不保证时效性, 仅在收到消息时校验卡片
	// https://docs.go-cqhttp.org/event/#%E6%8E%A5%E6%94%B6%E5%88%B0%E7%A6%BB%E7%BA%BF%E6%96%87%E4%BB%B6
	CUSTOM_NOTICE_TYPE_OFFLINE_FILE = "offline_file" // 接收到离线文件
)

// NoticeGroupUpload 群文件上传
type NoticeGroupUpload struct {
	Time       int64       `json:"time"`        // 事件发生的时间戳
	SelfID     int64       `json:"self_id"`     // 收到事件的机器人QQ
	PostType   string      `json:"post_type"`   // 上报类型 notice
	NoticeType string      `json:"notice_type"` // 通知类型 group_upload
	GroupID    int64       `json:"group_id"`    // 群号
	UserID     int64       `json:"user_id"`     // 发送者QQ号
	File       *NoticeFile `json:"file"`        // 文件信息
}

// NoticeFile 文件信息
type NoticeFile struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Busid int64  `json:"busid"`
}

// NoticeGroupAdmin 群管理员变动
type NoticeGroupAdmin struct {
	Time       int64  `json:"time"`        // 事件发生的时间戳
	SelfID     int64  `json:"self_id"`     // 收到事件的机器人QQ
	PostType   string `json:"post_type"`   // notice 上报类型
	NoticeType string `json:"notice_type"` // group_admin 通知类型
	SubType    string `json:"sub_type"`    // set unset 事件子类型，分别表示设置和取消管理员
	GroupID    int64  `json:"group_id"`    // 群号
	UserID     int64  `json:"user_id"`     // 管理员QQ号
}

// NoticeGroupDecrease 群成员减少
type NoticeGroupDecrease struct {
	Time       int64  `json:"time"`        // 事件发生的时间戳
	SelfID     int64  `json:"self_id"`     // 收到事件的机器人QQ
	PostType   string `json:"post_type"`   // notice 上报类型
	NoticeType string `json:"notice_type"` // group_decrease 通知类型
	SubType    string `json:"sub_type"`    // leave、kick、kick_me 事件子类型，分别表示主动退群、成员被踢、登录号被踢
	GroupID    int64  `json:"group_id"`    // 群号
	OperatorID int64  `json:"operator_id"` // 操作者 QQ 号（如果是主动退群，则和 user_id 相同）
	UserID     int64  `json:"user_id"`     // 离开者 QQ 号
}

// NoticeGroupIncrease 群成员增加
type NoticeGroupIncrease struct {
	Time       int64  `json:"time"`        // 事件发生的时间戳
	SelfID     int64  `json:"self_id"`     // 收到事件的机器人QQ
	PostType   string `json:"post_type"`   // notice 上报类型
	NoticeType string `json:"notice_type"` // group_increase 通知类型
	SubType    string `json:"sub_type"`    // approve、invite 事件子类型，分别表示管理员已同意入群、管理员邀请入群
	GroupID    int64  `json:"group_id"`    // 群号
	OperatorID int64  `json:"operator_id"` // 操作者 QQ 号
	UserID     int64  `json:"user_id"`     // 离开者 QQ 号
}

// NoticeGroupBan 群禁言
type NoticeGroupBan struct {
	Time       int64  `json:"time"`        // 事件发生的时间戳
	SelfID     int64  `json:"self_id"`     // 收到事件的机器人QQ
	PostType   string `json:"post_type"`   // notice 上报类型
	NoticeType string `json:"notice_type"` // group_ban	 通知类型
	SubType    string `json:"sub_type"`    // ban、lift_ban 事件子类型，分别表示禁言、解除禁言
	GroupID    int64  `json:"group_id"`    // 群号
	OperatorID int64  `json:"operator_id"` // 操作者 QQ 号
	UserID     int64  `json:"user_id"`     // 被禁言 QQ 号
	Duration   int64  `json:"duration"`    // 禁言时长，单位秒
}

// NoticeFriendAdd 好友添加
type NoticeFriendAdd struct {
	Time       int64  `json:"time"`        // 事件发生的时间戳
	SelfID     int64  `json:"self_id"`     // 收到事件的机器人QQ
	PostType   string `json:"post_type"`   // notice 上报类型
	NoticeType string `json:"notice_type"` // friend_add 通知类型
	UserID     int64  `json:"user_id"`     // 被禁言 QQ 号
}

// NoticeGroupRecall 群消息撤回
type NoticeGroupRecall struct {
	Time       int64  `json:"time"`        // 事件发生的时间戳
	SelfID     int64  `json:"self_id"`     // 收到事件的机器人QQ
	PostType   string `json:"post_type"`   // notice 上报类型
	NoticeType string `json:"notice_type"` // group_recall	 通知类型
	GroupID    int64  `json:"group_id"`    // 群号
	OperatorID int64  `json:"operator_id"` // 操作者 QQ 号
	UserID     int64  `json:"user_id"`     // 消息发送者 QQ 号
	MessageID  int64  `json:"message_id"`  // 被撤回的消息 ID
}

// NoticeFriendRecall 好友消息撤回
type NoticeFriendRecall struct {
	Time       int64  `json:"time"`        // 事件发生的时间戳
	SelfID     int64  `json:"self_id"`     // 收到事件的机器人QQ
	PostType   string `json:"post_type"`   // notice 上报类型
	NoticeType string `json:"notice_type"` // friend_recall	 通知类型
	UserID     int64  `json:"user_id"`     // 消息发送者 QQ 号
	MessageID  int64  `json:"message_id"`  // 被撤回的消息 ID
}

// NoticePoke 群内戳一戳
type NoticePoke struct {
	Time       int64  `json:"time"`        // 事件发生的时间戳
	SelfID     int64  `json:"self_id"`     // 收到事件的机器人QQ
	PostType   string `json:"post_type"`   // notice 上报类型
	NoticeType string `json:"notice_type"` // notify 通知类型
	SubType    string `json:"sub_type"`    // poke 提示类型
	GroupID    int64  `json:"group_id"`    // 群号
	UserID     int64  `json:"user_id"`     // 发送者 QQ 号
	TargetID   int64  `json:"target_id"`   // 被戳者 QQ 号
}

// NoticeLuckyKing 群红包运气王
type NoticeLuckyKing struct {
	Time       int64  `json:"time"`        // 事件发生的时间戳
	SelfID     int64  `json:"self_id"`     // 收到事件的机器人QQ
	PostType   string `json:"post_type"`   // notice 上报类型
	NoticeType string `json:"notice_type"` // notify 通知类型
	SubType    string `json:"sub_type"`    // lucky_king 提示类型
	GroupID    int64  `json:"group_id"`    // 群号
	UserID     int64  `json:"user_id"`     // 红包发送者 QQ 号
	TargetID   int64  `json:"target_id"`   // 运气王 QQ 号
}

// NoticeHonor 群成员荣誉变更
type NoticeHonor struct {
	Time       int64  `json:"time"`        // 事件发生的时间戳
	SelfID     int64  `json:"self_id"`     // 收到事件的机器人QQ
	PostType   string `json:"post_type"`   // notice 上报类型
	NoticeType string `json:"notice_type"` // notify 通知类型
	SubType    string `json:"sub_type"`    // honor 提示类型
	GroupID    int64  `json:"group_id"`    // 群号
	UserID     int64  `json:"user_id"`     // 成员 QQ 号
	HonorType  string `json:"honor_type"`  // talkative、performer、emotion  荣誉类型，分别表示龙王、群聊之火、快乐源泉
}
