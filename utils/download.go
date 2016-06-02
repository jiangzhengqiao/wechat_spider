package utils

import (
	"fmt"
	"github.com/fedesog/webdriver"
	"log"
)

func Download(url, platName string, page int, chromeDriver *webdriver.ChromeDriver) []WeChatList {
	log.Printf("开始采集第%v页，URL【%s】.", page, fmt.Sprintf(url, platName, page))

	results := make([]WeChatList, 0)
	session, err := chromeDriver.NewSession(webdriver.Capabilities{}, webdriver.Capabilities{})
	if err != nil {
		log.Println(err)
	}

	err = session.Url(fmt.Sprintf(url, platName, page))
	if err != nil {
		log.Println(err)
	}

	html, _ := session.FindElement(webdriver.XPath, "//*[@id='main']/div/div[2]/div")

	divs, _ := html.FindElements(webdriver.ClassName, "wx-rb")
	if len(divs) > 0 {
		for _, div := range divs {
			wechatUrl, _ := div.GetAttribute("href")

			val, _ := div.FindElement(webdriver.ClassName, "txt-box")

			h3, _ := val.FindElement(webdriver.TagName, "h3")

			wechatName, _ := h3.Text()

			h4, _ := val.FindElement(webdriver.TagName, "h4")
			h4, _ = h4.FindElement(webdriver.Name, "em_weixinhao")
			wechatSignal, _ := h4.Text()

			sps, _ := val.FindElements(webdriver.ClassName, "s-p3")
			var wechatIntroduced, wechatTit string
			if len(sps) > 1 {
				t, _ := sps[1].FindElement(webdriver.ClassName, "sp-tit")
				th, _ := t.Text()
				if th == "功能介绍：" {
					sptxt, _ := sps[0].FindElement(webdriver.ClassName, "sp-txt")
					wechatIntroduced, _ = sptxt.Text()
				}

				sptxt1, _ := sps[1].FindElement(webdriver.ClassName, "sp-txt")
				wechatTit, _ = sptxt1.Text()
			}

			ico, _ := div.FindElement(webdriver.ClassName, "pos-ico")
			box, _ := ico.FindElement(webdriver.ClassName, "pos-box")
			img, _ := box.FindElement(webdriver.TagName, "img")
			wechatCode, _ := img.GetAttribute("src")

			result := NewWeChatList()
			result.wechatUrl = wechatUrl
			result.wechatName = wechatName
			result.wechatSignal = wechatSignal
			result.wechatIntroduced = wechatIntroduced
			result.wechatTit = wechatTit
			result.wechatCode = wechatCode
			result.articleLists = downloadArticleList(wechatUrl, chromeDriver)
			results = append(results, *result)
		}
	}

	session.Delete()
	return results
}

func downloadArticleList(url string, chromeDriver *webdriver.ChromeDriver) []ArticleList {
	log.Printf("开始采集文章列表页，URL【%s】.", url)
	articleLists := make([]ArticleList, 0)
	session, err := chromeDriver.NewSession(webdriver.Capabilities{}, webdriver.Capabilities{})
	if err != nil {
		log.Println(err)
	}

	err = session.Url(url)
	if err != nil {
		log.Println(err)
	}

	div, _ := session.FindElement(webdriver.ClassName, "weui_msg_card_list")
	infos, _ := div.FindElements(webdriver.ClassName, "weui_msg_card")
	if len(infos) > 0 {
		for _, info := range infos {
			// date, _ := info.FindElement(webdriver.ClassName, "weui_msg_card_hd")
			titles, _ := info.FindElements(webdriver.ClassName, "weui_media_title")
			descs, _ := info.FindElements(webdriver.ClassName, "weui_media_desc")
			infos, _ := info.FindElements(webdriver.ClassName, "weui_media_extra_info")
			urls, _ := info.FindElements(webdriver.TagName, "h4")
			if len(titles) > 0 {
				for i := 0; i < len(titles); i++ {
					articleList := NewArticleList()
					t, _ := titles[i].Text()
					d, _ := descs[i].Text()
					in, _ := infos[i].Text()
					u, _ := urls[i].GetAttribute("hrefs")
					url := "http://mp.weixin.qq.com" + u
					articleList.title = t
					articleList.desc = d
					articleList.date = in
					articleList.url = url
					articleList.articleContent = downloadArticleContent(url, chromeDriver)
					articleLists = append(articleLists, *articleList)
				}
			}
		}
	}
	session.Delete()
	return articleLists
}

func downloadArticleContent(url string, chromeDriver *webdriver.ChromeDriver) ArticleContent {
	session, err := chromeDriver.NewSession(webdriver.Capabilities{}, webdriver.Capabilities{})
	if err != nil {
		log.Println(err)
	}

	err = session.Url(url)
	if err != nil {
		log.Println(err)
	}

	title, _ := session.FindElement(webdriver.ID, "activity-name")
	t, _ := title.Text()

	postDate, _ := session.FindElement(webdriver.ID, "post-date")
	date, _ := postDate.Text()

	imgContent, _ := session.FindElement(webdriver.ID, "img-content")
	content, _ := imgContent.Text()

	articleContent := NewArticleContent()
	articleContent.articleTitle = t
	articleContent.articleDate = date
	articleContent.articleContent = content
	articleContent.articleUrl = url

	session.Delete()
	return *articleContent
}

func DownloadTest(url string, session *webdriver.Session) []WeChatList {
	log.Printf("开始采集URL【%s】.", url)

	results := make([]WeChatList, 0)
	// session, err := chromeDriver.NewSession(webdriver.Capabilities{}, webdriver.Capabilities{})
	// if err != nil {
	// 	log.Println(err)
	// }

	// err = session.Url(fmt.Sprintf(url, platName, page))
	// if err != nil {
	// 	log.Println(err)
	// }

	html, _ := session.FindElement(webdriver.XPath, "//*[@id='main']/div/div[2]/div")

	divs, _ := html.FindElements(webdriver.ClassName, "wx-rb")
	if len(divs) > 0 {
		for _, div := range divs {
			wechatUrl, _ := div.GetAttribute("href")
			val, _ := div.FindElement(webdriver.ClassName, "txt-box")
			h3, _ := val.FindElement(webdriver.TagName, "h3")
			wechatName, _ := h3.Text()
			h4, _ := val.FindElement(webdriver.TagName, "h4")
			h4, _ = h4.FindElement(webdriver.Name, "em_weixinhao")
			wechatSignal, _ := h4.Text()

			sps, _ := val.FindElements(webdriver.ClassName, "s-p3")
			var wechatIntroduced, wechatTit string
			if len(sps) > 1 {
				t, _ := sps[1].FindElement(webdriver.ClassName, "sp-tit")
				th, _ := t.Text()
				if th == "功能介绍：" {
					sptxt, _ := sps[0].FindElement(webdriver.ClassName, "sp-txt")
					wechatIntroduced, _ = sptxt.Text()
				}

				sptxt1, _ := sps[1].FindElement(webdriver.ClassName, "sp-txt")
				wechatTit, _ = sptxt1.Text()
			}

			ico, _ := div.FindElement(webdriver.ClassName, "pos-ico")
			box, _ := ico.FindElement(webdriver.ClassName, "pos-box")
			img, _ := box.FindElement(webdriver.TagName, "img")
			wechatCode, _ := img.GetAttribute("src")
			div.Click()

			// 采集文章列表
			articleLists := make([]ArticleList, 0)
			wh, _ := session.WindowHandle()
			whs, _ := session.WindowHandles()
			for _, v := range whs {
				if wh.ID != v.ID {
					session.FocusOnWindow(v.ID)
					div, _ := session.FindElement(webdriver.ClassName, "weui_msg_card_list")
					infos, _ := div.FindElements(webdriver.ClassName, "weui_msg_card")
					if len(infos) > 0 {
						for _, info := range infos {
							// date, _ := info.FindElement(webdriver.ClassName, "weui_msg_card_hd")
							titles, _ := info.FindElements(webdriver.ClassName, "weui_media_title")
							descs, _ := info.FindElements(webdriver.ClassName, "weui_media_desc")
							infos, _ := info.FindElements(webdriver.ClassName, "weui_media_extra_info")
							urls, _ := info.FindElements(webdriver.TagName, "h4")
							if len(titles) > 0 {
								for i := 0; i < len(titles); i++ {
									articleList := NewArticleList()
									t, _ := titles[i].Text()
									d, _ := descs[i].Text()
									in, _ := infos[i].Text()
									u, _ := urls[i].GetAttribute("hrefs")
									url := "http://mp.weixin.qq.com" + u
									titles[i].Click()
									// 采集文章
									wh, _ := session.WindowHandle()
									whs, _ := session.WindowHandles()
									articleContent := NewArticleContent()
									// log.Println("窗口个数：\t", len(whs))
									for _, v := range whs {
										if wh.ID != v.ID {
											session.FocusOnWindow(v.ID)
											title, err := session.FindElement(webdriver.ID, "activity-name")
											if err == nil {
												t, _ := title.Text()

												postDate, _ := session.FindElement(webdriver.ID, "post-date")
												date, _ := postDate.Text()

												imgContent, _ := session.FindElement(webdriver.ID, "img-content")
												content, _ := imgContent.Text()

												articleContent.articleTitle = t
												articleContent.articleDate = date
												articleContent.articleContent = content
												articleContent.articleUrl = url
											}
											// session.CloseCurrentWindow()
											session.Back()
											session.FocusOnWindow(wh.ID)
										}
									}

									articleList.title = t
									articleList.desc = d
									articleList.date = in
									articleList.url = url
									articleList.articleContent = *articleContent
									articleLists = append(articleLists, *articleList)
								}
							}
						}
					}
					session.CloseCurrentWindow()
					session.FocusOnWindow(wh.ID)
				}
			}

			result := NewWeChatList()
			result.wechatUrl = wechatUrl
			result.wechatName = wechatName
			result.wechatSignal = wechatSignal
			result.wechatIntroduced = wechatIntroduced
			result.wechatTit = wechatTit
			result.wechatCode = wechatCode
			result.articleLists = articleLists
			results = append(results, *result)
		}
	}
	return results
}
