package backend

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"github.com/scjtqs2/bot_adapter/auth"
	"github.com/scjtqs2/bot_adapter/config"
	"github.com/scjtqs2/bot_adapter/pb/entity"
	"github.com/scjtqs2/bot_adapter/pb/service"
)

// AuthorizedService grpc的jwt 验证
type AuthorizedService struct {
	svc   service.AdapterServiceServer
	conf  *config.Config
	ausvc *auth.AuthServer
}

func (a AuthorizedService) CustomSetGroupPortrait(ctx context.Context, req *entity.CustomSetGroupPortraitReq) (*entity.CustomSetGroupPortraitRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomSetGroupPortrait token error:%v", err)
		return nil, err
	}
	return a.svc.CustomSetGroupPortrait(ctx, req)
}

func (a AuthorizedService) CustomGetWordSlices(ctx context.Context, req *entity.CustomGetWordSlicesReq) (*entity.CustomGetWordSlicesRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetWordSlices token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetWordSlices(ctx, req)
}

func (a AuthorizedService) CustomOcrImage(ctx context.Context, req *entity.CustomOcrImageReq) (*entity.CustomOcrImageRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomOcrImage token error:%v", err)
		return nil, err
	}
	return a.svc.CustomOcrImage(ctx, req)
}

func (a AuthorizedService) CustomGetGroupSystemMsg(ctx context.Context, req *entity.CustomGetGroupSystemMsgReq) (*entity.CustomGetGroupSystemMsgRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetGroupSystemMsg token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetGroupSystemMsg(ctx, req)
}

func (a AuthorizedService) CustomUploadGroupFile(ctx context.Context, req *entity.CustomUploadGroupFileReq) (*entity.CustomUploadGroupFileRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomUploadGroupFile token error:%v", err)
		return nil, err
	}
	return a.svc.CustomUploadGroupFile(ctx, req)
}

func (a AuthorizedService) CustomGetGroupFileSystemInfo(ctx context.Context, req *entity.CustomGetGroupFileSystemInfoReq) (*entity.CustomGetGroupFileSystemInfoRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetGroupFileSystemInfo token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetGroupFileSystemInfo(ctx, req)
}

func (a AuthorizedService) CustomGetGroupRootFiles(ctx context.Context, req *entity.CustomGetGroupRootFilesReq) (*entity.CustomGetGroupRootFilesRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetGroupRootFiles token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetGroupRootFiles(ctx, req)
}

func (a AuthorizedService) CustomGetGroupFilesByFolder(ctx context.Context, req *entity.CustomGetGroupFilesByFolderReq) (*entity.CustomGetGroupFilesByFolderRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetGroupFilesByFolder token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetGroupFilesByFolder(ctx, req)
}

func (a AuthorizedService) CustomGetGroupFileUrl(ctx context.Context, req *entity.CustomGetGroupFileUrlReq) (*entity.CustomGetGroupFileUrlRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetGroupFileUrl token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetGroupFileUrl(ctx, req)
}

func (a AuthorizedService) CustomGetStatus(ctx context.Context, req *entity.CustomGetStatusReq) (*entity.CustomGetStatusRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetStatus token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetStatus(ctx, req)
}

func (a AuthorizedService) CustomGetGroupAtAllRemain(ctx context.Context, req *entity.CustomGetGroupAtAllRemainReq) (*entity.CustomGetGroupAtAllRemainRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetGroupAtAllRemain token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetGroupAtAllRemain(ctx, req)
}

func (a AuthorizedService) CustomGetVipInfo(ctx context.Context, req *entity.CustomGetVipInfoReq) (*entity.CustomGetVipInfoRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetVipInfo token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetVipInfo(ctx, req)
}

func (a AuthorizedService) CustomSendGroupNotice(ctx context.Context, req *entity.CustomSendGroupNoticeReq) (*entity.CustomSendGroupNoticeRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomSendGroupNotice token error:%v", err)
		return nil, err
	}
	return a.svc.CustomSendGroupNotice(ctx, req)
}

func (a AuthorizedService) CustomReloadEventFilter(ctx context.Context, req *entity.CustomReloadEventFilterReq) (*entity.CustomReloadEventFilterRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomReloadEventFilter token error:%v", err)
		return nil, err
	}
	return a.svc.CustomReloadEventFilter(ctx, req)
}

func (a AuthorizedService) CustomDownloadFile(ctx context.Context, req *entity.CustomDownloadFileReq) (*entity.CustomDownloadFileRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomDownloadFile token error:%v", err)
		return nil, err
	}
	return a.svc.CustomDownloadFile(ctx, req)
}

func (a AuthorizedService) CustomGetOnlineClinets(ctx context.Context, req *entity.CustomGetOnlineClientsReq) (*entity.CustomGetOnlineClientsRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetOnlineClinets token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetOnlineClinets(ctx, req)
}

func (a AuthorizedService) CustomGetGroupMsgHistory(ctx context.Context, req *entity.CustomGetGroupMsgHistoryReq) (*entity.CustomGetGroupMsgHistoryRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetGroupMsgHistory token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetGroupMsgHistory(ctx, req)
}

func (a AuthorizedService) CustomSetEssenceMsg(ctx context.Context, req *entity.CustomSetEssenceMsgReq) (*entity.CustomSetEssenceMsgRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomSetEssenceMsg token error:%v", err)
		return nil, err
	}
	return a.svc.CustomSetEssenceMsg(ctx, req)
}

func (a AuthorizedService) CustomDeleteEssenceMsg(ctx context.Context, req *entity.CustomDeleteEssenceMsgReq) (*entity.CustomDeleteEssenceMsgRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomDeleteEssenceMsg token error:%v", err)
		return nil, err
	}
	return a.svc.CustomDeleteEssenceMsg(ctx, req)
}

func (a AuthorizedService) CustomGetEssenceMsgList(ctx context.Context, req *entity.CustomGetEssenceMsgListReq) (*entity.CustomGetEssenceMsgListRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetEssenceMsgList token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetEssenceMsgList(ctx, req)
}

func (a AuthorizedService) CustomCheckUrlSafely(ctx context.Context, req *entity.CustomCheckUrlSafelyReq) (*entity.CustomCheckUrlSafelyRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomCheckUrlSafely token error:%v", err)
		return nil, err
	}
	return a.svc.CustomCheckUrlSafely(ctx, req)
}

func (a AuthorizedService) CustomGetModelShow(ctx context.Context, req *entity.CustomGetModelShowReq) (*entity.CustomGetModelShowRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetModelShow token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetModelShow(ctx, req)
}

func (a AuthorizedService) CustomSetModelShow(ctx context.Context, req *entity.CustomSetModelShowReq) (*entity.CustomSetModelShowRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomSetModelShow token error:%v", err)
		return nil, err
	}
	return a.svc.CustomSetModelShow(ctx, req)
}

func (a AuthorizedService) CustomGetMsg(ctx context.Context, req *entity.CustomGetMsgReq) (*entity.CustomGetMsgRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetMsg token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetMsg(ctx, req)
}

func (a AuthorizedService) CustomGetImage(ctx context.Context, req *entity.CustomGetImageReq) (*entity.CustomGetImageRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetImage token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetImage(ctx, req)
}

func (a AuthorizedService) CustomSendGroupForwardMsg(ctx context.Context, req *entity.CustomSendGroupForwardMsgReq) (*entity.CustomSendGroupForwardMsgRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomSendGroupFordMsg token error:%v", err)
		return nil, err
	}
	return a.svc.CustomSendGroupForwardMsg(ctx, req)
}

func (a AuthorizedService) CustomGetForwardMsg(ctx context.Context, req *entity.CustomGetForwardMsgReq) (*entity.CustomGetForwardMsgRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CustomGetForwardMsg token error:%v", err)
		return nil, err
	}
	return a.svc.CustomGetForwardMsg(ctx, req)
}

func (a AuthorizedService) GetConfig(ctx context.Context, req *entity.Config) (*entity.Config, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetConfig token error:%v", err)
		return nil, err
	}
	return a.svc.GetConfig(ctx, req)
}

func (a AuthorizedService) UpdateConfig(ctx context.Context, req *entity.Config) (*entity.Config, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("UpdateConfig token error:%v", err)
		return nil, err
	}
	return a.svc.UpdateConfig(ctx, req)
}

func (a AuthorizedService) SetBotAdapterKill(ctx context.Context, req *entity.SetRestartReq) (*entity.SetRestartRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SetBotAdapterKill token error:%v", err)
		return nil, err
	}
	return a.svc.SetBotAdapterKill(ctx, req)
}

func (a AuthorizedService) GetAuthToken(ctx context.Context, req *entity.GetAuthTokenReq) (*entity.GetAuthTokenRsp, error) {
	token, err := a.ausvc.GetToken(req.GetAppId(), req.GetAppSecret())
	if err != nil {
		log.Errorf("GetAuthToken token error:%v", err)
		return nil, err
	}
	return &entity.GetAuthTokenRsp{
		Token: token,
	}, nil
}

func (a AuthorizedService) SendPrivateMsg(ctx context.Context, req *entity.SendPrivateMsgReq) (*entity.SendMsgRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SendPrivateMsg token error:%v", err)
		return nil, err
	}
	return a.svc.SendPrivateMsg(ctx, req)
}

func (a AuthorizedService) SendGroupMsg(ctx context.Context, req *entity.SendGroupMsgReq) (*entity.SendMsgRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SendGroupMsg token error:%v", err)
		return nil, err
	}
	return a.svc.SendGroupMsg(ctx, req)
}

func (a AuthorizedService) SendMsg(ctx context.Context, req *entity.SendMsgReq) (*entity.SendMsgRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SendMsg token error:%v", err)
		return nil, err
	}
	return a.svc.SendMsg(ctx, req)
}

func (a AuthorizedService) DeleteMsg(ctx context.Context, req *entity.DeleteMsgReq) (*entity.DeleteMsgRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("DeleteMsg token error:%v", err)
		return nil, err
	}
	return a.svc.DeleteMsg(ctx, req)
}

func (a AuthorizedService) GetMsg(ctx context.Context, req *entity.GetMsgReq) (*entity.GetMsgRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetMsg token error:%v", err)
		return nil, err
	}
	return a.svc.GetMsg(ctx, req)
}

func (a AuthorizedService) GetForwardMsg(ctx context.Context, req *entity.GetForwardMsgReq) (*entity.GetForwardMsgRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetForwardMsg token error:%v", err)
		return nil, err
	}
	return a.svc.GetForwardMsg(ctx, req)
}

func (a AuthorizedService) SendLike(ctx context.Context, req *entity.SendLikeReq) (*entity.SendLikeRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SendLike token error:%v", err)
		return nil, err
	}
	return a.svc.SendLike(ctx, req)
}

func (a AuthorizedService) SetGroupKick(ctx context.Context, req *entity.SetGroupKickReq) (*entity.SetGroupKickRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SetGroupKick token error:%v", err)
		return nil, err
	}
	return a.svc.SetGroupKick(ctx, req)
}

func (a AuthorizedService) SetGroupBan(ctx context.Context, req *entity.SetGroupBanReq) (*entity.SetGroupBanRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SetGroupBan token error:%v", err)
		return nil, err
	}
	return a.svc.SetGroupBan(ctx, req)
}

func (a AuthorizedService) SetGroupAnonymousBan(ctx context.Context, req *entity.SetGroupAnonymousBanReq) (*entity.SetGroupAnonymousBanRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SetGroupAnonymousBan token error:%v", err)
		return nil, err
	}
	return a.svc.SetGroupAnonymousBan(ctx, req)
}

func (a AuthorizedService) SetGroupWholeBan(ctx context.Context, req *entity.SetGroupWholeBanReq) (*entity.SetGroupWholeBanRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SetGroupWholeBan token error:%v", err)
		return nil, err
	}
	return a.svc.SetGroupWholeBan(ctx, req)
}

func (a AuthorizedService) SetGroupAdmin(ctx context.Context, req *entity.SetGroupAdminReq) (*entity.SetGroupAdminRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SetGroupAdmin token error:%v", err)
		return nil, err
	}
	return a.svc.SetGroupAdmin(ctx, req)
}

func (a AuthorizedService) SetGroupAnonymous(ctx context.Context, req *entity.SetGroupAnonymousReq) (*entity.SetGroupAnonymousRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SetGroupAnonymous token error:%v", err)
		return nil, err
	}
	return a.svc.SetGroupAnonymous(ctx, req)
}

func (a AuthorizedService) SetGroupCard(ctx context.Context, req *entity.SetGroupCardReq) (*entity.SetGroupCardRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SetGroupCard token error:%v", err)
		return nil, err
	}
	return a.svc.SetGroupCard(ctx, req)
}

func (a AuthorizedService) SetGroupName(ctx context.Context, req *entity.SetGroupNameReq) (*entity.SetGroupNameRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SetGroupName token error:%v", err)
		return nil, err
	}
	return a.svc.SetGroupName(ctx, req)
}

func (a AuthorizedService) SetGroupLeave(ctx context.Context, req *entity.SetGroupLeaveReq) (*entity.SetGroupLeaveRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SetGroupLeave token error:%v", err)
		return nil, err
	}
	return a.svc.SetGroupLeave(ctx, req)
}

func (a AuthorizedService) SetGroupSpecialTitle(ctx context.Context, req *entity.SetGroupSpecialTitleReq) (*entity.SetGroupSpecialTitleRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SetGroupSpecialTitle token error:%v", err)
		return nil, err
	}
	return a.svc.SetGroupSpecialTitle(ctx, req)
}

func (a AuthorizedService) SetGroupAddRequest(ctx context.Context, req *entity.SetGroupAddRequestReq) (*entity.SetGroupAddRequestRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SetGroupAddRequest token error:%v", err)
		return nil, err
	}
	return a.svc.SetGroupAddRequest(ctx, req)
}

func (a AuthorizedService) SetFriendAddRequest(ctx context.Context, req *entity.SetFriendAddRequestReq) (*entity.SetFriendaddRequestRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SetFriendAddRequest token error:%v", err)
		return nil, err
	}
	return a.svc.SetFriendAddRequest(ctx, req)
}

func (a AuthorizedService) GetLoginInfo(ctx context.Context, req *entity.GetLoginInfoReq) (*entity.GetLoginInfoRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetLoginInfo token error:%v", err)
		return nil, err
	}
	return a.svc.GetLoginInfo(ctx, req)
}

func (a AuthorizedService) GetStrangerInfo(ctx context.Context, req *entity.GetStrangerInfoReq) (*entity.GetStrangerInfoRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetStrangerInfo token error:%v", err)
		return nil, err
	}
	return a.svc.GetStrangerInfo(ctx, req)
}

func (a AuthorizedService) GetFriendList(ctx context.Context, req *entity.GetFriendListReq) (*entity.GetFriendListRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetFriendList token error:%v", err)
		return nil, err
	}
	return a.svc.GetFriendList(ctx, req)
}

func (a AuthorizedService) GetGroupInfo(ctx context.Context, req *entity.GetGroupInfoReq) (*entity.GetGroupInfoRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetGroupInfo token error:%v", err)
		return nil, err
	}
	return a.svc.GetGroupInfo(ctx, req)
}

func (a AuthorizedService) GetGroupList(ctx context.Context, req *entity.GetGroupListReq) (*entity.GetGroupListRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetGroupList token error:%v", err)
		return nil, err
	}
	return a.svc.GetGroupList(ctx, req)
}

func (a AuthorizedService) GetGroupMemberInfo(ctx context.Context, req *entity.GetGroupMemberInfoReq) (*entity.GetGroupMemberInfoRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetGroupMemberInfo token error:%v", err)
		return nil, err
	}
	return a.svc.GetGroupMemberInfo(ctx, req)
}

func (a AuthorizedService) GetGroupMemberList(ctx context.Context, req *entity.GetGroupMemberListReq) (*entity.GetGroupMemberListRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetGroupMemberList token error:%v", err)
		return nil, err
	}
	return a.svc.GetGroupMemberList(ctx, req)
}

func (a AuthorizedService) GetGroupHonorInfo(ctx context.Context, req *entity.GetGroupHonorInfoReq) (*entity.GetGroupHonorInfoRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetGroupHonorInfo token error:%v", err)
		return nil, err
	}
	return a.svc.GetGroupHonorInfo(ctx, req)
}

func (a AuthorizedService) GetCookies(ctx context.Context, req *entity.GetCookiesReq) (*entity.GetCookiesRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetCookies token error:%v", err)
		return nil, err
	}
	return a.svc.GetCookies(ctx, req)
}

func (a AuthorizedService) GetCsrfToken(ctx context.Context, req *entity.GetCsrfTokenReq) (*entity.GetCsrfTokenRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetCsrfToken token error:%v", err)
		return nil, err
	}
	return a.svc.GetCsrfToken(ctx, req)
}

func (a AuthorizedService) GetCredentials(ctx context.Context, req *entity.GetCredentialsReq) (*entity.GetCredentialsRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetCredentials token error:%v", err)
		return nil, err
	}
	return a.svc.GetCredentials(ctx, req)
}

func (a AuthorizedService) GetRecord(ctx context.Context, req *entity.GetRecordReq) (*entity.GetRecordRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetRecord token error:%v", err)
		return nil, err
	}
	return a.svc.GetRecord(ctx, req)
}

func (a AuthorizedService) GetImage(ctx context.Context, req *entity.GetImageReq) (*entity.GetImageRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetImage token error:%v", err)
		return nil, err
	}
	return a.svc.GetImage(ctx, req)
}

func (a AuthorizedService) CanSendImage(ctx context.Context, req *entity.CanSendImageReq) (*entity.CanSendImageRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CanSendImage token error:%v", err)
		return nil, err
	}
	return a.svc.CanSendImage(ctx, req)
}

func (a AuthorizedService) CanSendRecord(ctx context.Context, req *entity.CanSendRecordReq) (*entity.CanSendRecordRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CanSendRecord token error:%v", err)
		return nil, err
	}
	return a.svc.CanSendRecord(ctx, req)
}

func (a AuthorizedService) GetStatus(ctx context.Context, req *entity.GetStatusReq) (*entity.GetStatusRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetStatus token error:%v", err)
		return nil, err
	}
	return a.svc.GetStatus(ctx, req)
}

func (a AuthorizedService) GetVersionInfo(ctx context.Context, req *entity.GetVersionInfoReq) (*entity.GetVersionInfoRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("GetVersionInfo token error:%v", err)
		return nil, err
	}
	return a.svc.GetVersionInfo(ctx, req)
}

func (a AuthorizedService) SetRestart(ctx context.Context, req *entity.SetRestartReq) (*entity.SetRestartRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("SetRestart token error:%v", err)
		return nil, err
	}
	return a.svc.SetRestart(ctx, req)
}

func (a AuthorizedService) CleanCache(ctx context.Context, req *entity.CleanCacheReq) (*entity.CleanCacheRsp, error) {
	err := a.ausvc.CheckFormGrpcContext(ctx)
	if err != nil {
		log.Errorf("CleanCache token error:%v", err)
		return nil, err
	}
	return a.svc.CleanCache(ctx, req)
}

// NewAuthorizedService 初始化 token校验
func NewAuthorizedService(ct *dig.Container, svc service.AdapterServiceServer) (*AuthorizedService, error) {
	var a AuthorizedService
	a.svc = svc
	_ = ct.Invoke(func(opt *config.Config) {
		a.conf = opt
	})
	_ = ct.Invoke(func(au *auth.AuthServer) {
		a.ausvc = au
	})
	return &a, nil
}
