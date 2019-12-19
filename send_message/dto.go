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
	AccessToken   string             `json:"access_token,omitempty"`
	ToUser        string             `json:"touser,omitempty"`
	ToParty       string             `json:"toparty,omitempty"`
	ToTag         string             `json:"totag,omitempty"`
	MessageType   string             `json:"msgtype,omitempty"`
	AgentId       int                `json:"agentid,omitempty"`
	Safe          int                `json:"safe,omitempty"`
	EnableIdTrans int                `json:"enable_id_trans,omitempty"`
	Text          WeChatWorkText     `json:"text,omitempty"`
	Image         WeChatWorkImage    `json:"image,omitempty"`
	Voice         WeChatWorkVoice    `json:"voice,omitempty"`
	Video         WeChatWorkVideo    `json:"video,omitempty"`
	File          WeChatWorkFile     `json:"file,omitempty"`
	TextCard      WeChatWorkTextCard `json:"textcard,omitempty"`
	News          WeChatWorkNews     `json:"news,omitempty"`
	MPNews        WeChatWorkMPNews   `json:"mpnews,omitempty"`
	MarkDown      WeChatWorkMarkDown `json:"markdown,omitempty"`
	//MiniProgramNotice WeChatWorkMiniProgramNotice `json:"miniprogram_notice,omitempty"`
	TaskCard WeChatWorkTaskCard `json:"taskcard,omitempty"`
}

type WeChatWorkText struct {
	Content string `json:"content,omitempty"`
}

type WeChatWorkImage struct {
	MediaId string `json:"media_id,omitempty"`
}

type WeChatWorkVoice struct {
	MediaId string `json:"media_id,omitempty"`
}

type WeChatWorkVideo struct {
	MediaId     string `json:"media_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type WeChatWorkFile struct {
	MediaId string `json:"media_id,omitempty"`
}

type WeChatWorkTextCard struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
	Btntxt      string `json:"btntxt,omitempty"`
}

type WeChatWorkNews struct {
	Articles []WeChatWorkNewsArticle `json:"articles,omitempty"`
}
type WeChatWorkNewsArticle struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
	PicUrl      string `json:"picurl,omitempty"`
}

type WeChatWorkMPNews struct {
	Articles []WeChatWorkMPNewsArticle `json:"articles,omitempty"`
}
type WeChatWorkMPNewsArticle struct {
	Title            string `json:"title,omitempty"`
	ThumbMediaId     string `json:"thumb_media_id,omitempty"`
	Author           string `json:"author,omitempty"`
	ContentSourceUrl string `json:"content_source_url,omitempty"`
	Content          string `json:"content,omitempty"`
	Digest           string `json:"digest,omitempty"`
}

type WeChatWorkMarkDown struct {
	Content string `json:"content,omitempty"`
}

type WeChatWorkMiniProgramNotice struct {
	AppId             string                            `json:"appid,omitempty"`
	Page              string                            `json:"page,omitempty"`
	Title             string                            `json:"title,omitempty"`
	Description       string                            `json:"description,omitempty"`
	EmphasisFirstItem bool                              `json:"emphasis_first_item,omitempty"`
	ContentItem       []WeChatWorkMiniProgramNoticeItem `json:"content_item,omitempty"`
}

type WeChatWorkMiniProgramNoticeItem struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type WeChatWorkTaskCard struct {
	title       string `json:"title,omitempty"`
	description string `json:"description,omitempty"`
	url         string `json:"url,omitempty"`
	task_id     string `json:"task_id,omitempty"`
	btn         string `json:"btn,omitempty"`
}
type WeChatWorkTaskCardBtn struct {
	Key         string `json:"key,omitempty"`
	Name        string `json:"name,omitempty"`
	ReplaceName string `json:"replace_name,omitempty"`
	Color       string `json:"color,omitempty"`
	IsBold      bool   `json:"is_bold,omitempty"`
}

var WeChatWorkResponse struct {
	ErrCode      int    `json:"errcode,omitempty"`
	ErrMsg       string `json:"errmsg,omitempty"`
	InvalidUser  string `json:"invalidtag,omitempty"`
	InvalidParty string `json:"invalidparty,omitempty"`
	InvalidTag   string `json:"invalidtag,omitempty"`
}
