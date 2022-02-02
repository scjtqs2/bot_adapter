// Package client 对外提供的grpc client客户端封装 demo
package client

import (
	"context"
	"github.com/scjtqs2/bot_adapter/pb/entity"
	"github.com/scjtqs2/bot_adapter/pb/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"time"
)

type AdapterClientInterface interface {
	SendPrivateMsg(ctx context.Context, req *entity.SendPrivateMsgReq) (rsp *entity.SendMsgRsp, err error)
	SendGroupMsg(ctx context.Context, req *entity.SendGroupMsgReq) (rsp *entity.SendMsgRsp, err error)
	SendMsg(ctx context.Context, req *entity.SendMsgReq) (rsp *entity.SendMsgRsp, err error)
	DeleteMsg(ctx context.Context, req *entity.DeleteMsgReq) (rsp *entity.DeleteMsgRsp, err error)
	GetMsg(ctx context.Context, req *entity.GetMsgReq) (rsp *entity.GetMsgRsp, err error)
	GetForwardMsg(ctx context.Context, req *entity.GetForwardMsgReq) (rsp *entity.GetForwardMsgRsp, err error)
	SendLike(ctx context.Context, req *entity.SendLikeReq) (rsp *entity.SendLikeRsp, err error)
	SetGroupKick(ctx context.Context, req *entity.SetGroupKickReq) (rsp *entity.SetGroupKickRsp, err error)
	SetGroupBan(ctx context.Context, req *entity.SetGroupBanReq) (rsp *entity.SetGroupBanRsp, err error)
	SetGroupAnonymousBan(ctx context.Context, req *entity.SetGroupAnonymousBanReq) (rsp *entity.SetGroupAnonymousBanRsp, err error)
	SetGroupWholeBan(ctx context.Context, req *entity.SetGroupWholeBanReq) (rsp *entity.SetGroupWholeBanRsp, err error)
	SetGroupAdmin(ctx context.Context, req *entity.SetGroupAdminReq) (rsp *entity.SetGroupAdminRsp, err error)
	SetGroupAnonymous(ctx context.Context, req *entity.SetGroupAnonymousReq) (rsp *entity.SetGroupAnonymousRsp, err error)
	SetGroupCard(ctx context.Context, req *entity.SetGroupCardReq) (rsp *entity.SetGroupCardRsp, err error)
	SetGroupName(ctx context.Context, req *entity.SetGroupNameReq) (rsp *entity.SetGroupNameRsp, err error)
	SetGroupLeave(ctx context.Context, req *entity.SetGroupLeaveReq) (rsp *entity.SetGroupLeaveRsp, err error)
	SetGroupSpecialTitle(ctx context.Context, req *entity.SetGroupSpecialTitleReq) (rsp *entity.SetGroupSpecialTitleRsp, err error)
	SetGroupAddRequest(ctx context.Context, req *entity.SetGroupAddRequestReq) (rsp *entity.SetGroupAddRequestRsp, err error)
	SetFriendAddRequest(ctx context.Context, req *entity.SetFriendAddRequestReq) (rsp *entity.SetFriendaddRequestRsp, err error)
	GetLoginInfo(ctx context.Context, req *entity.GetLoginInfoReq) (rsp *entity.GetLoginInfoRsp, err error)
	GetStrangerInfo(ctx context.Context, req *entity.GetStrangerInfoReq) (rsp *entity.GetStrangerInfoRsp, err error)
	GetFriendList(ctx context.Context, req *entity.GetFriendListReq) (rsp *entity.GetFriendListRsp, err error)
	GetGroupInfo(ctx context.Context, req *entity.GetGroupInfoReq) (rsp *entity.GetGroupInfoRsp, err error)
	GetGroupList(ctx context.Context, req *entity.GetGroupListReq) (rsp *entity.GetGroupListRsp, err error)
	GetGroupMemberInfo(ctx context.Context, req *entity.GetGroupMemberInfoReq) (rsp *entity.GetGroupMemberInfoRsp, err error)
	GetGroupMemberList(ctx context.Context, req *entity.GetGroupMemberListReq) (rsp *entity.GetGroupMemberListRsp, err error)
	GetGroupHonorInfo(ctx context.Context, req *entity.GetGroupHonorInfoReq) (rsp *entity.GetGroupHonorInfoRsp, err error)
	GetCookies(ctx context.Context, req *entity.GetCookiesReq) (rsp *entity.GetCookiesRsp, err error)
	GetCsrfToken(ctx context.Context, req *entity.GetCsrfTokenReq) (rsp *entity.GetCsrfTokenRsp, err error)
	GetCredentials(ctx context.Context, req *entity.GetCredentialsReq) (rsp *entity.GetCredentialsRsp, err error)
	GetRecord(ctx context.Context, req *entity.GetRecordReq) (rsp *entity.GetRecordRsp, err error)
	GetImage(ctx context.Context, req *entity.GetImageReq) (rsp *entity.GetImageRsp, err error)
	CanSendImage(ctx context.Context, req *entity.CanSendImageReq) (rsp *entity.CanSendImageRsp, err error)
	CanSendRecord(ctx context.Context, req *entity.CanSendRecordReq) (rsp *entity.CanSendRecordRsp, err error)
	GetStatus(ctx context.Context, req *entity.GetStatusReq) (rsp *entity.GetStatusRsp, err error)
	GetVersionInfo(ctx context.Context, req *entity.GetVersionInfoReq) (rsp *entity.GetVersionInfoRsp, err error)
	SetRestart(ctx context.Context, req *entity.SetRestartReq) (rsp *entity.SetRestartRsp, err error)
	CleanCache(ctx context.Context, req *entity.CleanCacheReq) (rsp *entity.CleanCacheRsp, err error)

	// go-cqhttp 实现的一些非标准api
	CustomSetGroupPortrait(ctx context.Context, req *entity.CustomSetGroupPortraitReq) (*entity.CustomSetGroupPortraitRsp, error)
	CustomGetWordSlices(ctx context.Context, req *entity.CustomGetWordSlicesReq) (*entity.CustomGetWordSlicesRsp, error)
	CustomOcrImage(ctx context.Context, req *entity.CustomOcrImageReq) (*entity.CustomOcrImageRsp, error)
	CustomGetGroupSystemMsg(ctx context.Context, req *entity.CustomGetGroupSystemMsgReq) (*entity.CustomGetGroupSystemMsgRsp, error)
	CustomUploadGroupFile(ctx context.Context, req *entity.CustomUploadGroupFileReq) (*entity.CustomUploadGroupFileRsp, error)
	CustomGetGroupFileSystemInfo(ctx context.Context, req *entity.CustomGetGroupFileSystemInfoReq) (*entity.CustomGetGroupFileSystemInfoRsp, error)
	CustomGetGroupRootFiles(ctx context.Context, req *entity.CustomGetGroupRootFilesReq) (*entity.CustomGetGroupRootFilesRsp, error)
	CustomGetGroupFilesByFolder(ctx context.Context, req *entity.CustomGetGroupFilesByFolderReq) (*entity.CustomGetGroupFilesByFolderRsp, error)
	CustomGetGroupFileUrl(ctx context.Context, req *entity.CustomGetGroupFileUrlReq) (*entity.CustomGetGroupFileUrlRsp, error)
	CustomGetStatus(ctx context.Context, req *entity.CustomGetStatusReq) (*entity.CustomGetStatusRsp, error)
	CustomGetGroupAtAllRemain(ctx context.Context, req *entity.CustomGetGroupAtAllRemainReq) (*entity.CustomGetGroupAtAllRemainRsp, error)
	CustomGetVipInfo(ctx context.Context, req *entity.CustomGetVipInfoReq) (*entity.CustomGetVipInfoRsp, error)
	CustomSendGroupNotice(ctx context.Context, req *entity.CustomSendGroupNoticeReq) (*entity.CustomSendGroupNoticeRsp, error)
	CustomReloadEventFilter(ctx context.Context, req *entity.CustomReloadEventFilterReq) (*entity.CustomReloadEventFilterRsp, error)
	CustomDownloadFile(ctx context.Context, req *entity.CustomDownloadFileReq) (*entity.CustomDownloadFileRsp, error)
	CustomGetOnlineClinets(ctx context.Context, req *entity.CustomGetOnlineClientsReq) (*entity.CustomGetOnlineClientsRsp, error)
	CustomGetGroupMsgHistory(ctx context.Context, req *entity.CustomGetGroupMsgHistoryReq) (*entity.CustomGetGroupMsgHistoryRsp, error)
	CustomSetEssenceMsg(ctx context.Context, req *entity.CustomSetEssenceMsgReq) (*entity.CustomSetEssenceMsgRsp, error)
	CustomDeleteEssenceMsg(ctx context.Context, req *entity.CustomDeleteEssenceMsgReq) (*entity.CustomDeleteEssenceMsgRsp, error)
	CustomGetEssenceMsgList(ctx context.Context, req *entity.CustomGetEssenceMsgListReq) (*entity.CustomGetEssenceMsgListRsp, error)
	CustomCheckUrlSafely(ctx context.Context, req *entity.CustomCheckUrlSafelyReq) (*entity.CustomCheckUrlSafelyRsp, error)
	CustomGetModelShow(ctx context.Context, req *entity.CustomGetModelShowReq) (*entity.CustomGetModelShowRsp, error)
	CustomSetModelShow(ctx context.Context, req *entity.CustomSetModelShowReq) (*entity.CustomSetModelShowRsp, error)
	CustomGetMsg(ctx context.Context, req *entity.CustomGetMsgReq) (*entity.CustomGetMsgRsp, error)
	CustomGetImage(ctx context.Context, req *entity.CustomGetImageReq) (*entity.CustomGetImageRsp, error)
	CustomSendGroupForwardMsg(ctx context.Context, req *entity.CustomSendGroupForwardMsgReq) (*entity.CustomSendGroupForwardMsgRsp, error)
	CustomGetForwardMsg(ctx context.Context, req *entity.CustomGetForwardMsgReq) (*entity.CustomGetForwardMsgRsp, error)

	// 下面是bot_adapter自身的一些功能实现
	GetAuthToken(ctx context.Context, req *entity.GetAuthTokenReq) (rsp *entity.GetAuthTokenRsp, err error)
	GetConfig(ctx context.Context, req *entity.Config) (*entity.Config, error)
	UpdateConfig(ctx context.Context, req *entity.Config) (*entity.Config, error)
	SetBotAdapterKill(ctx context.Context, req *entity.SetRestartReq) (*entity.SetRestartRsp, error)
}

type AdapterService struct {
	client service.AdapterServiceClient
}

// NewAdapterServiceClient 新建grpc客户端
func NewAdapterServiceClient(addr string, appid string, appsecret string) (*AdapterService, error) {
	var ac AdapterService
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Timeout: 60 * time.Second,
		Time:    120 * time.Second,
	}))
	if err != nil {
		return nil, err
	}
	ac.client = service.NewAdapterServiceClient(conn)
	auth := &ac
	rsp, err := auth.client.GetAuthToken(context.TODO(), &entity.GetAuthTokenReq{AppId: appid, AppSecret: appsecret})
	if err != nil {
		return nil, err
	}
	token = rsp.Token
	// 自动刷新token
	go func() {
		for {
			rsp, err := auth.client.GetAuthToken(context.TODO(), &entity.GetAuthTokenReq{AppId: appid, AppSecret: appsecret})
			if err != nil {
				time.Sleep(time.Minute)
				return
			}
			token = rsp.Token
			time.Sleep(1 * time.Hour)
		}
	}()
	return auth, nil
}

// SendPrivateMsg 发送私聊
func (ac *AdapterService) SendPrivateMsg(ctx context.Context, req *entity.SendPrivateMsgReq) (rsp *entity.SendMsgRsp, err error) {
	return ac.client.SendPrivateMsg(GetContext(ctx), req)
}

// SendGroupMsg 发送群消息
func (ac *AdapterService) SendGroupMsg(ctx context.Context, req *entity.SendGroupMsgReq) (rsp *entity.SendMsgRsp, err error) {
	return ac.client.SendGroupMsg(GetContext(ctx), req)
}

// SendMsg 发送消息
func (ac *AdapterService) SendMsg(ctx context.Context, req *entity.SendMsgReq) (rsp *entity.SendMsgRsp, err error) {
	return ac.client.SendMsg(GetContext(ctx), req)
}

// DeleteMsg 撤回消息
func (ac *AdapterService) DeleteMsg(ctx context.Context, req *entity.DeleteMsgReq) (rsp *entity.DeleteMsgRsp, err error) {
	return ac.client.DeleteMsg(GetContext(ctx), req)
}

// GetMsg 拉取消息
func (ac *AdapterService) GetMsg(ctx context.Context, req *entity.GetMsgReq) (rsp *entity.GetMsgRsp, err error) {
	return ac.client.GetMsg(GetContext(ctx), req)
}

func (ac *AdapterService) GetForwardMsg(ctx context.Context, req *entity.GetForwardMsgReq) (rsp *entity.GetForwardMsgRsp, err error) {
	return ac.client.GetForwardMsg(GetContext(ctx), req)
}

func (ac *AdapterService) SendLike(ctx context.Context, req *entity.SendLikeReq) (rsp *entity.SendLikeRsp, err error) {
	return ac.client.SendLike(GetContext(ctx), req)
}

func (ac *AdapterService) SetGroupKick(ctx context.Context, req *entity.SetGroupKickReq) (rsp *entity.SetGroupKickRsp, err error) {
	return ac.client.SetGroupKick(GetContext(ctx), req)
}

func (ac *AdapterService) SetGroupBan(ctx context.Context, req *entity.SetGroupBanReq) (rsp *entity.SetGroupBanRsp, err error) {
	return ac.client.SetGroupBan(GetContext(ctx), req)
}

func (ac *AdapterService) SetGroupAnonymousBan(ctx context.Context, req *entity.SetGroupAnonymousBanReq) (rsp *entity.SetGroupAnonymousBanRsp, err error) {
	return ac.client.SetGroupAnonymousBan(GetContext(ctx), req)
}

func (ac *AdapterService) SetGroupWholeBan(ctx context.Context, req *entity.SetGroupWholeBanReq) (rsp *entity.SetGroupWholeBanRsp, err error) {
	return ac.client.SetGroupWholeBan(GetContext(ctx), req)
}

func (ac *AdapterService) SetGroupAdmin(ctx context.Context, req *entity.SetGroupAdminReq) (rsp *entity.SetGroupAdminRsp, err error) {
	return ac.client.SetGroupAdmin(GetContext(ctx), req)
}

func (ac *AdapterService) SetGroupAnonymous(ctx context.Context, req *entity.SetGroupAnonymousReq) (rsp *entity.SetGroupAnonymousRsp, err error) {
	return ac.client.SetGroupAnonymous(GetContext(ctx), req)
}

func (ac *AdapterService) SetGroupCard(ctx context.Context, req *entity.SetGroupCardReq) (rsp *entity.SetGroupCardRsp, err error) {
	return ac.client.SetGroupCard(GetContext(ctx), req)
}

func (ac *AdapterService) SetGroupName(ctx context.Context, req *entity.SetGroupNameReq) (rsp *entity.SetGroupNameRsp, err error) {
	return ac.client.SetGroupName(GetContext(ctx), req)
}

func (ac *AdapterService) SetGroupLeave(ctx context.Context, req *entity.SetGroupLeaveReq) (rsp *entity.SetGroupLeaveRsp, err error) {
	return ac.client.SetGroupLeave(GetContext(ctx), req)
}

func (ac *AdapterService) SetGroupSpecialTitle(ctx context.Context, req *entity.SetGroupSpecialTitleReq) (rsp *entity.SetGroupSpecialTitleRsp, err error) {
	return ac.client.SetGroupSpecialTitle(GetContext(ctx), req)
}

func (ac *AdapterService) SetGroupAddRequest(ctx context.Context, req *entity.SetGroupAddRequestReq) (rsp *entity.SetGroupAddRequestRsp, err error) {
	return ac.client.SetGroupAddRequest(GetContext(ctx), req)
}

func (ac *AdapterService) SetFriendAddRequest(ctx context.Context, req *entity.SetFriendAddRequestReq) (rsp *entity.SetFriendaddRequestRsp, err error) {
	return ac.client.SetFriendAddRequest(GetContext(ctx), req)
}

func (ac *AdapterService) GetLoginInfo(ctx context.Context, req *entity.GetLoginInfoReq) (rsp *entity.GetLoginInfoRsp, err error) {
	return ac.client.GetLoginInfo(GetContext(ctx), req)
}

func (ac *AdapterService) GetStrangerInfo(ctx context.Context, req *entity.GetStrangerInfoReq) (rsp *entity.GetStrangerInfoRsp, err error) {
	return ac.client.GetStrangerInfo(GetContext(ctx), req)
}

func (ac *AdapterService) GetFriendList(ctx context.Context, req *entity.GetFriendListReq) (rsp *entity.GetFriendListRsp, err error) {
	return ac.client.GetFriendList(GetContext(ctx), req)
}

func (ac *AdapterService) GetGroupInfo(ctx context.Context, req *entity.GetGroupInfoReq) (rsp *entity.GetGroupInfoRsp, err error) {
	return ac.client.GetGroupInfo(GetContext(ctx), req)
}

func (ac *AdapterService) GetGroupList(ctx context.Context, req *entity.GetGroupListReq) (rsp *entity.GetGroupListRsp, err error) {
	return ac.client.GetGroupList(GetContext(ctx), req)
}

func (ac *AdapterService) GetGroupMemberInfo(ctx context.Context, req *entity.GetGroupMemberInfoReq) (rsp *entity.GetGroupMemberInfoRsp, err error) {
	return ac.client.GetGroupMemberInfo(GetContext(ctx), req)
}

func (ac *AdapterService) GetGroupMemberList(ctx context.Context, req *entity.GetGroupMemberListReq) (rsp *entity.GetGroupMemberListRsp, err error) {
	return ac.client.GetGroupMemberList(GetContext(ctx), req)
}

func (ac *AdapterService) GetGroupHonorInfo(ctx context.Context, req *entity.GetGroupHonorInfoReq) (rsp *entity.GetGroupHonorInfoRsp, err error) {
	return ac.client.GetGroupHonorInfo(GetContext(ctx), req)
}

func (ac *AdapterService) GetCookies(ctx context.Context, req *entity.GetCookiesReq) (rsp *entity.GetCookiesRsp, err error) {
	return ac.client.GetCookies(GetContext(ctx), req)
}

func (ac *AdapterService) GetCsrfToken(ctx context.Context, req *entity.GetCsrfTokenReq) (rsp *entity.GetCsrfTokenRsp, err error) {
	return ac.client.GetCsrfToken(GetContext(ctx), req)
}

func (ac *AdapterService) GetCredentials(ctx context.Context, req *entity.GetCredentialsReq) (rsp *entity.GetCredentialsRsp, err error) {
	return ac.client.GetCredentials(GetContext(ctx), req)
}

func (ac *AdapterService) GetRecord(ctx context.Context, req *entity.GetRecordReq) (rsp *entity.GetRecordRsp, err error) {
	return ac.client.GetRecord(GetContext(ctx), req)
}

func (ac *AdapterService) GetImage(ctx context.Context, req *entity.GetImageReq) (rsp *entity.GetImageRsp, err error) {
	return ac.client.GetImage(GetContext(ctx), req)
}

func (ac *AdapterService) CanSendImage(ctx context.Context, req *entity.CanSendImageReq) (rsp *entity.CanSendImageRsp, err error) {
	return ac.client.CanSendImage(GetContext(ctx), req)
}

func (ac *AdapterService) CanSendRecord(ctx context.Context, req *entity.CanSendRecordReq) (rsp *entity.CanSendRecordRsp, err error) {
	return ac.client.CanSendRecord(GetContext(ctx), req)
}

func (ac *AdapterService) GetStatus(ctx context.Context, req *entity.GetStatusReq) (rsp *entity.GetStatusRsp, err error) {
	return ac.client.GetStatus(GetContext(ctx), req)
}

func (ac *AdapterService) GetVersionInfo(ctx context.Context, req *entity.GetVersionInfoReq) (rsp *entity.GetVersionInfoRsp, err error) {
	return ac.client.GetVersionInfo(GetContext(ctx), req)
}

func (ac *AdapterService) SetRestart(ctx context.Context, req *entity.SetRestartReq) (rsp *entity.SetRestartRsp, err error) {
	return ac.client.SetRestart(GetContext(ctx), req)
}

func (ac *AdapterService) CleanCache(ctx context.Context, req *entity.CleanCacheReq) (rsp *entity.CleanCacheRsp, err error) {
	return ac.client.CleanCache(GetContext(ctx), req)
}

func (ac *AdapterService) CustomSetGroupPortrait(ctx context.Context, req *entity.CustomSetGroupPortraitReq) (*entity.CustomSetGroupPortraitRsp, error) {
	return ac.client.CustomSetGroupPortrait(GetContext(ctx), req)
}

func (ac *AdapterService) CustomGetWordSlices(ctx context.Context, req *entity.CustomGetWordSlicesReq) (*entity.CustomGetWordSlicesRsp, error) {
	return ac.client.CustomGetWordSlices(GetContext(ctx), req)
}
func (ac *AdapterService) CustomOcrImage(ctx context.Context, req *entity.CustomOcrImageReq) (*entity.CustomOcrImageRsp, error) {
	return ac.client.CustomOcrImage(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetGroupSystemMsg(ctx context.Context, req *entity.CustomGetGroupSystemMsgReq) (*entity.CustomGetGroupSystemMsgRsp, error) {
	return ac.client.CustomGetGroupSystemMsg(GetContext(ctx), req)
}
func (ac *AdapterService) CustomUploadGroupFile(ctx context.Context, req *entity.CustomUploadGroupFileReq) (*entity.CustomUploadGroupFileRsp, error) {
	return ac.client.CustomUploadGroupFile(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetGroupFileSystemInfo(ctx context.Context, req *entity.CustomGetGroupFileSystemInfoReq) (*entity.CustomGetGroupFileSystemInfoRsp, error) {
	return ac.client.CustomGetGroupFileSystemInfo(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetGroupRootFiles(ctx context.Context, req *entity.CustomGetGroupRootFilesReq) (*entity.CustomGetGroupRootFilesRsp, error) {
	return ac.client.CustomGetGroupRootFiles(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetGroupFilesByFolder(ctx context.Context, req *entity.CustomGetGroupFilesByFolderReq) (*entity.CustomGetGroupFilesByFolderRsp, error) {
	return ac.client.CustomGetGroupFilesByFolder(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetGroupFileUrl(ctx context.Context, req *entity.CustomGetGroupFileUrlReq) (*entity.CustomGetGroupFileUrlRsp, error) {
	return ac.client.CustomGetGroupFileUrl(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetStatus(ctx context.Context, req *entity.CustomGetStatusReq) (*entity.CustomGetStatusRsp, error) {
	return ac.client.CustomGetStatus(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetGroupAtAllRemain(ctx context.Context, req *entity.CustomGetGroupAtAllRemainReq) (*entity.CustomGetGroupAtAllRemainRsp, error) {
	return ac.client.CustomGetGroupAtAllRemain(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetVipInfo(ctx context.Context, req *entity.CustomGetVipInfoReq) (*entity.CustomGetVipInfoRsp, error) {
	return ac.client.CustomGetVipInfo(GetContext(ctx), req)
}
func (ac *AdapterService) CustomSendGroupNotice(ctx context.Context, req *entity.CustomSendGroupNoticeReq) (*entity.CustomSendGroupNoticeRsp, error) {
	return ac.client.CustomSendGroupNotice(GetContext(ctx), req)
}
func (ac *AdapterService) CustomReloadEventFilter(ctx context.Context, req *entity.CustomReloadEventFilterReq) (*entity.CustomReloadEventFilterRsp, error) {
	return ac.client.CustomReloadEventFilter(GetContext(ctx), req)
}
func (ac *AdapterService) CustomDownloadFile(ctx context.Context, req *entity.CustomDownloadFileReq) (*entity.CustomDownloadFileRsp, error) {
	return ac.client.CustomDownloadFile(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetOnlineClinets(ctx context.Context, req *entity.CustomGetOnlineClientsReq) (*entity.CustomGetOnlineClientsRsp, error) {
	return ac.client.CustomGetOnlineClinets(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetGroupMsgHistory(ctx context.Context, req *entity.CustomGetGroupMsgHistoryReq) (*entity.CustomGetGroupMsgHistoryRsp, error) {
	return ac.client.CustomGetGroupMsgHistory(GetContext(ctx), req)
}
func (ac *AdapterService) CustomSetEssenceMsg(ctx context.Context, req *entity.CustomSetEssenceMsgReq) (*entity.CustomSetEssenceMsgRsp, error) {
	return ac.client.CustomSetEssenceMsg(GetContext(ctx), req)
}
func (ac *AdapterService) CustomDeleteEssenceMsg(ctx context.Context, req *entity.CustomDeleteEssenceMsgReq) (*entity.CustomDeleteEssenceMsgRsp, error) {
	return ac.client.CustomDeleteEssenceMsg(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetEssenceMsgList(ctx context.Context, req *entity.CustomGetEssenceMsgListReq) (*entity.CustomGetEssenceMsgListRsp, error) {
	return ac.client.CustomGetEssenceMsgList(GetContext(ctx), req)
}
func (ac *AdapterService) CustomCheckUrlSafely(ctx context.Context, req *entity.CustomCheckUrlSafelyReq) (*entity.CustomCheckUrlSafelyRsp, error) {
	return ac.client.CustomCheckUrlSafely(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetModelShow(ctx context.Context, req *entity.CustomGetModelShowReq) (*entity.CustomGetModelShowRsp, error) {
	return ac.client.CustomGetModelShow(GetContext(ctx), req)
}
func (ac *AdapterService) CustomSetModelShow(ctx context.Context, req *entity.CustomSetModelShowReq) (*entity.CustomSetModelShowRsp, error) {
	return ac.client.CustomSetModelShow(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetMsg(ctx context.Context, req *entity.CustomGetMsgReq) (*entity.CustomGetMsgRsp, error) {
	return ac.client.CustomGetMsg(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetImage(ctx context.Context, req *entity.CustomGetImageReq) (*entity.CustomGetImageRsp, error) {
	return ac.client.CustomGetImage(GetContext(ctx), req)
}
func (ac *AdapterService) CustomSendGroupForwardMsg(ctx context.Context, req *entity.CustomSendGroupForwardMsgReq) (*entity.CustomSendGroupForwardMsgRsp, error) {
	return ac.client.CustomSendGroupForwardMsg(GetContext(ctx), req)
}
func (ac *AdapterService) CustomGetForwardMsg(ctx context.Context, req *entity.CustomGetForwardMsgReq) (*entity.CustomGetForwardMsgRsp, error) {
	return ac.client.CustomGetForwardMsg(GetContext(ctx), req)
}

// 下面是bot_adapter自身的一些功能实现
func (ac *AdapterService) GetAuthToken(ctx context.Context, req *entity.GetAuthTokenReq) (rsp *entity.GetAuthTokenRsp, err error) {
	return ac.client.GetAuthToken(GetContext(ctx), req)
}
func (ac *AdapterService) GetConfig(ctx context.Context, req *entity.Config) (*entity.Config, error) {
	return ac.client.GetConfig(GetContext(ctx), req)
}
func (ac *AdapterService) UpdateConfig(ctx context.Context, req *entity.Config) (*entity.Config, error) {
	return ac.client.UpdateConfig(GetContext(ctx), req)
}
func (ac *AdapterService) SetBotAdapterKill(ctx context.Context, req *entity.SetRestartReq) (*entity.SetRestartRsp, error) {
	return ac.client.SetBotAdapterKill(GetContext(ctx), req)
}
