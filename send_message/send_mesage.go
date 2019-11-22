package send_message

import (
	"encoding/json"
	"errors"
	"github.com/pangpanglabs/goutils/behaviorlog"
	"github.com/pangpanglabs/goutils/httpreq"
	"net/http"
)

const (
	WeChatMsgType_Text     = "text" //文本
	WeChatMsgType_Image    = "image"
	WeChatMsgType_Voice    = "voice"
	WeChatMsgType_Video    = "video"
	WeChatMsgType_File     = "file"
	WeChatMsgType_TextCatd = "textcard" //文本卡片
	WeChatMsgType_MarkDown = "markdown" //markdown消息
)

//https://work.weixin.qq.com/api/doc#90000/90135/90236
func Send(request WeChatWorkRequest) error {
	senMsgUrl := WeChatWorkURL + "/send?access_token=" + request.AccessToken
	msg := make(map[string]interface{}, 0)
	detail := make(map[string]interface{}, 0)
	msg["touser"] = request.ToUser
	msg["toparty"] = request.ToParty
	msg["totag"] = request.ToParty
	msg["msgtype"] = request.MessageType
	msg["agentid"] = request.AgentId
	msg["safe"] = 0
	msg["enable_id_trans"] = 0

	switch request.MessageType {
	case WeChatMsgType_Text:
		detail = makeDetailMsgOfText(request.WeChatWorkText)
	case WeChatMsgType_TextCatd:
		detail = makeDetailMsgOfTextCard(request.WeChatWorkTextCard)
	case WeChatMsgType_MarkDown:
		detail = makeDetailMsgOfMarkDown(request.WeChatWorkMarkDown)
	default:
		return errors.New("miss MsgType")
	}

	msg[request.MessageType] = detail
	_, err := httpreq.New(http.MethodPost, senMsgUrl, msg).
		WithBehaviorLogContext(behaviorlog.FromCtx(nil)).
		Call(&WeChatWorkResponse)
	if err != nil {
		return err
	}
	if WeChatWorkResponse.ErrCode == 0 && WeChatWorkResponse.ErrMsg == "ok" {
		return nil
	}
	errmsg, _ := json.Marshal(WeChatWorkResponse)
	return errors.New("send WeChatWork message error. " + string(errmsg))
}

func makeDetailMsgOfText(textCard WeChatWorkText) map[string]interface{} {
	detail := make(map[string]interface{}, 0)
	detail["content"] = textCard.Content
	return detail
}

func makeDetailMsgOfTextCard(textCard WeChatWorkTextCard) map[string]interface{} {
	detail := make(map[string]interface{}, 0)
	detail["title"] = textCard.Title
	detail["description"] = textCard.Description
	detail["url"] = textCard.Url
	detail["btntxt"] = textCard.Btntxt
	return detail
}

func makeDetailMsgOfMarkDown(markDown WeChatWorkMarkDown) map[string]interface{} {
	detail := make(map[string]interface{}, 0)
	detail["content"] = markDown.Content
	return detail
}
