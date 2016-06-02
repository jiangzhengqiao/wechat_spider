package main

import (
	"github.com/fedesog/webdriver"
	"log"
	"time"

	"weixin.com/utils"
)

func main() {
	chromeDriver := webdriver.NewChromeDriver(utils.CHROME_DRIVER_PATH)
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}

	session, err := chromeDriver.NewSession(webdriver.Capabilities{}, webdriver.Capabilities{})
	if err != nil {
		log.Println(err)
	}

	// 打开网站
	err = session.Url("http://shuju.wdzj.com")

	if err != nil {
		log.Println(err)
	}

	// 测试URL
	tables, _ := session.FindElements(webdriver.ClassName, "data_table")
	trs, _ := tables[1].FindElements(webdriver.TagName, "tr")
	if len(trs) > 0 {
		for _, tr := range trs {
			tds, _ := tr.FindElements(webdriver.TagName, "td")
			if len(tds) != 6 {
				log.Println("数据列表页显示不完整。")
			} else {
				as, _ := tds[1].FindElements(webdriver.TagName, "a")
				as[0].Click()
				// platName, _ := as[0].Text()
				wh, _ := session.WindowHandle()
				whs, _ := session.WindowHandles()
				// log.Println("窗口个数：\t", len(whs))
				for _, v := range whs {
					if wh.ID != v.ID {
						session.FocusOnWindow(v.ID)
						title, _ := session.Title()
						url, _ := session.GetUrl()
						if title == "" {
							log.Println(url)
						}

						// if url == "http://shuju.wdzj.com/plat-info-0.html" {
						// 	log.Printf("%s 存在问题，请检查档案后台。", platName)
						// }
						// 获取完信息之后再切换回去 fack，找了一下午
						log.Println(title)
						// time.Sleep(2 * time.Second)
						session.CloseCurrentWindow()
						session.FocusOnWindow(wh.ID)
					}
				}
			}
		}
	} else {
		log.Println("数据页面无数据。")
	}

	time.Sleep(10 * time.Second)
	session.Delete()
	chromeDriver.Stop()
}
