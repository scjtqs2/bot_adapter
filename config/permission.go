package config

import "errors"

// api 权限 常量 get_token为公开接口，不需要权限
const (
	/**
	onebot-11 的标准接口
	*/
	PERMISSION_FOR_SEND_PRIVATE_MSG        = "send_private_msg"          // 发送私聊消息
	PERMISSION_FOR_SEND_GROUP_MSG          = "send_group_msg"            // 发送私聊消息
	PERMISSION_FOR_SEND_MSG                = "send_msg"                  // 发送消息
	PERMISSION_FOR_DELETE_MSG              = "delete_msg"                // 撤回消息
	PERMISSION_FOR_GET_MSG                 = "get_msg"                   // 获取消息
	PERMISSION_FOR_GET_FORWARD_MSG         = "get_forward_msg"           // 获取合并转发消息
	PERMISSION_FOR_SEND_LIKE               = "send_like"                 // 发送好友赞
	PERMISSION_FOR_SET_GROUP_KICK          = "set_group_kick"            // 群组踢人
	PERMISSION_FOR_SET_GROUP_BAN           = "set_group_ban"             // 群组单人禁言
	PERMISSION_FOR_SET_GROUP_ANONYMOUS_BAN = "set_group_anonymous_ban"   // 群组匿名用户禁言
	PERMISSION_FOR_SET_GROUP_WHOLE_BAN     = "set_group_group_whole_ban" // 群组全员禁言
	PERMISSION_FOR_SET_GROUP_ADMIN         = "set_group_admin"           // 群组设置管理员
	PERMISSION_FOR_SET_GROUP_ANONYMOUS     = "set_group_anonymous"       // 群组匿名
	PERMISSION_FOR_SET_GROUP_CARD          = "set_group_card"            // 设置群名片（备注）
	PERMISSION_FOR_SET_GROUP_NAME          = "set_group_name"            // 设置群名
	PERMISSION_FOR_SET_GROUP_LEAVE         = "set_group_leave"           // 退出群组
	PERMISSION_FOR_SET_GROUP_SPECIAL_TITLE = "set_group_special_title"   // 设置群组专属头衔
	PERMISSION_FOR_SET_FRIEND_ADD_REQUEST  = "set_friend_add_request"    // 处理好友请求
	PERMISSION_FOR_SET_GROUP_ADD_REQUEST   = "set_group_add_request"     // 处理加群请求/邀请
	PERMISSION_FOR_GET_LOGIN_INFO          = "get_login_info"            // 获取登录号信息
	PERMISSION_FOR_GET_STRANGER_INFO       = "get_stranger_info"         // 获取陌生人信息
	PERMISSION_FOR_GET_FRIEND_LIST         = "get_friend_list"           // 获取好友列表
	PERMISSION_FOR_GET_GROUP_INFO          = "get_group_info"            // 获取群信息
	PERMISSION_FOR_GET_GROUP_LIST          = "get_group_list"            // 获取群列表
	PERMISSION_FOR_GET_GROUP_MEMBER_INFO   = "get_group_member_info"     // 获取群成员信息
	PERMISSION_FOR_GET_GROUP_MEMBER_LIST   = "get_group_member_list"     // 获取群成员列表
	PERMISSION_FOR_GET_GROUP_HONOR_INFO    = "get_group_honor_info"      // 获取群荣誉信息
	PERMISSION_FOR_GET_COOKIES             = "get_cookies"               // 获取Cookies
	PERMISSION_FOR_GET_CSRF_TOKEN          = "get_csrf_token"            // 获取 CSRF Token
	PERMISSION_FOR_GET_CREDENTIALS         = "get_credentials"           // 获取 QQ 相关接口凭证
	PERMISSION_FOR_GET_RECORD              = "get_record"                // 获取语音
	PERMISSION_FOR_GET_IMAGE               = "get_image"                 // 获取图片
	PERMISSION_FOR_CAN_SEND_IMAGE          = "can_send_image"            // 检查是否可以发送图片
	PERMISSION_FOR_CAN_SEND_RECORD         = "can_send_record"           // 检查是否可以发送语音
	PERMISSION_FOR_GET_STATUS              = "get_status"                // 获取运行状态
	PERMISSION_FOR_GET_VERSION_INFO        = "get_version_info"          // 获取版本信息
	PERMISSION_FOR_SET_RESTART             = "set_restart"               // 重启 OneBot 实现
	PERMISSION_FOR_CLEAN_CACHE             = "clean_cache"               // 清理缓存

	// go-cqhttp实现的一些非 onebot-11 的接口 以及于 onebot标准略微差异的API
	PERMISSION_FOR_CUSTOM_SET_GROUP_PROTRAIT         = "custom_set_group_portrait"         // 设置群头像
	PERMISSION_FOR_CUSTOM_GET_WORD_SLICES            = "custom_get_word_slices"            // 获取中文分词（隐藏API)
	PERMISSION_FOR_CUSTOM_OCR_IMAGE                  = "custom_ocr_image"                  // 图片OCR
	PERMISSION_FOR_CUSTOM_GET_GROUP_SYSTEM_MSG       = "custom_get_group_system_msg"       // 获取群系统消息
	PERMISSION_FOR_CUSTOM_UPLOAD_GROUP_FILE          = "custom_upload_group_file"          // 上传群文件
	PERMISSION_FOR_CUSTOM_GET_GROUP_AT_ALL_REMAIN    = "custom_get_group_at_all_remain"    // 获取群@全体成员 剩余次数
	PERMISSION_FOR_CUSTOM_GET_VIP_INFO               = "custom_get_vip_info"               // 获取VIP信息
	PERMISSION_FOR_CUSTOM_SEND_GROUP_NOTICE          = "custom_send_group_notice"          // 发送群公告
	PERMISSION_FOR_RELOAD_EVENT_FILTER               = "custom_reload_event_filter"        // 重载事件过滤器
	PERMISSION_FOR_CUSTOM_SEND_GROUP_FORWARD_msg     = "custom_send_group_forward_msg"     // 发送合并转发（群）
	PERMISSION_FOR_CUSTOM_GET_GROUP_FILE_SYSTEM_INFO = "custom_get_group_file_system_info" // 获取群文件系统信息
	PERMISSION_FOR_CUSTOM_GET_GROUP_ROOT_FILES       = "custom_get_group_root_files"       // 获取群根目录文件列表
	PERMISSION_FOR_CUSTOM_GET_GROUP_FILES_BY_FOLDER  = "custom_get_group_files_by_folder"  // 获取群子目录文件列表
	PERMISSION_FOR_CUSTOM_GET_GROUP_FILE_URL         = "custom_get_group_file_url"         // 获取群文件资源连接
	PERMISSION_FOR_CUSTOM_DOWNLOAD_FILE              = "custom_download_file"              // 下载文件到缓存目录
	PERMISSION_FOR_CUSTOM_GET_ONLINE_CLIENTS         = "custom_get_online_clients"         // 获取当前账号在线客户端列表
	PERMISSION_FOR_CUSTOM_GET_GROUP_MSG_HISTORY      = "custom_get_group_msg_history"      // 获取群消息历史记录
	PERMISSION_FOR_CUSTOM_SET_ESSENCE_MSG            = "custom_set_essence_msg"            // 设置精华消息
	PERMISSION_FOR_CUSTOM_DELETE_ESSENCE_MSG         = "custom_delete_essence_msg"         // 移出精华消息
	PERMISSION_FOR_CUSTOM_GET_ESSENCE_MSG_LIST       = "custom_get_essence_msg_list"       // 获取精华消息列表
	PERMISSION_FOR_CUSTOM_CHECK_URL_SAFELY           = "custom_check_url_safely"           // 检查链接安全性
	PERMISSION_FOR_CUSTOM_GET_MODEL_SHOW             = "custom_get_model_show"             // 获取在线机型
	PERMISSION_FOR_CUSTOM_SET_MODEL_SHOW             = "custom_set_model_show"             // 设置在线机型

	// 其他功能
	PERMISSION_FOR_GET_CONFIG           = "get_config"           // 拉取 bot_adapter 的配置文件（yaml的bytes)
	PERMISSION_FOR_UPDATE_CONFIG        = "update_config"        // 直接更新覆盖 bot_adapter的配置
	PERMISSION_FOR_SET_BOT_ADAPTER_KILL = "set_bot_adapter_kill" // 用来重启 bot_adapter的docker，一般用于更新配置之后让其生效
)

// 推送事件 权限配置
const (
	/**
	消息类型事件
	*/
	PUSH_PERMISSION_MESSAGE_PRIVATE = "message_private"
	PUSH_PERMISSION_MESSAGE_GROUP   = "message_group"
	/**
	通知类事件
	*/
	PUSH_PERMISSION_NOTICE_GROUP_UPLOAD      = "notice_group_upload"
	PUSH_PERMISSION_NOTICE_GROUP_ADMIN       = "notice_group_admin"
	PUSH_PERMISSION_NOTICE_GROUP_DECREASE    = "notice_group_decrease"
	PUSH_PERMISSION_NOTICE_GROUP_INCREASE    = "notice_group_increase"
	PUSH_PERMISSION_NOTICE_GROUP_BAN         = "notice_group_ban"
	PUSH_PERMISSION_NOTICE_FRIEND_ADD        = "notice_friend_add"
	PUSH_PERMISSION_NOTICE_GROUP_RECALL      = "notice_group_recall"
	PUSH_PERMISSION_NOTICE_FRIEND_RECALL     = "notice_friend_recall"
	PUSH_PERMISSION_NOTICE_NOTIFY_POKE       = "notice_notify_poke"
	PUSH_PERMISSION_NOTICE_NOTIFY_LUCKY_KING = "notice_notify_lucky_king"
	PUSH_PERMISSION_NOTICE_NOTIFY_HONOR      = "notice_notify_honor"
	/**
	请求事件
	*/
	PUSH_PERMISSION_REQUEST_FRIEND = "request_friend"
	PUSH_PERMISSION_REQUEST_GROUP  = "request_group"
	/**
	元事件
	*/
	PUSH_PERMISSION_META_EVENT_LIFECYCLE = "meta_event_lifecycle"
	PUSH_PERMISSION_META_EVENT_HEARTBEAT = "meta_event_heartbeat"

	// go-cqhttp实现的扩展 event
	// https://docs.go-cqhttp.org/event/#%E7%BE%A4%E6%88%90%E5%91%98%E5%90%8D%E7%89%87%E6%9B%B4%E6%96%B0
	PUSH_CUSTOM_PERMISSION_NOTICE_GROUP_CARD = "custom_notice_group_card" // 群成员名片更新
	// https://docs.go-cqhttp.org/event/#%E6%8E%A5%E6%94%B6%E5%88%B0%E7%A6%BB%E7%BA%BF%E6%96%87%E4%BB%B6
	PUSH_CUSTOM_PERMISSION_NOTICE_OFFLINE_FILE = "custom_notice_offline_file" // 接收到的离线文件
)

// CheckOnlyCqhttp 校验是否是cqhttp 公开api的权限
func CheckOnlyCqhttp(permission string) bool {
	permissions := []string{
		PERMISSION_FOR_SEND_PRIVATE_MSG,
		PERMISSION_FOR_SEND_GROUP_MSG,
		PERMISSION_FOR_SEND_MSG,
		PERMISSION_FOR_DELETE_MSG,
		PERMISSION_FOR_GET_MSG,
		PERMISSION_FOR_GET_FORWARD_MSG,
		PERMISSION_FOR_SEND_LIKE,
		PERMISSION_FOR_SET_GROUP_KICK,
		PERMISSION_FOR_SET_GROUP_BAN,
		PERMISSION_FOR_SET_GROUP_ANONYMOUS_BAN,
		PERMISSION_FOR_SET_GROUP_WHOLE_BAN,
		PERMISSION_FOR_SET_GROUP_ADMIN,
		PERMISSION_FOR_SET_GROUP_ANONYMOUS,
		PERMISSION_FOR_SET_GROUP_CARD,
		PERMISSION_FOR_SET_GROUP_NAME,
		PERMISSION_FOR_SET_GROUP_LEAVE,
		PERMISSION_FOR_SET_GROUP_SPECIAL_TITLE,
		PERMISSION_FOR_SET_FRIEND_ADD_REQUEST,
		PERMISSION_FOR_SET_GROUP_ADD_REQUEST,
		PERMISSION_FOR_GET_LOGIN_INFO,
		PERMISSION_FOR_GET_STRANGER_INFO,
		PERMISSION_FOR_GET_FRIEND_LIST,
		PERMISSION_FOR_GET_GROUP_INFO,
		PERMISSION_FOR_GET_GROUP_LIST,
		PERMISSION_FOR_GET_GROUP_MEMBER_INFO,
		PERMISSION_FOR_GET_GROUP_MEMBER_LIST,
		PERMISSION_FOR_GET_GROUP_HONOR_INFO,
		PERMISSION_FOR_GET_COOKIES,
		PERMISSION_FOR_GET_CSRF_TOKEN,
		PERMISSION_FOR_GET_CREDENTIALS,
		PERMISSION_FOR_GET_RECORD,
		PERMISSION_FOR_GET_IMAGE,
		PERMISSION_FOR_CAN_SEND_IMAGE,
		PERMISSION_FOR_CAN_SEND_RECORD,
		PERMISSION_FOR_GET_STATUS,
		PERMISSION_FOR_GET_VERSION_INFO,
		PERMISSION_FOR_SET_RESTART,
		PERMISSION_FOR_CLEAN_CACHE,

		/**
		go-cqhttp的非标准api
		*/
		PERMISSION_FOR_CUSTOM_SET_GROUP_PROTRAIT,         // 设置群头像
		PERMISSION_FOR_CUSTOM_GET_WORD_SLICES,            // 获取中文分词（隐藏API)
		PERMISSION_FOR_CUSTOM_OCR_IMAGE,                  // 图片OCR
		PERMISSION_FOR_CUSTOM_GET_GROUP_SYSTEM_MSG,       // 获取群系统消息
		PERMISSION_FOR_CUSTOM_UPLOAD_GROUP_FILE,          // 上传群文件
		PERMISSION_FOR_CUSTOM_GET_GROUP_AT_ALL_REMAIN,    // 获取群@全体成员 剩余次数
		PERMISSION_FOR_CUSTOM_GET_VIP_INFO,               // 获取VIP信息
		PERMISSION_FOR_CUSTOM_SEND_GROUP_NOTICE,          // 发送群公告
		PERMISSION_FOR_RELOAD_EVENT_FILTER,               // 重载事件过滤器
		PERMISSION_FOR_CUSTOM_SEND_GROUP_FORWARD_msg,     // 发送合并转发（群）
		PERMISSION_FOR_CUSTOM_GET_GROUP_FILE_SYSTEM_INFO, // 获取群文件系统信息
		PERMISSION_FOR_CUSTOM_GET_GROUP_ROOT_FILES,       // 获取群根目录文件列表
		PERMISSION_FOR_CUSTOM_GET_GROUP_FILES_BY_FOLDER,  // 获取群子目录文件列表
		PERMISSION_FOR_CUSTOM_GET_GROUP_FILE_URL,         // 获取群文件资源连接
		PERMISSION_FOR_CUSTOM_DOWNLOAD_FILE,              // 下载文件到缓存目录
		PERMISSION_FOR_CUSTOM_GET_ONLINE_CLIENTS,         // 获取当前账号在线客户端列表
		PERMISSION_FOR_CUSTOM_GET_GROUP_MSG_HISTORY,      // 获取群消息历史记录
		PERMISSION_FOR_CUSTOM_SET_ESSENCE_MSG,            // 设置精华消息
		PERMISSION_FOR_CUSTOM_DELETE_ESSENCE_MSG,         // 移出精华消息
		PERMISSION_FOR_CUSTOM_GET_ESSENCE_MSG_LIST,       // 获取精华消息列表
		PERMISSION_FOR_CUSTOM_CHECK_URL_SAFELY,           // 检查链接安全性
		PERMISSION_FOR_CUSTOM_GET_MODEL_SHOW,             // 获取在线机型
		PERMISSION_FOR_CUSTOM_SET_MODEL_SHOW,             // 设置在线机型
	}
	for _, s := range permissions {
		if s == permission {
			return true
		}
	}
	return false
}

// GetPermission 从一堆权限配置中读取到指定的权限配置信息
func GetPermission(ps []*PermissionConfig, appid string) (*PermissionConfig, error) {
	for _, p := range ps {
		if p.AppID == appid {
			return p, nil
		}
	}
	return nil, errors.New("Permission NOT FOUND")
}

// CheckPermission 从集合ps中检测是否存在 权限 permission
func CheckPermission(permission string, ps []string) bool {
	for _, p := range ps {
		if p == permission {
			return true
		}
	}
	return false
}
