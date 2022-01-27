package config

// api的url常量定义
const (
	API_URL_SEND_PRIVATE_MSG        = "send_private_msg"        // 发送私聊消息
	API_URL_SEND_GROUP_MSG          = "send_group_msg"          // 发送群消息
	API_URL_SEND_MSG                = "send_msg"                // 发送消息
	API_URL_DELETE_MSG              = "delete_msg"              // 撤回消息
	API_URL_GET_MSG                 = "get_msg"                 // 获取消息
	API_URL_GET_FORWARD_MSG         = "get_forward_msg"         // 获取合并转发消息
	API_URL_SEND_LIKE               = "send_like"               // 发送好友赞
	API_URL_SET_GROUP_KICK          = "set_group_kick"          // 群组踢人
	API_URL_SET_GROUP_BAN           = "set_group_ban"           // 群组单人禁言
	API_URL_SET_GROUP_ANONYMOUS_BAN = "set_group_anonymous_ban" // 群组匿名用户禁言
	API_URL_SET_GROUP_WHOLE_BAN     = "set_group_whole_ban"     // 群组全员禁言
	API_URL_SET_GROUP_ADMIN         = "set_group_admin"         // 群组设置管理员
	API_URL_SET_GROUP_ANONYMOUS     = "set_group_anonymous"     // 群组匿名
	API_URL_SET_GROUP_CARD          = "set_group_card"          // 设置群名片（备注）
	API_URL_SET_GROUP_NAME          = "set_group_name"          // 设置群名
	API_URL_SET_GROUP_LEAVE         = "set_group_leave"         // 退出群组
	API_URL_SET_GROUP_SPECIAL_TITLE = "set_group_special_title" // 设置群组专属头衔
	API_URL_SET_FRIEND_ADD_REQUEST  = "set_friend_add_request"  // 处理好友请求
	API_URL_SET_GROUP_ADD_REQUEST   = "set_group_add_request"   // 处理加群请求/邀请
	API_URL_GET_LOGIN_INFO          = "get_login_info"          // 获取登录号信息
	API_URL_GET_STRANGER_INFO       = "get_stranger_info"       // 获取陌生人信息
	API_URL_GET_FRIEND_LIST         = "get_friend_list"         // 获取好友列表
	API_URL_GET_GROUP_INFO          = "get_group_info"          // 获取群信息
	API_URL_GET_GROUP_LIST          = "get_group_list"          // 获取群列表
	API_URL_GET_GROUP_MEMBER_INFO   = "get_group_member_info"   // 获取群成员信息
	API_URL_GET_GROUP_MEMBER_LIST   = "get_group_member_list"   // 获取群成员列表
	API_URL_GET_GROUP_HONOR_INFO    = "get_group_honor_info"    // 获取群荣誉信息
	API_URL_GET_COOKIES             = "get_cookies"             // 获取Cookies
	API_URL_GET_CSRF_TOKEN          = "get_csrf_token"          // 获取 CSRF Token
	API_URL_GET_CREDENTIALS         = "get_credentials"         // 获取 QQ 相关接口凭证
	API_URL_GET_RECORD              = "get_record"              // 获取语音
	API_URL_GET_IMAGE               = "get_image"               // 获取图片
	API_URL_CAN_SEND_IMAGE          = "can_send_image"          // 检查是否可以发送图片
	API_URL_CAN_SEND_RECORD         = "can_send_record"         // 检查是否可以发送语音
	API_URL_GET_STATUS              = "get_status"              // 获取运行状态
	API_URL_GET_VERSION_INFO        = "get_version_info"        // 获取版本信息
	API_URL_SET_RESTART             = "set_restart"             // 重启 OneBot 实现
	API_URL_CLEAN_CACHE             = "clean_cache"             // 清理缓存

	// go-cqhttp 实现的一些和标准不一样的api

	API_CUSTOM_URL_SET_GROUP_PORTRAIT         = "set_group_portrait"         // 设置群头像
	API_CUSTOM_URL_GET_WORD_SLICES            = ".get_word_slices"           // 获取中文分词（隐藏API)
	API_CUSTOM_URL_OCR_IMAGE                  = "ocr_image"                  // 图片OCR
	API_CUSTOM_URL_GET_GROUP_SYSTEM_MSG       = "get_group_system_msg"       // 获取群系统消息
	API_CUSTOM_URL_UPLOAD_GROUP_FILE          = "upload_group_file"          // 上传群文件
	API_CUSTOM_URL_GET_GROUP_FILE_SYSTEM_INFO = "get_group_file_system_info" // 获取群文件系统信息
	API_CUSTOM_URL_GET_GROUP_ROOT_FILES       = "get_group_root_files"       // 获取群根目录文件列表
	API_CUSTOM_URL_GET_GROUP_FILES_BY_FOLDER  = "get_group_files_by_folder"  // 获取群子目录文件列表
	API_CUSTOM_URL_GET_GROUP_FILE_URL         = "get_group_file_url"         // 获取群文件资源链接
	API_CUSTOM_URL_GET_STATUS                 = "get_status"                 // 获取状态
	API_CUSTOM_URL_GET_GROUP_AT_ALL_REMAIN    = "get_group_at_all_remain"    // 获取群@全体成员 剩余次数
	API_CUSTOM_URL_GET_VIP_INFO               = "_get_vip_info"              // 获取VIP信息
	API_CUSTOM_URL_SEND_GROUP_NOTICE          = "_send_group_notice"         // 发送群公告
	API_CUSTOM_URL_RELOAD_EVENT_FILTER        = "reload_event_filter"        // 重载事件过滤器
	API_CUSTOM_URL_DOWNLOAD_FILE              = "download_file"              // 下载文件到缓存目录
	API_CUSTOM_URL_GET_ONLINE_CLIENTS         = "get_online_clients"         // 获取当前账号在线客户端列表
	API_CUSTOM_URL_GET_GROUP_MSG_HISTORY      = "get_group_msg_history"      // 获取群消息历史记录
	API_CUSTOM_URL_SET_ESSENCE_MSG            = "set_essence_msg"            // 设置精华消息
	API_CUSTOM_URL_DELETE_ESSENCE_MSG         = "delete_essence_msg"         // 移除精华消息
	API_CUSTOM_URL_GET_ESSENCE_MSG_LIST       = "get_essence_msg_list"       // 获取精华消息列表
	API_CUSTOM_URL_CHECK_URL_SAFELY           = "check_url_safely"           // 检查链接安全性
	API_CUSTOM_URL_GET_MODEL_SHOW             = "_get_model_show"            // 获取在线机型
	API_CUSTOM_URL_SET_MODEL_SHOW             = "_set_model_show"            // 设置在线机型
	API_CUSTOM_URL_SEND_GROUP_FORWARD_MSG     = "send_group_forward_msg"     // 发送合并转发（群）
	API_CUSTOM_URL_GET_MSG                    = "get_msg"                    // 获取消息
	API_CUSTOM_URL_GET_IMAGE                  = "get_image"                  // 获取图片信息
	API_CUSTOM_URL_GET_FORWARD_MSG            = "get_forward_msg"            // 获取并转发内容

)
