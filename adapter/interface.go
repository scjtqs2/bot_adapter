// Package adapter 处理不同的协议渠道
package adapter

import (
	"errors"

	"github.com/scjtqs2/bot_adapter/config"
	"github.com/scjtqs2/bot_adapter/pb/entity"
)

// MSG 消息Map
type MSG map[string]interface{}

// AdapterSvc 适配器定义
type AdapterSvc interface {
	Send(action string, msg interface{}) ([]byte, error) // 统一的发送数据到go-cqhttp
	GetChan() chan []byte                                // 统一的接收上报消息
	// 下面的是 go-cqhttp的 onebot11的标准接口
	SendPrivateMsg(req *entity.SendPrivateMsgReq) (rsp *entity.SendMsgRsp, err error)
	SendGroupMsg(req *entity.SendGroupMsgReq) (rsp *entity.SendMsgRsp, err error)
	SendMsg(req *entity.SendMsgReq) (rsp *entity.SendMsgRsp, err error)
	DeleteMsg(req *entity.DeleteMsgReq) (rsp *entity.DeleteMsgRsp, err error)
	GetMsg(req *entity.GetMsgReq) (rsp *entity.GetMsgRsp, err error)
	GetForwardMsg(req *entity.GetForwardMsgReq) (rsp *entity.GetForwardMsgRsp, err error)
	SendLike(req *entity.SendLikeReq) (rsp *entity.SendLikeRsp, err error)
	SetGroupKick(req *entity.SetGroupKickReq) (rsp *entity.SetGroupKickRsp, err error)
	SetGroupBan(req *entity.SetGroupBanReq) (rsp *entity.SetGroupBanRsp, err error)
	SetGroupAnonymousBan(req *entity.SetGroupAnonymousBanReq) (rsp *entity.SetGroupAnonymousBanRsp, err error)
	SetGroupWholeBan(req *entity.SetGroupWholeBanReq) (rsp *entity.SetGroupWholeBanRsp, err error)
	SetGroupAdmin(req *entity.SetGroupAdminReq) (rsp *entity.SetGroupAdminRsp, err error)
	SetGroupAnonymous(req *entity.SetGroupAnonymousReq) (rsp *entity.SetGroupAnonymousRsp, err error)
	SetGroupCard(req *entity.SetGroupCardReq) (rsp *entity.SetGroupCardRsp, err error)
	SetGroupName(req *entity.SetGroupNameReq) (rsp *entity.SetGroupNameRsp, err error)
	SetGroupLeave(req *entity.SetGroupLeaveReq) (rsp *entity.SetGroupLeaveRsp, err error)
	SetGroupSpecialTitle(req *entity.SetGroupSpecialTitleReq) (rsp *entity.SetGroupSpecialTitleRsp, err error)
	SetGroupAddRequest(req *entity.SetGroupAddRequestReq) (rsp *entity.SetGroupAddRequestRsp, err error)
	SetFriendAddRequest(req *entity.SetFriendAddRequestReq) (rsp *entity.SetFriendaddRequestRsp, err error)
	GetLoginInfo(req *entity.GetLoginInfoReq) (rsp *entity.GetLoginInfoRsp, err error)
	GetStrangerInfo(req *entity.GetStrangerInfoReq) (rsp *entity.GetStrangerInfoRsp, err error)
	GetFriendList(req *entity.GetFriendListReq) (rsp *entity.GetFriendListRsp, err error)
	GetGroupInfo(req *entity.GetGroupInfoReq) (rsp *entity.GetGroupInfoRsp, err error)
	GetGroupList(req *entity.GetGroupListReq) (rsp *entity.GetGroupListRsp, err error)
	GetGroupMemberInfo(req *entity.GetGroupMemberInfoReq) (rsp *entity.GetGroupMemberInfoRsp, err error)
	GetGroupMemberList(req *entity.GetGroupMemberListReq) (rsp *entity.GetGroupMemberListRsp, err error)
	GetGroupHonorInfo(req *entity.GetGroupHonorInfoReq) (rsp *entity.GetGroupHonorInfoRsp, err error)
	GetCookies(req *entity.GetCookiesReq) (rsp *entity.GetCookiesRsp, err error)
	GetCsrfToken(req *entity.GetCsrfTokenReq) (rsp *entity.GetCsrfTokenRsp, err error)
	GetCredentials(req *entity.GetCredentialsReq) (rsp *entity.GetCredentialsRsp, err error)
	GetRecord(req *entity.GetRecordReq) (rsp *entity.GetRecordRsp, err error)
	GetImage(req *entity.GetImageReq) (rsp *entity.GetImageRsp, err error)
	CanSendImage(req *entity.CanSendImageReq) (rsp *entity.CanSendImageRsp, err error)
	CanSendRecord(req *entity.CanSendRecordReq) (rsp *entity.CanSendRecordRsp, err error)
	GetStatus(req *entity.GetStatusReq) (rsp *entity.GetStatusRsp, err error)
	GetVersionInfo(req *entity.GetVersionInfoReq) (rsp *entity.GetVersionInfoRsp, err error)
	SetRestart(req *entity.SetRestartReq) (rsp *entity.SetRestartRsp, err error)
	CleanCache(req *entity.CleanCacheReq) (rsp *entity.CleanCacheRsp, err error)

	// go-cqhttp 实现的一些非标准api
	CustomSetGroupPortrait(*entity.CustomSetGroupPortraitReq) (*entity.CustomSetGroupPortraitRsp, error)
	CustomGetWordSlices(*entity.CustomGetWordSlicesReq) (*entity.CustomGetWordSlicesRsp, error)
	CustomOcrImage(*entity.CustomOcrImageReq) (*entity.CustomOcrImageRsp, error)
	CustomGetGroupSystemMsg(*entity.CustomGetGroupSystemMsgReq) (*entity.CustomGetGroupSystemMsgRsp, error)
	CustomUploadGroupFile(*entity.CustomUploadGroupFileReq) (*entity.CustomUploadGroupFileRsp, error)
	CustomGetGroupFileSystemInfo(*entity.CustomGetGroupFileSystemInfoReq) (*entity.CustomGetGroupFileSystemInfoRsp, error)
	CustomGetGroupRootFiles(*entity.CustomGetGroupRootFilesReq) (*entity.CustomGetGroupRootFilesRsp, error)
	CustomGetGroupFilesByFolder(*entity.CustomGetGroupFilesByFolderReq) (*entity.CustomGetGroupFilesByFolderRsp, error)
	CustomGetGroupFileUrl(*entity.CustomGetGroupFileUrlReq) (*entity.CustomGetGroupFileUrlRsp, error)
	CustomGetStatus(*entity.CustomGetStatusReq) (*entity.CustomGetStatusRsp, error)
	CustomGetGroupAtAllRemain(*entity.CustomGetGroupAtAllRemainReq) (*entity.CustomGetGroupAtAllRemainRsp, error)
	CustomGetVipInfo(*entity.CustomGetVipInfoReq) (*entity.CustomGetVipInfoRsp, error)
	CustomSendGroupNotice(*entity.CustomSendGroupNoticeReq) (*entity.CustomSendGroupNoticeRsp, error)
	CustomReloadEventFilter(*entity.CustomReloadEventFilterReq) (*entity.CustomReloadEventFilterRsp, error)
	CustomDownloadFile(*entity.CustomDownloadFileReq) (*entity.CustomDownloadFileRsp, error)
	CustomGetOnlineClinets(*entity.CustomGetOnlineClientsReq) (*entity.CustomGetOnlineClientsRsp, error)
	CustomGetGroupMsgHistory(*entity.CustomGetGroupMsgHistoryReq) (*entity.CustomGetGroupMsgHistoryRsp, error)
	CustomSetEssenceMsg(*entity.CustomSetEssenceMsgReq) (*entity.CustomSetEssenceMsgRsp, error)
	CustomDeleteEssenceMsg(*entity.CustomDeleteEssenceMsgReq) (*entity.CustomDeleteEssenceMsgRsp, error)
	CustomGetEssenceMsgList(*entity.CustomGetEssenceMsgListReq) (*entity.CustomGetEssenceMsgListRsp, error)
	CustomCheckUrlSafely(*entity.CustomCheckUrlSafelyReq) (*entity.CustomCheckUrlSafelyRsp, error)
	CustomGetModelShow(*entity.CustomGetModelShowReq) (*entity.CustomGetModelShowRsp, error)
	CustomSetModelShow(*entity.CustomSetModelShowReq) (*entity.CustomSetModelShowRsp, error)
	CustomGetMsg(*entity.CustomGetMsgReq) (*entity.CustomGetMsgRsp, error)
	CustomGetImage(*entity.CustomGetImageReq) (*entity.CustomGetImageRsp, error)
	CustomSendGroupForwardMsg(*entity.CustomSendGroupForwardMsgReq) (*entity.CustomSendGroupForwardMsgRsp, error)
	CustomGetForwardMsg(*entity.CustomGetForwardMsgReq) (*entity.CustomGetForwardMsgRsp, error)
}

// NewAdapter 初始化适配器
func NewAdapter(cfg *config.Config) (AdapterSvc, error) {
	var (
		adpt AdapterSvc
		err  error
	)
	switch cfg.ServerType {
	case config.SERVER_TYPE_HTTP:
		adpt, err = HttpInit(cfg)
	default:
		err = errors.New("error adapter")
	}
	return adpt, err
}
