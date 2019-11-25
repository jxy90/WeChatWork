package send_message

import (
	"encoding/json"
	"errors"
	"github.com/pangpanglabs/goutils/behaviorlog"
	"github.com/pangpanglabs/goutils/httpreq"
	"net/http"
)

//https://work.weixin.qq.com/api/doc#90000/90135/90236
func Send(request WeChatWorkRequest) error {
	switch request.MessageType {
	case MsgType_Text,
		MsgType_Image,
		MsgType_Voice,
		MsgType_Video,
		MsgType_File,
		MsgType_TextCard,
		MsgType_News,
		MsgType_MPNews,
		MsgType_MarkDown,
		MsgType_TaskCard:
	default:
		errors.New("invalid MessageType:" + request.MessageType)
	}
	senMsgUrl := WeChatWorkURL + "/send?access_token=" + request.AccessToken
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return errors.New("struct Marshal err")
	}
	requestMap := make(map[string]interface{}, 0)
	err = json.Unmarshal(requestBytes, &requestMap)
	if err != nil {
		return errors.New("struct Unmarshal err")
	}
	_, err = httpreq.New(http.MethodPost, senMsgUrl, requestMap).
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
