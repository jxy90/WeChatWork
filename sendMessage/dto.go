package sendMessage

type WeChatWorkRequest struct {
	AccessToken        string             `json:"access_token"`
	ToUser             string             `json:"touser"`
	ToParty            string             `json:"toparty"`
	ToTag              string             `json:"totag"`
	MessageType        string             `json:"type"`
	AgentId            string             `json:"agentid"`
	Safe               string             `json:"safe"`
	EnableIdTrans      string             `json:"enable_id_trans"`
	WeChatWorkText     WeChatWorkText     `json:"text"`
	WeChatWorkTextCard WeChatWorkTextCard `json:"textcard"`
	WeChatWorkMarkDown WeChatWorkMarkDown `json:"markdown"`
}
type WeChatWorkText struct {
	Content string `json:"content"`
}
type WeChatWorkImage struct {
	MediaId string `json:"media_id"`
}
type WeChatWorkVoice struct {
	MediaId string `json:"media_id"`
}
type WeChatWorkVideo struct {
	MediaId     string `json:"media_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type WeChatWorkFile struct {
	MediaId string `json:"media_id"`
}

type WeChatWorkTextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Btntxt      string `json:"btntxt"`
}

type WeChatWorkNews struct {
	Articles []WeChatWorkNewsArticle `json:"articles" `
}
type WeChatWorkNewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picUrl"`
}

type WeChatWorkMPNews struct {
	Articles []WeChatWorkMPNewsArticle `json:"articles" `
}
type WeChatWorkMPNewsArticle struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	ContentSourceUrl string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
}

type WeChatWorkMarkDown struct {
	Content string `json:"content" `
}

var WeChatWorkResponse struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	InvalidUser  string `json:"invalidtag"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
}
