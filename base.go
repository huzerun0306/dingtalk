package dingtalk

type msgTypeType string

const (
	TEXT       msgTypeType = "text"
	LINK       msgTypeType = "link"
	MARKDOWN   msgTypeType = "markdown"
	ACTION_CARD msgTypeType = "actionCard"
	FEED_CARD   msgTypeType = "feedCard"
)

type textModel struct {
	Content string `json:"content,omitempty"`
}

type atModel struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

type linkModel struct {
	Text       string `json:"text,omitempty"`
	Title      string `json:"title,omitempty"`
	PicUrl     string `json:"picUrl,omitempty"`
	MessageUrl string `json:"messageUrl,omitempty"`
}

type markDownModel struct {
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
}

type actionCardBtnOrientationType string

const (
	horizontal actionCardBtnOrientationType = "0" // 横向
	vertical   actionCardBtnOrientationType = "1" // 竖向
)

type actionCardModel struct {
	Title          string                       `json:"title,omitempty"`
	Text           string                       `json:"text,omitempty"`
	BtnOrientation actionCardBtnOrientationType `json:"btnOrientation,omitempty"`
	SingleTitle    string                       `json:"singleTitle,omitempty"`
	SingleURL      string                       `json:"singleURL,omitempty"`
	Btns           []actionCardMultiBtnModel    `json:"btns,omitempty"`
}

type actionCardMultiBtnModel struct {
	Title     string `json:"title,omitempty"`
	ActionURL string `json:"actionURL,omitempty"`
}

type feedCardModel struct {
	Links []feedCardLinkModel `json:"links,omitempty"`
}

type feedCardLinkModel struct {
	Title      string `json:"title,omitempty"`
	MessageURL string `json:"messageURL,omitempty"`
	PicURL     string `json:"picURL,omitempty"`
}