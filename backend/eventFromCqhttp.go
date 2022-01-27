package backend

import (
	"bytes"
	"encoding/json"
	"github.com/scjtqs2/bot_adapter/config"
	"github.com/scjtqs2/bot_adapter/event"
	"github.com/scjtqs2/bot_adapter/sha256"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// EventResoveBackend 上报消息处理结构体
type EventResoveBackend struct {
	EventChan chan []byte
	Config    *config.Config
}

// DoEvent 处理事件
func (e *EventResoveBackend) DoEvent() {
	for msg := range e.EventChan {
		go e.doPlugsPush(msg)
	}
}

// doPlugsPush 分发数据到每个插件
func (e *EventResoveBackend) doPlugsPush(msg []byte) {
	for _, plugin := range e.Config.Plugins {
		if e.filterMsgKeyWords(msg, plugin.RegisterKeyWords) {
			message := gjson.ParseBytes(msg)
			log.Warnf("plugin %s 拦截事件：type=%s message_type=%s message_id=%s message=%s", plugin.PluginName, message.Get("post_type").String(), message.Get("message_type").String(), message.Get("message_id").String(), config.LimitedString(message.Get("message").String()))
			go e.doPush(msg, plugin)
			return // 这里直接拦截，所以不是continue
		}
	}
	// 无拦截，全部推送
	for _, plugin := range e.Config.Plugins {
		go e.doPush(msg, plugin)
	}
}

// parseEventPrimissCheck 解析消息内容 校验推送权限
func (e *EventResoveBackend) filterMsgKeyWords(data []byte, keywords []*config.RegisterKeyWords) bool {
	msg := gjson.ParseBytes(data)
	if msg.Get("post_type").String() == "message" && keywords != nil {
		message := msg.Get("message")
		if message.Type == gjson.JSON {
			node1 := message.Array()[0]
			if node1.Get("type").String() == "text" {
				return e.filterPrefixKeyWords(node1.Get("data.text").String(), msg.Get("message_type").String(), keywords)
			}
			return false
		}
		return e.filterPrefixKeyWords(message.String(), msg.Get("message_type").String(), keywords)
	}
	return false
}

// filterPrefixKeyWords 利用协程并发处理字符串前缀判断
func (e *EventResoveBackend) filterPrefixKeyWords(msg string, msg_type string, keys []*config.RegisterKeyWords) bool {
	flag := false
	if msg == "" || msg_type == "" || keys == nil {
		return false
	}
	wg := &sync.WaitGroup{}
	filter := func(f *bool, msg string, key []string, w *sync.WaitGroup) {
		defer w.Done()
		if *f {
			return
		}
		wg2 := &sync.WaitGroup{}
		for _, k := range key {
			keyword := k
			wg2.Add(1)
			go func(f *bool, msg, keyword string, wg2 *sync.WaitGroup) {
				defer wg2.Done()
				if *f {
					return
				}
				if strings.HasPrefix(msg, keyword) {
					*f = true
				}
			}(f, msg, keyword, wg2)
		}
		wg2.Wait()
	}
	for _, v := range keys {
		if v.MsgType == "" || v.MsgType != msg_type || v.PrefixKeyWords == nil {
			continue
		}
		wg.Add(1)
		go filter(&flag, msg, v.PrefixKeyWords, wg)
	}
	wg.Wait()
	return flag
}

func (e *EventResoveBackend) doPush(data []byte, plugin *config.PluginConfig) {
	if plugin.PostAddr == "" { // 空地址。不推送
		return
	}
	if !e.parseEventPrimissCheck(data, plugin) {
		return
	}
	enc, err := sha256.Encrypt(data, plugin.EncryptKey)
	if err != nil {
		log.Warnf("plugin %s encrypt msg faild err:%s", plugin.PluginName, err.Error())
		return
	}
	e.doPost(sha256.EncryptObj{Encrypt: enc}, plugin)
}

// parseEventPrimissCheck 解析消息内容 校验推送权限
func (e *EventResoveBackend) parseEventPrimissCheck(data []byte, plugin *config.PluginConfig) bool {
	msg := gjson.ParseBytes(data)
	// log.Infof("post_type %s", msg.Get("post_type").String())
	permissionOpt, err := config.GetPermission(e.Config.Permissions, plugin.AppID)
	if err != nil {
		// TODO 未配置推送权限，默认全开
		return true
	}
	if permissionOpt.IsAdmin || permissionOpt.IsOnlyCqhttp {
		return true
	}
	if permissionOpt.PushPermissions == nil {
		return false
	}
	switch msg.Get("post_type").String() {
	case "message": // 消息事件
		switch msg.Get("message_type").String() {
		case event.MESSAGE_TYPE_PRIVATE:
			// var req event.MessagePrivate
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_MESSAGE_PRIVATE, permissionOpt.PushPermissions)
		case event.MESSAGE_TYPE_GROUP:
			// var req event.MessageGroup
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_MESSAGE_GROUP, permissionOpt.PushPermissions)
		}
	case "notice": // 通知事件
		switch msg.Get("notice_type").String() {
		case event.NOTICE_TYPE_FRIEND_ADD:
			// var req event.NoticeFriendAdd
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_NOTICE_FRIEND_ADD, permissionOpt.PushPermissions)
		case event.NOTICE_TYPE_FRIEND_RECALL:
			// var req event.NoticeFriendRecall
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_NOTICE_FRIEND_RECALL, permissionOpt.PushPermissions)
		case event.NOTICE_TYPE_GROUP_BAN:
			// var req event.NoticeGroupBan
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_NOTICE_GROUP_BAN, permissionOpt.PushPermissions)
		case event.NOTICE_TYPE_GROUP_DECREASE:
			// var req event.NoticeGroupDecrease
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_NOTICE_GROUP_DECREASE, permissionOpt.PushPermissions)
		case event.NOTICE_TYPE_GROUP_INCREASE:
			// var req event.NoticeGroupIncrease
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_NOTICE_GROUP_INCREASE, permissionOpt.PushPermissions)
		case event.NOTICE_TYPE_GROUP_ADMIN:
			// var req event.NoticeGroupAdmin
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_NOTICE_GROUP_ADMIN, permissionOpt.PushPermissions)
		case event.NOTICE_TYPE_GROUP_RECALL:
			// var req event.NoticeGroupRecall
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_NOTICE_GROUP_RECALL, permissionOpt.PushPermissions)
		case event.NOTICE_TYPE_GROUP_UPLOAD:
			// var req event.NoticeGroupUpload
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_NOTICE_GROUP_UPLOAD, permissionOpt.PushPermissions)
		case event.NOTICE_TYPE_POKE:
			// var req event.NoticePoke
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_NOTICE_NOTIFY_POKE, permissionOpt.PushPermissions)
		case event.NOTICE_TYPE_HONOR:
			// var req event.NoticeHonor
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_NOTICE_NOTIFY_HONOR, permissionOpt.PushPermissions)
		case event.NOTICE_TYPE_LUCKY_KING:
			// var req event.NoticeLuckyKing
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_NOTICE_NOTIFY_LUCKY_KING, permissionOpt.PushPermissions)
		case event.CUSTOM_NOTICE_TYPE_GROUP_CARD:
			return config.CheckPermission(config.PUSH_CUSTOM_PERMISSION_NOTICE_GROUP_CARD, permissionOpt.PushPermissions)
		case event.CUSTOM_NOTICE_TYPE_OFFLINE_FILE:
			return config.CheckPermission(config.PUSH_CUSTOM_PERMISSION_NOTICE_OFFLINE_FILE, permissionOpt.PushPermissions)
		}
	case "request": // 请求事件
		switch msg.Get("request_type").String() {
		case event.REQUEST_TYPE_FRIEND:
			// var req event.RequestFriend
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_REQUEST_FRIEND, permissionOpt.PushPermissions)
		case event.REQUEST_TYPE_GROUP:
			// var req event.RequestGroup
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_REQUEST_GROUP, permissionOpt.PushPermissions)
		}
	case "meta_event": // 元事件
		switch msg.Get("meta_event_type").String() {
		case event.META_EVENT_LIFECYCLE:
			// var req event.MetaEventLifecycle
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_META_EVENT_LIFECYCLE, permissionOpt.PushPermissions)
		case event.META_EVENT_HEARTBEAT:
			// var req event.MetaEventHeartbeat
			// _ = json.Unmarshal([]byte(msg.Raw), &req)
			return config.CheckPermission(config.PUSH_PERMISSION_META_EVENT_HEARTBEAT, permissionOpt.PushPermissions)
		}
	}
	return false
}

// doPost 正真的推送
func (e *EventResoveBackend) doPost(obj sha256.EncryptObj, plugin *config.PluginConfig) {
	var res *http.Response
	body, _ := json.Marshal(obj)
	client := http.Client{Timeout: time.Second * time.Duration(time.Second*2)}
	header := make(http.Header)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	header.Set("X-Lark-Request-Timestamp", timestamp)
	nonce := sha256.RandString(8)
	header.Set("X-Lark-Request-Nonce", nonce)
	signature := sha256.CalculateSignature(timestamp, nonce, plugin.EncryptKey, string(body))
	header.Set("X-Lark-Signature", signature)
	header.Set("Content-Type", "application/json")
	maxRetries := 3
	for i := 0; i <= maxRetries; i++ {
		req, err := http.NewRequest("POST", plugin.PostAddr, bytes.NewReader(body))
		if err != nil {
			log.Warnf("推送数据到 %s  %s 时创建请求失败: %v", plugin.PluginName, plugin.PostAddr, err)
			return
		}
		req.Header = header
		res, err = client.Do(req)
		if res != nil {
			//goland:noinspection GoDeferInLoop
			defer res.Body.Close()
		}
		if err == nil {
			break
		}
		if i < maxRetries {
			log.Warnf("发送数据到 %s %s 失败: %v 将进行低 %d 次重试", plugin.PluginName, plugin.PostAddr, err, i+1)
		} else {
			log.Warnf("发送数据到 %s 数据 %s 到 %v 失败: %v 停止发送：已达重试上线", body, plugin.PluginName, plugin.PostAddr, err)
			return
		}
	}
}
