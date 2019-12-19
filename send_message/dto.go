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
	AccessToken   string             `json:"access_token,ommitempty"`
	ToUser        string             `json:"touser,ommitempty"`
	ToParty       string             `json:"toparty,ommitempty"`
	ToTag         string             `json:"totag,ommitempty"`
	MessageType   string             `json:"msgtype,ommitempty"`
	AgentId       int                `json:"agentid,ommitempty"`
	Safe          int                `json:"safe,ommitempty"`
	EnableIdTrans int                `json:"enable_id_trans,ommitempty"`
	Text          WeChatWorkText     `json:"text,ommitempty"`
	Image         WeChatWorkImage    `json:"image,ommitempty"`
	Voice         WeChatWorkVoice    `json:"voice,ommitempty"`
	Video         WeChatWorkVideo    `json:"video,ommitempty"`
	File          WeChatWorkFile     `json:"file,ommitempty"`
	TextCard      WeChatWorkTextCard `json:"textcard,ommitempty"`
	News          WeChatWorkNews     `json:"news,ommitempty"`
	MPNews        WeChatWorkMPNews   `json:"mpnews,ommitempty"`
	MarkDown      WeChatWorkMarkDown `json:"markdown,ommitempty"`
	//MiniProgramNotice WeChatWorkMiniProgramNotice `json:"miniprogram_notice,ommitempty"`
	TaskCard WeChatWorkTaskCard `json:"taskcard,ommitempty"`
}

type WeChatWorkText struct {
	Content string `json:"content,ommitempty"`
}

type WeChatWorkImage struct {
	MediaId string `json:"media_id,ommitempty"`
}

type WeChatWorkVoice struct {
	MediaId string `json:"media_id,ommitempty"`
}

type WeChatWorkVideo struct {
	MediaId     string `json:"media_id,ommitempty"`
	Title       string `json:"title,ommitempty"`
	Description string `json:"description,ommitempty"`
}

type WeChatWorkFile struct {
	MediaId string `json:"media_id,ommitempty"`
}

type WeChatWorkTextCard struct {
	Title       string `json:"title,ommitempty"`
	Description string `json:"description,ommitempty"`
	Url         string `json:"url,ommitempty"`
	Btntxt      string `json:"btntxt,ommitempty"`
}

type WeChatWorkNews struct {
	Articles []WeChatWorkNewsArticle `json:"articles,ommitempty"`
}
type WeChatWorkNewsArticle struct {
	Title       string `json:"title,ommitempty"`
	Description string `json:"description,ommitempty"`
	Url         string `json:"url,ommitempty"`
	PicUrl      string `json:"picurl,ommitempty"`
}

type WeChatWorkMPNews struct {
	Articles []WeChatWorkMPNewsArticle `json:"articles,ommitempty"`
}
type WeChatWorkMPNewsArticle struct {
	Title            string `json:"title,ommitempty"`
	ThumbMediaId     string `json:"thumb_media_id,ommitempty"`
	Author           string `json:"author,ommitempty"`
	ContentSourceUrl string `json:"content_source_url,ommitempty"`
	Content          string `json:"content,ommitempty"`
	Digest           string `json:"digest,ommitempty"`
}

type WeChatWorkMarkDown struct {
	Content string `json:"content,ommitempty"`
}

type WeChatWorkMiniProgramNotice struct {
	AppId             string                            `json:"appid,ommitempty"`
	Page              string                            `json:"page,ommitempty"`
	Title             string                            `json:"title,ommitempty"`
	Description       string                            `json:"description,ommitempty"`
	EmphasisFirstItem bool                              `json:"emphasis_first_item,ommitempty"`
	ContentItem       []WeChatWorkMiniProgramNoticeItem `json:"content_item,ommitempty"`
}

type WeChatWorkMiniProgramNoticeItem struct {
	Key   string `json:"key,ommitempty"`
	Value string `json:"value,ommitempty"`
}

type WeChatWorkTaskCard struct {
	title       string `json:"title,ommitempty"`
	description string `json:"description,ommitempty"`
	url         string `json:"url,ommitempty"`
	task_id     string `json:"task_id,ommitempty"`
	btn         string `json:"btn,ommitempty"`
}
type WeChatWorkTaskCardBtn struct {
	Key         string `json:"key,ommitempty"`
	Name        string `json:"name,ommitempty"`
	ReplaceName string `json:"replace_name,ommitempty"`
	Color       string `json:"color,ommitempty"`
	IsBold      bool   `json:"is_bold,ommitempty"`
}

var WeChatWorkResponse struct {
	ErrCode      int    `json:"errcode,ommitempty"`
	ErrMsg       string `json:"errmsg,ommitempty"`
	InvalidUser  string `json:"invalidtag,ommitempty"`
	InvalidParty string `json:"invalidparty,ommitempty"`
	InvalidTag   string `json:"invalidtag,ommitempty"`
}
