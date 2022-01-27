package event

// Event 事件的接口
type Event interface {
	GetType() string
}

const (
	EVENT_TYPE_MESSAGE_PRIVATE      = "message_private"
	EVENT_TYPE_MESSAGE_GROUP        = "message_group"
	EVENT_TYPE_META_HEARTBEAT       = "meta_heartbeat"
	EVENT_TYPE_META_LIFECYCLE       = "meta_lifecycle"
	EVENT_TYPE_NOTICE_GROUPUPLOAD   = "notice_groupupload"
	EVENT_TYPE_NOTICE_GROUPADMIIN   = "notice_groupadmin"
	EVENT_TYPE_NOTICE_GROUPDECREASE = "notice_groupdecrease"
	EVENT_TYPE_NOTICE_GROUPINCREASE = "notice_groupincrease"
	EVENT_TYPE_NOTICE_GROUPBAN      = "notice_groupban"
	EVENT_TYPE_NOTICE_GROUPRECAL    = "notice_groupcal"
	EVENT_TYPE_NOTICE_FRIENDADD     = "notice_friendadd"
	EVENT_TYPE_NOTICE_FRIENDRECALL  = "notice_friendrecall"
	EVENT_TYPE_NOTICE_POKE          = "notice_poke"
	EVENT_TYPE_NOTICE_HONOR         = "notice_honor"
	EVENT_TYPE_NOTICE_LUCKYKING     = "notice_luckyking"
	EVENT_TYPE_REQUEST_FRIEND       = "request_friend"
	EVENT_TYPE_REQUEST_GROUP        = "request_group"
)

// GetType 获取消息类型
func (e *MessagePrivate) GetType() string {
	return EVENT_TYPE_MESSAGE_PRIVATE
}

// GetType 获取消息类型
func (e *MessageGroup) GetType() string {
	return EVENT_TYPE_MESSAGE_GROUP
}

// GetType 获取消息类型
func (e *MetaEventLifecycle) GetType() string {
	return EVENT_TYPE_META_LIFECYCLE
}

// GetType 获取消息类型
func (e *MetaEventHeartbeat) GetType() string {
	return EVENT_TYPE_META_HEARTBEAT
}

// GetType 获取消息类型
func (e *RequestFriend) GetType() string {
	return EVENT_TYPE_REQUEST_FRIEND
}

// GetType 获取消息类型
func (e *RequestGroup) GetType() string {
	return EVENT_TYPE_REQUEST_GROUP
}

// GetType 获取消息类型
func (e *NoticeFriendAdd) GetType() string {
	return EVENT_TYPE_NOTICE_FRIENDADD
}

// GetType 获取消息类型
func (e *NoticeFriendRecall) GetType() string {
	return EVENT_TYPE_NOTICE_FRIENDRECALL
}

// GetType 获取消息类型
func (e *NoticeGroupAdmin) GetType() string {
	return EVENT_TYPE_NOTICE_GROUPADMIIN
}

// GetType 获取消息类型
func (e *NoticeGroupBan) GetType() string {
	return EVENT_TYPE_NOTICE_GROUPBAN
}

// GetType 获取消息类型
func (e *NoticeGroupDecrease) GetType() string {
	return EVENT_TYPE_NOTICE_GROUPDECREASE
}

// GetType 获取消息类型
func (e *NoticeGroupIncrease) GetType() string {
	return EVENT_TYPE_NOTICE_GROUPINCREASE
}

// GetType 获取消息类型
func (e *NoticeGroupUpload) GetType() string {
	return EVENT_TYPE_NOTICE_GROUPUPLOAD
}

// GetType 获取消息类型
func (e *NoticeGroupRecall) GetType() string {
	return EVENT_TYPE_NOTICE_GROUPRECAL
}

// GetType 获取消息类型
func (e *NoticeHonor) GetType() string {
	return EVENT_TYPE_NOTICE_HONOR
}

// GetType 获取消息类型
func (e *NoticeLuckyKing) GetType() string {
	return EVENT_TYPE_NOTICE_LUCKYKING
}

// GetType 获取消息类型
func (e *NoticePoke) GetType() string {
	return EVENT_TYPE_NOTICE_POKE
}
