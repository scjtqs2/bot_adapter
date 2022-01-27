package backend

import (
	"context"
	"fmt"
	"github.com/scjtqs2/bot_adapter/auth"
	"github.com/scjtqs2/bot_adapter/config"
	"github.com/scjtqs2/bot_adapter/pb/entity"
	"github.com/scjtqs2/bot_adapter/pb/service"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// PermissionService 权限验证处理 暂时不处理，后面有空再做
type PermissionService struct {
	svc  service.AdapterServiceServer
	conf *config.Config
	auth *auth.AuthServer
}

func (p *PermissionService) CustomSetGroupPortrait(ctx context.Context, req *entity.CustomSetGroupPortraitReq) (*entity.CustomSetGroupPortraitRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_SET_GROUP_PROTRAIT); err != nil {
		return nil, err
	}
	return p.svc.CustomSetGroupPortrait(ctx, req)
}

func (p *PermissionService) CustomGetWordSlices(ctx context.Context, req *entity.CustomGetWordSlicesReq) (*entity.CustomGetWordSlicesRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_GET_WORD_SLICES); err != nil {
		return nil, err
	}
	return p.svc.CustomGetWordSlices(ctx, req)
}

func (p *PermissionService) CustomOcrImage(ctx context.Context, req *entity.CustomOcrImageReq) (*entity.CustomOcrImageRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_OCR_IMAGE); err != nil {
		return nil, err
	}
	return p.svc.CustomOcrImage(ctx, req)
}

func (p *PermissionService) CustomGetGroupSystemMsg(ctx context.Context, req *entity.CustomGetGroupSystemMsgReq) (*entity.CustomGetGroupSystemMsgRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_GET_GROUP_SYSTEM_MSG); err != nil {
		return nil, err
	}
	return p.svc.CustomGetGroupSystemMsg(ctx, req)
}

func (p *PermissionService) CustomUploadGroupFile(ctx context.Context, req *entity.CustomUploadGroupFileReq) (*entity.CustomUploadGroupFileRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_UPLOAD_GROUP_FILE); err != nil {
		return nil, err
	}
	return p.svc.CustomUploadGroupFile(ctx, req)
}

func (p *PermissionService) CustomGetGroupFileSystemInfo(ctx context.Context, req *entity.CustomGetGroupFileSystemInfoReq) (*entity.CustomGetGroupFileSystemInfoRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_GET_GROUP_FILE_SYSTEM_INFO); err != nil {
		return nil, err
	}
	return p.svc.CustomGetGroupFileSystemInfo(ctx, req)
}

func (p *PermissionService) CustomGetGroupRootFiles(ctx context.Context, req *entity.CustomGetGroupRootFilesReq) (*entity.CustomGetGroupRootFilesRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_GET_GROUP_ROOT_FILES); err != nil {
		return nil, err
	}
	return p.svc.CustomGetGroupRootFiles(ctx, req)
}

func (p *PermissionService) CustomGetGroupFilesByFolder(ctx context.Context, req *entity.CustomGetGroupFilesByFolderReq) (*entity.CustomGetGroupFilesByFolderRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_GET_GROUP_FILES_BY_FOLDER); err != nil {
		return nil, err
	}
	return p.svc.CustomGetGroupFilesByFolder(ctx, req)
}

func (p *PermissionService) CustomGetGroupFileUrl(ctx context.Context, req *entity.CustomGetGroupFileUrlReq) (*entity.CustomGetGroupFileUrlRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_GET_GROUP_FILE_URL); err != nil {
		return nil, err
	}
	return p.svc.CustomGetGroupFileUrl(ctx, req)
}

func (p *PermissionService) CustomGetStatus(ctx context.Context, req *entity.CustomGetStatusReq) (*entity.CustomGetStatusRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_STATUS); err != nil {
		return nil, err
	}
	return p.svc.CustomGetStatus(ctx, req)
}

func (p *PermissionService) CustomGetGroupAtAllRemain(ctx context.Context, req *entity.CustomGetGroupAtAllRemainReq) (*entity.CustomGetGroupAtAllRemainRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_GET_GROUP_AT_ALL_REMAIN); err != nil {
		return nil, err
	}
	return p.svc.CustomGetGroupAtAllRemain(ctx, req)
}

func (p *PermissionService) CustomGetVipInfo(ctx context.Context, req *entity.CustomGetVipInfoReq) (*entity.CustomGetVipInfoRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_GET_VIP_INFO); err != nil {
		return nil, err
	}
	return p.svc.CustomGetVipInfo(ctx, req)
}

func (p *PermissionService) CustomSendGroupNotice(ctx context.Context, req *entity.CustomSendGroupNoticeReq) (*entity.CustomSendGroupNoticeRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_SEND_GROUP_NOTICE); err != nil {
		return nil, err
	}
	return p.svc.CustomSendGroupNotice(ctx, req)
}

func (p *PermissionService) CustomReloadEventFilter(ctx context.Context, req *entity.CustomReloadEventFilterReq) (*entity.CustomReloadEventFilterRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_RELOAD_EVENT_FILTER); err != nil {
		return nil, err
	}
	return p.svc.CustomReloadEventFilter(ctx, req)
}

func (p *PermissionService) CustomDownloadFile(ctx context.Context, req *entity.CustomDownloadFileReq) (*entity.CustomDownloadFileRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_DOWNLOAD_FILE); err != nil {
		return nil, err
	}
	return p.svc.CustomDownloadFile(ctx, req)
}

func (p *PermissionService) CustomGetOnlineClinets(ctx context.Context, req *entity.CustomGetOnlineClientsReq) (*entity.CustomGetOnlineClientsRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_GET_ONLINE_CLIENTS); err != nil {
		return nil, err
	}
	return p.svc.CustomGetOnlineClinets(ctx, req)
}

func (p *PermissionService) CustomGetGroupMsgHistory(ctx context.Context, req *entity.CustomGetGroupMsgHistoryReq) (*entity.CustomGetGroupMsgHistoryRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_GET_GROUP_MSG_HISTORY); err != nil {
		return nil, err
	}
	return p.svc.CustomGetGroupMsgHistory(ctx, req)
}

func (p *PermissionService) CustomSetEssenceMsg(ctx context.Context, req *entity.CustomSetEssenceMsgReq) (*entity.CustomSetEssenceMsgRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_SET_ESSENCE_MSG); err != nil {
		return nil, err
	}
	return p.svc.CustomSetEssenceMsg(ctx, req)
}

func (p *PermissionService) CustomDeleteEssenceMsg(ctx context.Context, req *entity.CustomDeleteEssenceMsgReq) (*entity.CustomDeleteEssenceMsgRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_DELETE_ESSENCE_MSG); err != nil {
		return nil, err
	}
	return p.svc.CustomDeleteEssenceMsg(ctx, req)
}

func (p *PermissionService) CustomGetEssenceMsgList(ctx context.Context, req *entity.CustomGetEssenceMsgListReq) (*entity.CustomGetEssenceMsgListRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_GET_ESSENCE_MSG_LIST); err != nil {
		return nil, err
	}
	return p.svc.CustomGetEssenceMsgList(ctx, req)
}

func (p *PermissionService) CustomCheckUrlSafely(ctx context.Context, req *entity.CustomCheckUrlSafelyReq) (*entity.CustomCheckUrlSafelyRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_CHECK_URL_SAFELY); err != nil {
		return nil, err
	}
	return p.svc.CustomCheckUrlSafely(ctx, req)
}

func (p *PermissionService) CustomGetModelShow(ctx context.Context, req *entity.CustomGetModelShowReq) (*entity.CustomGetModelShowRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_GET_MODEL_SHOW); err != nil {
		return nil, err
	}
	return p.svc.CustomGetModelShow(ctx, req)
}

func (p *PermissionService) CustomSetModelShow(ctx context.Context, req *entity.CustomSetModelShowReq) (*entity.CustomSetModelShowRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_SET_MODEL_SHOW); err != nil {
		return nil, err
	}
	return p.svc.CustomSetModelShow(ctx, req)
}

func (p *PermissionService) CustomGetMsg(ctx context.Context, req *entity.CustomGetMsgReq) (*entity.CustomGetMsgRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_MSG); err != nil {
		return nil, err
	}
	return p.svc.CustomGetMsg(ctx, req)
}

func (p *PermissionService) CustomGetImage(ctx context.Context, req *entity.CustomGetImageReq) (*entity.CustomGetImageRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_IMAGE); err != nil {
		return nil, err
	}
	return p.svc.CustomGetImage(ctx, req)
}

func (p *PermissionService) CustomSendGroupForwardMsg(ctx context.Context, req *entity.CustomSendGroupForwardMsgReq) (*entity.CustomSendGroupForwardMsgRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CUSTOM_SEND_GROUP_FORWARD_msg); err != nil {
		return nil, err
	}
	return p.svc.CustomSendGroupForwardMsg(ctx, req)
}

func (p *PermissionService) CustomGetForwardMsg(ctx context.Context, req *entity.CustomGetForwardMsgReq) (*entity.CustomGetForwardMsgRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_FORWARD_MSG); err != nil {
		return nil, err
	}
	return p.svc.CustomGetForwardMsg(ctx, req)
}

func (p *PermissionService) GetConfig(ctx context.Context, e *entity.Config) (*entity.Config, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_CONFIG); err != nil {
		return nil, err
	}
	return p.svc.GetConfig(ctx, e)
}

func (p *PermissionService) UpdateConfig(ctx context.Context, e *entity.Config) (*entity.Config, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_UPDATE_CONFIG); err != nil {
		return nil, err
	}
	return p.svc.UpdateConfig(ctx, e)
}

func (p *PermissionService) SetBotAdapterKill(ctx context.Context, req *entity.SetRestartReq) (*entity.SetRestartRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SET_BOT_ADAPTER_KILL); err != nil {
		return nil, err
	}
	return p.svc.SetBotAdapterKill(ctx, req)
}

func (p *PermissionService) GetAuthToken(ctx context.Context, req *entity.GetAuthTokenReq) (*entity.GetAuthTokenRsp, error) {
	return p.svc.GetAuthToken(ctx, req)
}

func (p *PermissionService) SendPrivateMsg(ctx context.Context, req *entity.SendPrivateMsgReq) (*entity.SendMsgRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SEND_PRIVATE_MSG); err != nil {
		return nil, err
	}
	return p.svc.SendPrivateMsg(ctx, req)
}

func (p *PermissionService) SendGroupMsg(ctx context.Context, req *entity.SendGroupMsgReq) (*entity.SendMsgRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SEND_GROUP_MSG); err != nil {
		return nil, err
	}
	return p.svc.SendGroupMsg(ctx, req)
}

func (p *PermissionService) SendMsg(ctx context.Context, req *entity.SendMsgReq) (*entity.SendMsgRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SEND_MSG); err != nil {
		return nil, err
	}
	return p.svc.SendMsg(ctx, req)
}

func (p *PermissionService) DeleteMsg(ctx context.Context, req *entity.DeleteMsgReq) (*entity.DeleteMsgRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_DELETE_MSG); err != nil {
		return nil, err
	}
	return p.svc.DeleteMsg(ctx, req)
}

func (p *PermissionService) GetMsg(ctx context.Context, req *entity.GetMsgReq) (*entity.GetMsgRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_MSG); err != nil {
		return nil, err
	}
	return p.svc.GetMsg(ctx, req)
}

func (p *PermissionService) GetForwardMsg(ctx context.Context, req *entity.GetForwardMsgReq) (*entity.GetForwardMsgRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_FORWARD_MSG); err != nil {
		return nil, err
	}
	return p.svc.GetForwardMsg(ctx, req)
}

func (p *PermissionService) SendLike(ctx context.Context, req *entity.SendLikeReq) (*entity.SendLikeRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SEND_LIKE); err != nil {
		return nil, err
	}
	return p.svc.SendLike(ctx, req)
}

func (p *PermissionService) SetGroupKick(ctx context.Context, req *entity.SetGroupKickReq) (*entity.SetGroupKickRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SET_GROUP_KICK); err != nil {
		return nil, err
	}
	return p.svc.SetGroupKick(ctx, req)
}

func (p *PermissionService) SetGroupBan(ctx context.Context, req *entity.SetGroupBanReq) (*entity.SetGroupBanRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SET_GROUP_BAN); err != nil {
		return nil, err
	}
	return p.svc.SetGroupBan(ctx, req)
}

func (p *PermissionService) SetGroupAnonymousBan(ctx context.Context, req *entity.SetGroupAnonymousBanReq) (*entity.SetGroupAnonymousBanRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SET_GROUP_ANONYMOUS_BAN); err != nil {
		return nil, err
	}
	return p.svc.SetGroupAnonymousBan(ctx, req)
}

func (p *PermissionService) SetGroupWholeBan(ctx context.Context, req *entity.SetGroupWholeBanReq) (*entity.SetGroupWholeBanRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SET_GROUP_WHOLE_BAN); err != nil {
		return nil, err
	}
	return p.svc.SetGroupWholeBan(ctx, req)
}

func (p *PermissionService) SetGroupAdmin(ctx context.Context, req *entity.SetGroupAdminReq) (*entity.SetGroupAdminRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SET_GROUP_ADMIN); err != nil {
		return nil, err
	}
	return p.svc.SetGroupAdmin(ctx, req)
}

func (p *PermissionService) SetGroupAnonymous(ctx context.Context, req *entity.SetGroupAnonymousReq) (*entity.SetGroupAnonymousRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SET_GROUP_ANONYMOUS); err != nil {
		return nil, err
	}
	return p.svc.SetGroupAnonymous(ctx, req)
}

func (p *PermissionService) SetGroupCard(ctx context.Context, req *entity.SetGroupCardReq) (*entity.SetGroupCardRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SET_GROUP_CARD); err != nil {
		return nil, err
	}
	return p.svc.SetGroupCard(ctx, req)
}

func (p *PermissionService) SetGroupName(ctx context.Context, req *entity.SetGroupNameReq) (*entity.SetGroupNameRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SET_GROUP_NAME); err != nil {
		return nil, err
	}
	return p.svc.SetGroupName(ctx, req)
}

func (p *PermissionService) SetGroupLeave(ctx context.Context, req *entity.SetGroupLeaveReq) (*entity.SetGroupLeaveRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SET_GROUP_LEAVE); err != nil {
		return nil, err
	}
	return p.svc.SetGroupLeave(ctx, req)
}

func (p *PermissionService) SetGroupSpecialTitle(ctx context.Context, req *entity.SetGroupSpecialTitleReq) (*entity.SetGroupSpecialTitleRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SET_GROUP_SPECIAL_TITLE); err != nil {
		return nil, err
	}
	return p.svc.SetGroupSpecialTitle(ctx, req)
}

func (p *PermissionService) SetGroupAddRequest(ctx context.Context, req *entity.SetGroupAddRequestReq) (*entity.SetGroupAddRequestRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SET_GROUP_ADD_REQUEST); err != nil {
		return nil, err
	}
	return p.svc.SetGroupAddRequest(ctx, req)
}

func (p *PermissionService) SetFriendAddRequest(ctx context.Context, req *entity.SetFriendAddRequestReq) (*entity.SetFriendaddRequestRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SET_FRIEND_ADD_REQUEST); err != nil {
		return nil, err
	}
	return p.svc.SetFriendAddRequest(ctx, req)
}

func (p *PermissionService) GetLoginInfo(ctx context.Context, req *entity.GetLoginInfoReq) (*entity.GetLoginInfoRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_CONFIG); err != nil {
		return nil, err
	}
	return p.svc.GetLoginInfo(ctx, req)
}

func (p *PermissionService) GetStrangerInfo(ctx context.Context, req *entity.GetStrangerInfoReq) (*entity.GetStrangerInfoRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_CONFIG); err != nil {
		return nil, err
	}
	return p.svc.GetStrangerInfo(ctx, req)
}

func (p *PermissionService) GetFriendList(ctx context.Context, req *entity.GetFriendListReq) (*entity.GetFriendListRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_FRIEND_LIST); err != nil {
		return nil, err
	}
	return p.svc.GetFriendList(ctx, req)
}

func (p *PermissionService) GetGroupInfo(ctx context.Context, req *entity.GetGroupInfoReq) (*entity.GetGroupInfoRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_GROUP_INFO); err != nil {
		return nil, err
	}
	return p.svc.GetGroupInfo(ctx, req)
}

func (p *PermissionService) GetGroupList(ctx context.Context, req *entity.GetGroupListReq) (*entity.GetGroupListRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_GROUP_LIST); err != nil {
		return nil, err
	}
	return p.svc.GetGroupList(ctx, req)
}

func (p *PermissionService) GetGroupMemberInfo(ctx context.Context, req *entity.GetGroupMemberInfoReq) (*entity.GetGroupMemberInfoRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_GROUP_MEMBER_INFO); err != nil {
		return nil, err
	}
	return p.svc.GetGroupMemberInfo(ctx, req)
}

func (p *PermissionService) GetGroupMemberList(ctx context.Context, req *entity.GetGroupMemberListReq) (*entity.GetGroupMemberListRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_GROUP_MEMBER_LIST); err != nil {
		return nil, err
	}
	return p.svc.GetGroupMemberList(ctx, req)
}

func (p *PermissionService) GetGroupHonorInfo(ctx context.Context, req *entity.GetGroupHonorInfoReq) (*entity.GetGroupHonorInfoRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_GROUP_HONOR_INFO); err != nil {
		return nil, err
	}
	return p.svc.GetGroupHonorInfo(ctx, req)
}

func (p *PermissionService) GetCookies(ctx context.Context, req *entity.GetCookiesReq) (*entity.GetCookiesRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_COOKIES); err != nil {
		return nil, err
	}
	return p.svc.GetCookies(ctx, req)
}

func (p *PermissionService) GetCsrfToken(ctx context.Context, req *entity.GetCsrfTokenReq) (*entity.GetCsrfTokenRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_CSRF_TOKEN); err != nil {
		return nil, err
	}
	return p.svc.GetCsrfToken(ctx, req)
}

func (p *PermissionService) GetCredentials(ctx context.Context, req *entity.GetCredentialsReq) (*entity.GetCredentialsRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_CREDENTIALS); err != nil {
		return nil, err
	}
	return p.svc.GetCredentials(ctx, req)
}

func (p *PermissionService) GetRecord(ctx context.Context, req *entity.GetRecordReq) (*entity.GetRecordRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_RECORD); err != nil {
		return nil, err
	}
	return p.svc.GetRecord(ctx, req)
}

func (p *PermissionService) GetImage(ctx context.Context, req *entity.GetImageReq) (*entity.GetImageRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_IMAGE); err != nil {
		return nil, err
	}
	return p.svc.GetImage(ctx, req)
}

func (p *PermissionService) CanSendImage(ctx context.Context, req *entity.CanSendImageReq) (*entity.CanSendImageRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CAN_SEND_IMAGE); err != nil {
		return nil, err
	}
	return p.svc.CanSendImage(ctx, req)
}

func (p *PermissionService) CanSendRecord(ctx context.Context, req *entity.CanSendRecordReq) (*entity.CanSendRecordRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CAN_SEND_RECORD); err != nil {
		return nil, err
	}
	return p.svc.CanSendRecord(ctx, req)
}

func (p *PermissionService) GetStatus(ctx context.Context, req *entity.GetStatusReq) (*entity.GetStatusRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_STATUS); err != nil {
		return nil, err
	}
	return p.svc.GetStatus(ctx, req)
}

func (p *PermissionService) GetVersionInfo(ctx context.Context, req *entity.GetVersionInfoReq) (*entity.GetVersionInfoRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_GET_VERSION_INFO); err != nil {
		return nil, err
	}
	return p.svc.GetVersionInfo(ctx, req)
}

func (p *PermissionService) SetRestart(ctx context.Context, req *entity.SetRestartReq) (*entity.SetRestartRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_SET_RESTART); err != nil {
		return nil, err
	}
	return p.svc.SetRestart(ctx, req)
}

func (p *PermissionService) CleanCache(ctx context.Context, req *entity.CleanCacheReq) (*entity.CleanCacheRsp, error) {
	if err := p.checkPermission(ctx, config.PERMISSION_FOR_CLEAN_CACHE); err != nil {
		return nil, err
	}
	return p.svc.CleanCache(ctx, req)
}

// NewPermissionService 初始化 权限校验
func NewPermissionService(ct *dig.Container, svc service.AdapterServiceServer) (*PermissionService, error) {
	var p PermissionService
	p.svc = svc
	_ = ct.Invoke(func(opt *config.Config) {
		p.conf = opt
	})
	_ = ct.Invoke(func(au *auth.AuthServer) {
		p.auth = au
	})
	return &p, nil
}

func (p *PermissionService) checkPermission(ctx context.Context, permission string) error {
	pl, err := p.auth.GetPayloadFromGrpcContex(ctx)
	if err != nil {
		log.Errorf("checkPermission %s error err:%s", permission, err.Error())
		return err
	}
	appid := pl.JWTID
	ps, err := config.GetPermission(p.conf.Permissions, appid)
	if err != nil {
		log.Errorf("checkPermission %s error err:%s", permission, err.Error())
		return err
	}
	if ps.IsAdmin {
		return nil
	}
	if ps.IsOnlyCqhttp {
		if !config.CheckOnlyCqhttp(permission) {
			log.Errorf("checkPermission onlycqhttp %s error", permission)
			return fmt.Errorf("checkPermission  %s failed err:not onlycqhttp permissions", permission)
		}
		return nil
	}
	if !config.CheckPermission(permission, ps.ApiPermissions) {
		log.Errorf("checkPermission  %s error no permiss", permission)
		return fmt.Errorf("checkPermission  %s failed no permiss", permission)
	}
	return nil
}
