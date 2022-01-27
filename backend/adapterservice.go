package backend

import (
	"context"
	"github.com/scjtqs2/bot_adapter/adapter"
	"github.com/scjtqs2/bot_adapter/config"
	"github.com/scjtqs2/bot_adapter/handler"
	"github.com/scjtqs2/bot_adapter/pb/entity"
	"go.uber.org/dig"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

// AdapterService 真实的业务处理
type AdapterService struct {
	adapterHander *handler.AdapterHandler
	adapter       adapter.AdapterInterface
}

func (a AdapterService) CustomSetGroupPortrait(ctx context.Context, req *entity.CustomSetGroupPortraitReq) (*entity.CustomSetGroupPortraitRsp, error) {
	return a.adapter.CustomSetGroupPortrait(req)
}

func (a AdapterService) CustomGetWordSlices(ctx context.Context, req *entity.CustomGetWordSlicesReq) (*entity.CustomGetWordSlicesRsp, error) {
	return a.adapter.CustomGetWordSlices(req)
}

func (a AdapterService) CustomOcrImage(ctx context.Context, req *entity.CustomOcrImageReq) (*entity.CustomOcrImageRsp, error) {
	return a.adapter.CustomOcrImage(req)
}

func (a AdapterService) CustomGetGroupSystemMsg(ctx context.Context, req *entity.CustomGetGroupSystemMsgReq) (*entity.CustomGetGroupSystemMsgRsp, error) {
	return a.adapter.CustomGetGroupSystemMsg(req)
}

func (a AdapterService) CustomUploadGroupFile(ctx context.Context, req *entity.CustomUploadGroupFileReq) (*entity.CustomUploadGroupFileRsp, error) {
	return a.adapter.CustomUploadGroupFile(req)
}

func (a AdapterService) CustomGetGroupFileSystemInfo(ctx context.Context, req *entity.CustomGetGroupFileSystemInfoReq) (*entity.CustomGetGroupFileSystemInfoRsp, error) {
	return a.adapter.CustomGetGroupFileSystemInfo(req)
}

func (a AdapterService) CustomGetGroupRootFiles(ctx context.Context, req *entity.CustomGetGroupRootFilesReq) (*entity.CustomGetGroupRootFilesRsp, error) {
	return a.adapter.CustomGetGroupRootFiles(req)
}

func (a AdapterService) CustomGetGroupFilesByFolder(ctx context.Context, req *entity.CustomGetGroupFilesByFolderReq) (*entity.CustomGetGroupFilesByFolderRsp, error) {
	return a.adapter.CustomGetGroupFilesByFolder(req)
}

func (a AdapterService) CustomGetGroupFileUrl(ctx context.Context, req *entity.CustomGetGroupFileUrlReq) (*entity.CustomGetGroupFileUrlRsp, error) {
	return a.adapter.CustomGetGroupFileUrl(req)
}

func (a AdapterService) CustomGetStatus(ctx context.Context, req *entity.CustomGetStatusReq) (*entity.CustomGetStatusRsp, error) {
	return a.adapter.CustomGetStatus(req)
}

func (a AdapterService) CustomGetGroupAtAllRemain(ctx context.Context, req *entity.CustomGetGroupAtAllRemainReq) (*entity.CustomGetGroupAtAllRemainRsp, error) {
	return a.adapter.CustomGetGroupAtAllRemain(req)
}

func (a AdapterService) CustomGetVipInfo(ctx context.Context, req *entity.CustomGetVipInfoReq) (*entity.CustomGetVipInfoRsp, error) {
	return a.adapter.CustomGetVipInfo(req)
}

func (a AdapterService) CustomSendGroupNotice(ctx context.Context, req *entity.CustomSendGroupNoticeReq) (*entity.CustomSendGroupNoticeRsp, error) {
	return a.adapter.CustomSendGroupNotice(req)
}

func (a AdapterService) CustomReloadEventFilter(ctx context.Context, req *entity.CustomReloadEventFilterReq) (*entity.CustomReloadEventFilterRsp, error) {
	return a.adapter.CustomReloadEventFilter(req)
}

func (a AdapterService) CustomDownloadFile(ctx context.Context, req *entity.CustomDownloadFileReq) (*entity.CustomDownloadFileRsp, error) {
	return a.adapter.CustomDownloadFile(req)
}

func (a AdapterService) CustomGetOnlineClinets(ctx context.Context, req *entity.CustomGetOnlineClientsReq) (*entity.CustomGetOnlineClientsRsp, error) {
	return a.adapter.CustomGetOnlineClinets(req)
}

func (a AdapterService) CustomGetGroupMsgHistory(ctx context.Context, req *entity.CustomGetGroupMsgHistoryReq) (*entity.CustomGetGroupMsgHistoryRsp, error) {
	return a.adapter.CustomGetGroupMsgHistory(req)
}

func (a AdapterService) CustomSetEssenceMsg(ctx context.Context, req *entity.CustomSetEssenceMsgReq) (*entity.CustomSetEssenceMsgRsp, error) {
	return a.adapter.CustomSetEssenceMsg(req)
}

func (a AdapterService) CustomDeleteEssenceMsg(ctx context.Context, req *entity.CustomDeleteEssenceMsgReq) (*entity.CustomDeleteEssenceMsgRsp, error) {
	return a.adapter.CustomDeleteEssenceMsg(req)
}

func (a AdapterService) CustomGetEssenceMsgList(ctx context.Context, req *entity.CustomGetEssenceMsgListReq) (*entity.CustomGetEssenceMsgListRsp, error) {
	return a.adapter.CustomGetEssenceMsgList(req)
}

func (a AdapterService) CustomCheckUrlSafely(ctx context.Context, req *entity.CustomCheckUrlSafelyReq) (*entity.CustomCheckUrlSafelyRsp, error) {
	return a.adapter.CustomCheckUrlSafely(req)
}

func (a AdapterService) CustomGetModelShow(ctx context.Context, req *entity.CustomGetModelShowReq) (*entity.CustomGetModelShowRsp, error) {
	return a.adapter.CustomGetModelShow(req)
}

func (a AdapterService) CustomSetModelShow(ctx context.Context, req *entity.CustomSetModelShowReq) (*entity.CustomSetModelShowRsp, error) {
	return a.adapter.CustomSetModelShow(req)
}

func (a AdapterService) CustomGetMsg(ctx context.Context, req *entity.CustomGetMsgReq) (*entity.CustomGetMsgRsp, error) {
	return a.adapter.CustomGetMsg(req)
}

func (a AdapterService) CustomGetImage(ctx context.Context, req *entity.CustomGetImageReq) (*entity.CustomGetImageRsp, error) {
	return a.adapter.CustomGetImage(req)
}

func (a AdapterService) CustomSendGroupForwardMsg(ctx context.Context, req *entity.CustomSendGroupForwardMsgReq) (*entity.CustomSendGroupForwardMsgRsp, error) {
	return a.adapter.CustomSendGroupForwardMsg(req)
}

func (a AdapterService) CustomGetForwardMsg(ctx context.Context, req *entity.CustomGetForwardMsgReq) (*entity.CustomGetForwardMsgRsp, error) {
	return a.adapter.CustomGetForwardMsg(req)
}

func (a AdapterService) GetConfig(ctx context.Context, req *entity.Config) (*entity.Config, error) {
	cfg := config.Parse()
	c, err := yaml.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	return &entity.Config{
		Config: c,
	}, nil
}

func (a AdapterService) UpdateConfig(ctx context.Context, req *entity.Config) (*entity.Config, error) {
	var cfg config.Config
	err := yaml.Unmarshal(req.Config, &cfg)
	if err != nil {
		return nil, err
	}
	err = cfg.Save()
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (a AdapterService) SetBotAdapterKill(ctx context.Context, req *entity.SetRestartReq) (*entity.SetRestartRsp, error) {
	go func() {
		if req.GetDelay() == 0 {
			req.Delay = 2
		}
		time.Sleep(time.Duration(req.GetDelay()) * time.Second)
		os.Exit(1)
	}()
	return nil, nil
}

func (a AdapterService) GetAuthToken(ctx context.Context, req *entity.GetAuthTokenReq) (*entity.GetAuthTokenRsp, error) {
	panic("implement me")
}

func (a AdapterService) SendPrivateMsg(ctx context.Context, req *entity.SendPrivateMsgReq) (*entity.SendMsgRsp, error) {
	return a.adapter.SendPrivateMsg(req)
}

func (a AdapterService) SendGroupMsg(ctx context.Context, req *entity.SendGroupMsgReq) (*entity.SendMsgRsp, error) {
	return a.adapter.SendGroupMsg(req)
}

func (a AdapterService) SendMsg(ctx context.Context, req *entity.SendMsgReq) (*entity.SendMsgRsp, error) {
	return a.adapter.SendMsg(req)
}

func (a AdapterService) DeleteMsg(ctx context.Context, req *entity.DeleteMsgReq) (*entity.DeleteMsgRsp, error) {
	return a.adapter.DeleteMsg(req)
}

func (a AdapterService) GetMsg(ctx context.Context, req *entity.GetMsgReq) (*entity.GetMsgRsp, error) {
	return a.adapter.GetMsg(req)
}

func (a AdapterService) GetForwardMsg(ctx context.Context, req *entity.GetForwardMsgReq) (*entity.GetForwardMsgRsp, error) {
	return a.adapter.GetForwardMsg(req)
}

func (a AdapterService) SendLike(ctx context.Context, req *entity.SendLikeReq) (*entity.SendLikeRsp, error) {
	return a.adapter.SendLike(req)
}

func (a AdapterService) SetGroupKick(ctx context.Context, req *entity.SetGroupKickReq) (*entity.SetGroupKickRsp, error) {
	return a.adapter.SetGroupKick(req)
}

func (a AdapterService) SetGroupBan(ctx context.Context, req *entity.SetGroupBanReq) (*entity.SetGroupBanRsp, error) {
	return a.adapter.SetGroupBan(req)
}

func (a AdapterService) SetGroupAnonymousBan(ctx context.Context, req *entity.SetGroupAnonymousBanReq) (*entity.SetGroupAnonymousBanRsp, error) {
	return a.adapter.SetGroupAnonymousBan(req)
}

func (a AdapterService) SetGroupWholeBan(ctx context.Context, req *entity.SetGroupWholeBanReq) (*entity.SetGroupWholeBanRsp, error) {
	return a.adapter.SetGroupWholeBan(req)
}

func (a AdapterService) SetGroupAdmin(ctx context.Context, req *entity.SetGroupAdminReq) (*entity.SetGroupAdminRsp, error) {
	return a.adapter.SetGroupAdmin(req)
}

func (a AdapterService) SetGroupAnonymous(ctx context.Context, req *entity.SetGroupAnonymousReq) (*entity.SetGroupAnonymousRsp, error) {
	return a.adapter.SetGroupAnonymous(req)
}

func (a AdapterService) SetGroupCard(ctx context.Context, req *entity.SetGroupCardReq) (*entity.SetGroupCardRsp, error) {
	return a.adapter.SetGroupCard(req)
}

func (a AdapterService) SetGroupName(ctx context.Context, req *entity.SetGroupNameReq) (*entity.SetGroupNameRsp, error) {
	return a.adapter.SetGroupName(req)
}

func (a AdapterService) SetGroupLeave(ctx context.Context, req *entity.SetGroupLeaveReq) (*entity.SetGroupLeaveRsp, error) {
	return a.adapter.SetGroupLeave(req)
}

func (a AdapterService) SetGroupSpecialTitle(ctx context.Context, req *entity.SetGroupSpecialTitleReq) (*entity.SetGroupSpecialTitleRsp, error) {
	return a.adapter.SetGroupSpecialTitle(req)
}

func (a AdapterService) SetGroupAddRequest(ctx context.Context, req *entity.SetGroupAddRequestReq) (*entity.SetGroupAddRequestRsp, error) {
	return a.adapter.SetGroupAddRequest(req)
}

func (a AdapterService) SetFriendAddRequest(ctx context.Context, req *entity.SetFriendAddRequestReq) (*entity.SetFriendaddRequestRsp, error) {
	return a.adapter.SetFriendAddRequest(req)
}

func (a AdapterService) GetLoginInfo(ctx context.Context, req *entity.GetLoginInfoReq) (*entity.GetLoginInfoRsp, error) {
	return a.adapter.GetLoginInfo(req)
}

func (a AdapterService) GetStrangerInfo(ctx context.Context, req *entity.GetStrangerInfoReq) (*entity.GetStrangerInfoRsp, error) {
	return a.adapter.GetStrangerInfo(req)
}

func (a AdapterService) GetFriendList(ctx context.Context, req *entity.GetFriendListReq) (*entity.GetFriendListRsp, error) {
	return a.adapter.GetFriendList(req)
}

func (a AdapterService) GetGroupInfo(ctx context.Context, req *entity.GetGroupInfoReq) (*entity.GetGroupInfoRsp, error) {
	return a.adapter.GetGroupInfo(req)
}

func (a AdapterService) GetGroupList(ctx context.Context, req *entity.GetGroupListReq) (*entity.GetGroupListRsp, error) {
	return a.adapter.GetGroupList(req)
}

func (a AdapterService) GetGroupMemberInfo(ctx context.Context, req *entity.GetGroupMemberInfoReq) (*entity.GetGroupMemberInfoRsp, error) {
	return a.adapter.GetGroupMemberInfo(req)
}

func (a AdapterService) GetGroupMemberList(ctx context.Context, req *entity.GetGroupMemberListReq) (*entity.GetGroupMemberListRsp, error) {
	return a.adapter.GetGroupMemberList(req)
}

func (a AdapterService) GetGroupHonorInfo(ctx context.Context, req *entity.GetGroupHonorInfoReq) (*entity.GetGroupHonorInfoRsp, error) {
	return a.adapter.GetGroupHonorInfo(req)
}

func (a AdapterService) GetCookies(ctx context.Context, req *entity.GetCookiesReq) (*entity.GetCookiesRsp, error) {
	return a.adapter.GetCookies(req)
}

func (a AdapterService) GetCsrfToken(ctx context.Context, req *entity.GetCsrfTokenReq) (*entity.GetCsrfTokenRsp, error) {
	return a.adapter.GetCsrfToken(req)
}

func (a AdapterService) GetCredentials(ctx context.Context, req *entity.GetCredentialsReq) (*entity.GetCredentialsRsp, error) {
	return a.adapter.GetCredentials(req)
}

func (a AdapterService) GetRecord(ctx context.Context, req *entity.GetRecordReq) (*entity.GetRecordRsp, error) {
	return a.adapter.GetRecord(req)
}

func (a AdapterService) GetImage(ctx context.Context, req *entity.GetImageReq) (*entity.GetImageRsp, error) {
	return a.adapter.GetImage(req)
}

func (a AdapterService) CanSendImage(ctx context.Context, req *entity.CanSendImageReq) (*entity.CanSendImageRsp, error) {
	return a.adapter.CanSendImage(req)
}

func (a AdapterService) CanSendRecord(ctx context.Context, req *entity.CanSendRecordReq) (*entity.CanSendRecordRsp, error) {
	return a.adapter.CanSendRecord(req)
}

func (a AdapterService) GetStatus(ctx context.Context, req *entity.GetStatusReq) (*entity.GetStatusRsp, error) {
	return a.adapter.GetStatus(req)
}

func (a AdapterService) GetVersionInfo(ctx context.Context, req *entity.GetVersionInfoReq) (*entity.GetVersionInfoRsp, error) {
	return a.adapter.GetVersionInfo(req)
}

func (a AdapterService) SetRestart(ctx context.Context, req *entity.SetRestartReq) (*entity.SetRestartRsp, error) {
	return a.adapter.SetRestart(req)
}

func (a AdapterService) CleanCache(ctx context.Context, req *entity.CleanCacheReq) (*entity.CleanCacheRsp, error) {
	return a.adapter.CleanCache(req)
}

// NewAdapterService 初始化
func NewAdapterService(ct *dig.Container) (*AdapterService, error) {
	var adapt AdapterService
	adapt.adapterHander, _ = handler.NewAdapterHandler(ct)
	_ = ct.Invoke(func(a adapter.AdapterInterface) {
		adapt.adapter = a
	})
	return &adapt, nil
}
