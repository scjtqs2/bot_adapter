package adapter

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/scjtqs2/bot_adapter/config"
	"github.com/scjtqs2/bot_adapter/pb/entity"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"time"
)

// HttpAdapter http方法的结构体
type HttpAdapter struct {
	conf      *config.Config
	EventChan chan []byte
}

func (h *HttpAdapter) CustomSetGroupPortrait(req *entity.CustomSetGroupPortraitReq) (rsp *entity.CustomSetGroupPortraitRsp, err error) {
	url := config.API_CUSTOM_URL_SET_GROUP_PORTRAIT
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) CustomGetWordSlices(req *entity.CustomGetWordSlicesReq) (rsp *entity.CustomGetWordSlicesRsp, err error) {
	url := config.API_CUSTOM_URL_GET_WORD_SLICES
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomOcrImage(req *entity.CustomOcrImageReq) (rsp *entity.CustomOcrImageRsp, err error) {
	url := config.API_CUSTOM_URL_OCR_IMAGE
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomGetGroupSystemMsg(req *entity.CustomGetGroupSystemMsgReq) (rsp *entity.CustomGetGroupSystemMsgRsp, err error) {
	url := config.API_CUSTOM_URL_GET_GROUP_SYSTEM_MSG
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomUploadGroupFile(req *entity.CustomUploadGroupFileReq) (rsp *entity.CustomUploadGroupFileRsp, err error) {
	url := config.API_CUSTOM_URL_UPLOAD_GROUP_FILE
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) CustomGetGroupFileSystemInfo(req *entity.CustomGetGroupFileSystemInfoReq) (rsp *entity.CustomGetGroupFileSystemInfoRsp, err error) {
	url := config.API_CUSTOM_URL_GET_GROUP_FILE_SYSTEM_INFO
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomGetGroupRootFiles(req *entity.CustomGetGroupRootFilesReq) (rsp *entity.CustomGetGroupRootFilesRsp, err error) {
	url := config.API_CUSTOM_URL_GET_GROUP_ROOT_FILES
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) CustomGetGroupFilesByFolder(req *entity.CustomGetGroupFilesByFolderReq) (rsp *entity.CustomGetGroupFilesByFolderRsp, err error) {
	url := config.API_CUSTOM_URL_GET_GROUP_FILES_BY_FOLDER
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) CustomGetGroupFileUrl(req *entity.CustomGetGroupFileUrlReq) (rsp *entity.CustomGetGroupFileUrlRsp, err error) {
	url := config.API_CUSTOM_URL_GET_GROUP_FILE_URL
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) CustomGetStatus(req *entity.CustomGetStatusReq) (rsp *entity.CustomGetStatusRsp, err error) {
	url := config.API_URL_GET_STATUS
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomGetGroupAtAllRemain(req *entity.CustomGetGroupAtAllRemainReq) (rsp *entity.CustomGetGroupAtAllRemainRsp, err error) {
	url := config.API_CUSTOM_URL_GET_GROUP_AT_ALL_REMAIN
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomGetVipInfo(req *entity.CustomGetVipInfoReq) (rsp *entity.CustomGetVipInfoRsp, err error) {
	url := config.API_CUSTOM_URL_GET_VIP_INFO
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomSendGroupNotice(req *entity.CustomSendGroupNoticeReq) (rsp *entity.CustomSendGroupNoticeRsp, err error) {
	url := config.API_CUSTOM_URL_SEND_GROUP_NOTICE
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) CustomReloadEventFilter(req *entity.CustomReloadEventFilterReq) (rsp *entity.CustomReloadEventFilterRsp, err error) {
	url := config.API_CUSTOM_URL_RELOAD_EVENT_FILTER
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) CustomDownloadFile(req *entity.CustomDownloadFileReq) (rsp *entity.CustomDownloadFileRsp, err error) {
	url := config.API_CUSTOM_URL_DOWNLOAD_FILE
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) CustomGetOnlineClinets(req *entity.CustomGetOnlineClientsReq) (rsp *entity.CustomGetOnlineClientsRsp, err error) {
	url := config.API_CUSTOM_URL_GET_ONLINE_CLIENTS
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomGetGroupMsgHistory(req *entity.CustomGetGroupMsgHistoryReq) (rsp *entity.CustomGetGroupMsgHistoryRsp, err error) {
	url := config.API_CUSTOM_URL_GET_GROUP_MSG_HISTORY
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomSetEssenceMsg(req *entity.CustomSetEssenceMsgReq) (rsp *entity.CustomSetEssenceMsgRsp, err error) {
	url := config.API_CUSTOM_URL_SET_ESSENCE_MSG
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) CustomDeleteEssenceMsg(req *entity.CustomDeleteEssenceMsgReq) (rsp *entity.CustomDeleteEssenceMsgRsp, err error) {
	url := config.API_CUSTOM_URL_DELETE_ESSENCE_MSG
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) CustomGetEssenceMsgList(req *entity.CustomGetEssenceMsgListReq) (rsp *entity.CustomGetEssenceMsgListRsp, err error) {
	url := config.API_CUSTOM_URL_GET_ESSENCE_MSG_LIST
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomCheckUrlSafely(req *entity.CustomCheckUrlSafelyReq) (rsp *entity.CustomCheckUrlSafelyRsp, err error) {
	url := config.API_CUSTOM_URL_CHECK_URL_SAFELY
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomGetModelShow(req *entity.CustomGetModelShowReq) (rsp *entity.CustomGetModelShowRsp, err error) {
	url := config.API_CUSTOM_URL_GET_MODEL_SHOW
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomSetModelShow(req *entity.CustomSetModelShowReq) (rsp *entity.CustomSetModelShowRsp, err error) {
	url := config.API_CUSTOM_URL_SET_MODEL_SHOW
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) CustomGetMsg(req *entity.CustomGetMsgReq) (rsp *entity.CustomGetMsgRsp, err error) {
	url := config.API_CUSTOM_URL_GET_MSG
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomGetImage(req *entity.CustomGetImageReq) (rsp *entity.CustomGetImageRsp, err error) {
	url := config.API_CUSTOM_URL_GET_IMAGE
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomSendGroupForwardMsg(req *entity.CustomSendGroupForwardMsgReq) (rsp *entity.CustomSendGroupForwardMsgRsp, err error) {
	url := config.API_CUSTOM_URL_SEND_GROUP_FORWARD_MSG
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CustomGetForwardMsg(req *entity.CustomGetForwardMsgReq) (rsp *entity.CustomGetForwardMsgRsp, err error) {
	url := config.API_CUSTOM_URL_GET_FORWARD_MSG
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

// SendPrivateMsg 发送私聊
func (h *HttpAdapter) SendPrivateMsg(req *entity.SendPrivateMsgReq) (rsp *entity.SendMsgRsp, err error) {
	url := config.API_URL_SEND_PRIVATE_MSG
	msg := MSG{
		"user_id":     req.GetUserId(),
		"auto_escape": req.GetAutoEscape(),
		"message":     string(req.GetMessage()),
	}
	if req.GetIsArrayMessage() {
		message := gjson.ParseBytes(req.GetMessage())
		msg["message"] = message.Map()
	}
	res, err := h.Send(url, msg)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

// SendGroupMsg 发送群消息
func (h *HttpAdapter) SendGroupMsg(req *entity.SendGroupMsgReq) (rsp *entity.SendMsgRsp, err error) {
	url := config.API_URL_SEND_GROUP_MSG
	msg := MSG{
		"group_id":    req.GetGroupId(),
		"auto_escape": req.GetAutoEscape(),
		"message":     string(req.GetMessage()),
	}
	if req.GetIsArrayMessage() {
		message := gjson.ParseBytes(req.GetMessage())
		msg["message"] = message.Map()
	}
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

// SendMsg 发送消息
func (h *HttpAdapter) SendMsg(req *entity.SendMsgReq) (rsp *entity.SendMsgRsp, err error) {
	url := config.API_URL_SEND_MSG
	msg := MSG{
		"user_id":      req.GetUserId(),
		"group_id":     req.GetUserId(),
		"message_type": req.GetMessageType(),
		"auto_escape":  req.GetAutoEscape(),
		"message":      string(req.GetMessage()),
	}
	if req.GetIsArrayMessage() {
		message := gjson.ParseBytes(req.GetMessage())
		msg["message"] = message.Map()
	}
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

// DeleteMsg 撤回消息
func (h *HttpAdapter) DeleteMsg(req *entity.DeleteMsgReq) (rsp *entity.DeleteMsgRsp, err error) {
	url := config.API_URL_DELETE_MSG
	_, err = h.Send(url, req)
	return nil, err
}

// GetMsg 拉取消息
func (h *HttpAdapter) GetMsg(req *entity.GetMsgReq) (rsp *entity.GetMsgRsp, err error) {
	url := config.API_URL_GET_MSG
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetForwardMsg(req *entity.GetForwardMsgReq) (rsp *entity.GetForwardMsgRsp, err error) {
	url := config.API_URL_GET_FORWARD_MSG
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) SendLike(req *entity.SendLikeReq) (rsp *entity.SendLikeRsp, err error) {
	url := config.API_URL_SEND_LIKE
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) SetGroupKick(req *entity.SetGroupKickReq) (rsp *entity.SetGroupKickRsp, err error) {
	url := config.API_URL_SET_GROUP_KICK
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) SetGroupBan(req *entity.SetGroupBanReq) (rsp *entity.SetGroupBanRsp, err error) {
	url := config.API_URL_SET_GROUP_BAN
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) SetGroupAnonymousBan(req *entity.SetGroupAnonymousBanReq) (rsp *entity.SetGroupAnonymousBanRsp, err error) {
	url := config.API_URL_SET_GROUP_ANONYMOUS_BAN
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) SetGroupWholeBan(req *entity.SetGroupWholeBanReq) (rsp *entity.SetGroupWholeBanRsp, err error) {
	url := config.PERMISSION_FOR_SET_GROUP_WHOLE_BAN
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) SetGroupAdmin(req *entity.SetGroupAdminReq) (rsp *entity.SetGroupAdminRsp, err error) {
	url := config.API_URL_SET_GROUP_ADMIN
	_, err = h.Send(url, req)
	return rsp, err
}

func (h *HttpAdapter) SetGroupAnonymous(req *entity.SetGroupAnonymousReq) (rsp *entity.SetGroupAnonymousRsp, err error) {
	url := config.API_URL_SET_GROUP_ANONYMOUS
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) SetGroupCard(req *entity.SetGroupCardReq) (rsp *entity.SetGroupCardRsp, err error) {
	url := config.API_URL_SET_GROUP_CARD
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) SetGroupName(req *entity.SetGroupNameReq) (rsp *entity.SetGroupNameRsp, err error) {
	url := config.API_URL_SET_GROUP_NAME
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) SetGroupLeave(req *entity.SetGroupLeaveReq) (rsp *entity.SetGroupLeaveRsp, err error) {
	url := config.API_URL_SET_GROUP_LEAVE
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) SetGroupSpecialTitle(req *entity.SetGroupSpecialTitleReq) (rsp *entity.SetGroupSpecialTitleRsp, err error) {
	url := config.API_URL_SET_GROUP_SPECIAL_TITLE
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) SetGroupAddRequest(req *entity.SetGroupAddRequestReq) (rsp *entity.SetGroupAddRequestRsp, err error) {
	url := config.API_URL_SET_GROUP_ADD_REQUEST
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) SetFriendAddRequest(req *entity.SetFriendAddRequestReq) (rsp *entity.SetFriendaddRequestRsp, err error) {
	url := config.API_URL_SET_FRIEND_ADD_REQUEST
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) GetLoginInfo(req *entity.GetLoginInfoReq) (rsp *entity.GetLoginInfoRsp, err error) {
	url := config.API_URL_GET_LOGIN_INFO
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetStrangerInfo(req *entity.GetStrangerInfoReq) (rsp *entity.GetStrangerInfoRsp, err error) {
	url := config.API_URL_GET_STRANGER_INFO
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetFriendList(req *entity.GetFriendListReq) (rsp *entity.GetFriendListRsp, err error) {
	url := config.API_URL_GET_FRIEND_LIST
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetGroupInfo(req *entity.GetGroupInfoReq) (rsp *entity.GetGroupInfoRsp, err error) {
	url := config.API_URL_GET_GROUP_INFO
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetGroupList(req *entity.GetGroupListReq) (rsp *entity.GetGroupListRsp, err error) {
	url := config.API_URL_GET_GROUP_LIST
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetGroupMemberInfo(req *entity.GetGroupMemberInfoReq) (rsp *entity.GetGroupMemberInfoRsp, err error) {
	url := config.API_URL_GET_GROUP_MEMBER_INFO
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetGroupMemberList(req *entity.GetGroupMemberListReq) (rsp *entity.GetGroupMemberListRsp, err error) {
	url := config.API_URL_GET_GROUP_MEMBER_LIST
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetGroupHonorInfo(req *entity.GetGroupHonorInfoReq) (rsp *entity.GetGroupHonorInfoRsp, err error) {
	url := config.API_URL_GET_GROUP_HONOR_INFO
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetCookies(req *entity.GetCookiesReq) (rsp *entity.GetCookiesRsp, err error) {
	url := config.API_URL_GET_COOKIES
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetCsrfToken(req *entity.GetCsrfTokenReq) (rsp *entity.GetCsrfTokenRsp, err error) {
	url := config.API_URL_GET_CSRF_TOKEN
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetCredentials(req *entity.GetCredentialsReq) (rsp *entity.GetCredentialsRsp, err error) {
	url := config.API_URL_GET_CREDENTIALS
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetRecord(req *entity.GetRecordReq) (rsp *entity.GetRecordRsp, err error) {
	url := config.API_URL_GET_RECORD
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetImage(req *entity.GetImageReq) (rsp *entity.GetImageRsp, err error) {
	url := config.API_URL_GET_IMAGE
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CanSendImage(req *entity.CanSendImageReq) (rsp *entity.CanSendImageRsp, err error) {
	url := config.API_URL_CAN_SEND_IMAGE
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) CanSendRecord(req *entity.CanSendRecordReq) (rsp *entity.CanSendRecordRsp, err error) {
	url := config.API_URL_CAN_SEND_RECORD
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetStatus(req *entity.GetStatusReq) (rsp *entity.GetStatusRsp, err error) {
	url := config.API_URL_GET_STATUS
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) GetVersionInfo(req *entity.GetVersionInfoReq) (rsp *entity.GetVersionInfoRsp, err error) {
	url := config.API_URL_GET_VERSION_INFO
	res, err := h.Send(url, req)
	_ = json.Unmarshal(res, &rsp)
	return rsp, err
}

func (h *HttpAdapter) SetRestart(req *entity.SetRestartReq) (rsp *entity.SetRestartRsp, err error) {
	url := config.API_URL_SET_RESTART
	_, err = h.Send(url, req)
	return nil, err
}

func (h *HttpAdapter) CleanCache(req *entity.CleanCacheReq) (rsp *entity.CleanCacheRsp, err error) {
	url := config.API_URL_CLEAN_CACHE
	_, err = h.Send(url, req)
	return nil, err
}

// Send 通过Http渠道发送数据到 cqhttp
func (h *HttpAdapter) Send(action string, msg interface{}) ([]byte, error) {
	var res *http.Response
	body, _ := json.Marshal(msg)
	client := http.Client{Timeout: time.Second * time.Duration(time.Second*2)}
	header := make(http.Header)
	header.Set("Authorization", fmt.Sprintf("Bearer %s", h.conf.HTTPConfig.Token))
	header.Set("Content-Type", "application/json")
	maxRetries := 3
	for i := 0; i <= maxRetries; i++ {
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", h.conf.HTTPConfig.ServerAddr, action), bytes.NewReader(body))
		if err != nil {
			log.Warnf("发送数据到cqhttp %s 时创建请求失败: %v", h.conf.HTTPConfig.ServerAddr, err)
			return nil, err
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
			log.Warnf("发送数据到cqhttp %s 失败: %v 将进行低 %d 次重试", h.conf.HTTPConfig.ServerAddr, err, i+1)
		} else {
			log.Warnf("发送数据到cqhttp 数据 %s 到 %v 失败: %v 停止发送：已达重试上线", body, h.conf.HTTPConfig.ServerAddr, err)
			return nil, err
		}
	}
	r, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if gjson.ValidBytes(r) {
		result := gjson.ParseBytes(r)
		switch result.Get("retcode").Int() {
		case 0:
			return []byte(result.Get("data").String()), nil
		default:
			if result.Get("msg").String() != "" {
				return nil, fmt.Errorf(fmt.Sprintf("%s|%s", result.Get("msg").String(), result.Get("wording").String()))
			}
			return nil, errors.New("failed")
		}
	}
	return nil, errors.New("err respnsed")
}

// GetChan 拿到对应的chan，用来读取收到的消息
func (h *HttpAdapter) GetChan() chan []byte {
	return h.EventChan
}

// Init 初始化 http server，用于接收来自http的post
func (h *HttpAdapter) Init() {
	app := iris.New()
	app.Post("/msginput", h.msginput)
	go func() {
		_ = app.Run(iris.Addr(h.conf.HTTPConfig.LocalHost))
	}()
}

func (h *HttpAdapter) msginput(ctx iris.Context) {
	raw, _ := ctx.GetBody()
	if err := h.checkSignature(ctx, raw); err != nil {
		_, _ = ctx.JSON(MSG{
			"code": -1,
			"msg":  "signatreu err",
		})
		return
	}
	// fmt.Printf("post body: %s \r\n", string(raw))
	_, _ = ctx.JSON(MSG{
		"code": 200,
		"msg":  "received",
	})
	h.EventChan <- raw // 把接收到的消息写入chan内
}

// checkSignature 校验签名
func (h *HttpAdapter) checkSignature(ctx iris.Context, rawData []byte) error {

	signature := ctx.GetHeader("X-Signature")
	if signature != "" {
		signature = signature[5:]
	}
	mac := hmac.New(sha1.New, []byte(h.conf.HTTPConfig.Secret))
	_, _ = mac.Write(rawData)
	signatrueNow := hex.EncodeToString(mac.Sum(nil))
	// log.Infof("signature=%s ; signatureNow:%s", signature, signatrueNow)
	if signature != signatrueNow {
		return errors.New("signatreu err")
	}
	return nil
}

// HttpInit 初始化http的方法
func HttpInit(cfg *config.Config) (AdapterInterface, error) {
	adpt := &HttpAdapter{
		EventChan: make(chan []byte, 200),
	}
	err := adpt.InitConfig(cfg)
	if err != nil {
		return nil, err
	}
	adpt.Init()
	return adpt, nil
}

// InitConfig 加载配置文件
func (h *HttpAdapter) InitConfig(opt *config.Config) error {
	if opt.HTTPConfig.LocalHost == "" {
		log.Error("error addr")
		return errors.New("error addr")
	}
	if opt.HTTPConfig.LocalHost == "" {
		log.Error("error host")
		return errors.New("error host")
	}
	h.conf = opt
	return nil
}
