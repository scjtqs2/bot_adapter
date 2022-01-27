package event

const (
	META_EVENT_LIFECYCLE = "lifecycle" // 生命周期
	META_EVENT_HEARTBEAT = "heartbeat" // 心跳
)

// MetaEventHeartbeat 心跳
type MetaEventHeartbeat struct {
	Time          int64            `json:"time"`            // 事件发生的时间戳
	SelfID        int64            `json:"self_id"`         // 收到事件的机器人 QQ 号
	PostType      string           `json:"post_type"`       // meta_event 上报类型
	MetaEventType string           `json:"meta_event_type"` // 元事件类型
	Interval      int64            `json:"interval"`        // 到下次心跳的间隔，单位毫秒
	Status        *HeartBeatStatus `json:"status"`          // 状态信息
}

// MetaEventLifecycle 生命周期
type MetaEventLifecycle struct {
	Time          int64  `json:"time"`            // 事件发生的时间戳
	SelfID        int64  `json:"self_id"`         // 收到事件的机器人 QQ 号
	PostType      string `json:"post_type"`       // meta_event 上报类型
	MetaEventType string `json:"meta_event_type"` // 元事件类型
	SubType       string `json:"sub_type"`        // enable、disable、connect 事件子类型，分别表示 OneBot 启用、停用、WebSocket 连接成功
}

// HeartBeatStatus 心跳的status字段
type HeartBeatStatus struct {
	AppEnabled     bool           `json:"app_enabled"`
	AppGood        bool           `json:"app_good"`
	AppInitialized bool           `json:"app_initialized"`
	Good           bool           `json:"good"`
	Online         bool           `json:"online"`
	Stat           *HeartBeatStat `json:"stat"`
}

// HeartBeatStat 心跳的 stat 字段
type HeartBeatStat struct {
	PacketReceived  int64 `json:"PacketReceived"`
	PacketSent      int64 `json:"PacketSent"`
	PacketLost      int64 `json:"PacketLost"`
	MessageReceived int64 `json:"MessageReceived"`
	MessageSent     int64 `json:"MessageSent"`
	LastMessageTime int64 `json:"LastMessageTime"`
	DisconnectTimes int64 `json:"DisconnectTimes"`
	LostTimes       int64 `json:"LostTimes"`
}
