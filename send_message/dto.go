package send_message

const (
	MsgType_Text     = "text"
	MsgType_Image    = "image"
	MsgType_Voice    = "voice"
	MsgType_Video    = "video"
	MsgType_File     = "file"
	MsgType_TextCard = "textcard"
	MsgType_News     = "news"
	MsgType_MPNews   = "mpnews"
	MsgType_MarkDown = "markdown"
	//MsgType_MiniProgramNotice = "miniprogram_notice"
	MsgType_TaskCard = "taskcard"
)

type WeChatWorkRequest struct {
	AccessToken   string             `json:"access_token"`
	ToUser        string             `json:"touser"`
	ToParty       string             `json:"toparty"`
	ToTag         string             `json:"totag"`
	MessageType   string             `json:"msgtype"`
	AgentId       int                `json:"agentid"`
	Safe          int                `json:"safe"`
	EnableIdTrans int                `json:"enable_id_trans"`
	Text          WeChatWorkText     `json:"text"`
	Image         WeChatWorkImage    `json:"image"`
	Voice         WeChatWorkVoice    `json:"voice"`
	Video         WeChatWorkVideo    `json:"video"`
	File          WeChatWorkFile     `json:"file"`
	TextCard      WeChatWorkTextCard `json:"textcard"`
	News          WeChatWorkNews     `json:"news"`
	MPNews        WeChatWorkMPNews   `json:"mpnews"`
	MarkDown      WeChatWorkMarkDown `json:"markdown"`
	//MiniProgramNotice WeChatWorkMiniProgramNotice `json:"miniprogram_notice"`
	TaskCard WeChatWorkTaskCard `json:"taskcard"`
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
	PicUrl      string `json:"picurl"`
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

type WeChatWorkMiniProgramNotice struct {
	AppId             string                            `json:"appid" `
	Page              string                            `json:"page" `
	Title             string                            `json:"title" `
	Description       string                            `json:"description" `
	EmphasisFirstItem bool                              `json:"emphasis_first_item" `
	ContentItem       []WeChatWorkMiniProgramNoticeItem `json:"content_item" `
}

type WeChatWorkMiniProgramNoticeItem struct {
	Key   string `json:"key" `
	Value string `json:"value" `
}

type WeChatWorkTaskCard struct {
	title       string `json:"title" `
	description string `json:"description" `
	url         string `json:"url" `
	task_id     string `json:"task_id" `
	btn         string `json:"btn" `
}
type WeChatWorkTaskCardBtn struct {
	Key         string `json:"key" `
	Name        string `json:"name" `
	ReplaceName string `json:"replace_name" `
	Color       string `json:"color" `
	IsBold      bool   `json:"is_bold" `
}

var WeChatWorkResponse struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	InvalidUser  string `json:"invalidtag"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
}
