package utils

import (
	"log"
)

func Pipeline(results []WeChatList) {
	if len(results) > 0 {
		var i int = 1
		for _, result := range results {
			log.Println("1,-----------------------------------------")
			log.Println("微信地址:", result.wechatUrl)
			log.Println("微信名:", result.wechatName)
			log.Println("微信号:", result.wechatSignal)
			log.Println("微信认证:", result.wechatTit)
			log.Println("功能介绍:", result.wechatIntroduced)
			log.Println("二维码地址:", result.wechatCode)

			if len(result.articleLists) > 0 {
				for _, articleList := range result.articleLists {
					log.Println("2,-----------------------------------------")
					log.Println("标题:", articleList.title)
					log.Println("摘要:", articleList.desc)
					log.Println("时间:", articleList.date)
					log.Println("链接地址:", articleList.url)

					// if len(articleList.articleContents) > 0 {
					// 	for _, articleContent := range articleList.articleContents {
					log.Println("3,-----------------------------------------")
					log.Println("标题:", articleList.articleContent.articleTitle)
					log.Println("内容:", articleList.articleContent.articleContent)
					log.Println("时间:", articleList.articleContent.articleDate)
					log.Println("地址:", articleList.articleContent.articleUrl)
					i++
					// 	}
					// }
				}
			}
			log.Printf("微信号%s,有文章%v篇。", result.wechatName, i)
		}
	}
}
