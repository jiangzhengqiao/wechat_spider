package utils

// 公众号列表
type WeChatList struct {
	wechatUrl        string        // 微信地址
	wechatName       string        // 微信名
	wechatSignal     string        // 微信号
	wechatIntroduced string        // 功能介绍
	wechatTit        string        // 微信认证
	wechatCode       string        // 二维码
	articleLists     []ArticleList // 文章列表
}

// 文章列表
type ArticleList struct {
	title          string         // 标题
	desc           string         // 摘要
	date           string         // 时间
	url            string         // 地址
	articleContent ArticleContent // 文章
}

// 文章
type ArticleContent struct {
	articleTitle   string // 标题
	articleDate    string // 时间
	articleUrl     string // 地址
	articleContent string // 内容
	articleHtml    string // 原始内容
}

func NewWeChatList() *WeChatList {
	return &WeChatList{}
}

func NewArticleList() *ArticleList {
	return &ArticleList{}
}

func NewArticleContent() *ArticleContent {
	return &ArticleContent{}
}

// func (this *WeChatList) GetArticleList() []ArticleList {
// 	// 根据URL采集文章列表 this.wechatUrl
// }
