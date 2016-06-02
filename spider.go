package main

import (
	"fmt"
	"github.com/fedesog/webdriver"
	"log"
	"weixin.com/utils"
)

func main() {
	platName := "小企业E家"
	chromeDriver := webdriver.NewChromeDriver(utils.CHROME_DRIVER_PATH)
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}

	url := "http://weixin.sogou.com/weixin?query=%s&type=1&ie=utf8"

	session, err := chromeDriver.NewSession(webdriver.Capabilities{}, webdriver.Capabilities{})
	if err != nil {
		log.Println(err)
	}

	url = fmt.Sprintf(url, platName)
	err = session.Url(url)

	if err != nil {
		log.Println(err)
	}

	// pageNum := utils.GetPageNum(url, platName, chromeDriver)

	// if pageNum > 1 {
	// for i := 1; i <= pageNum; i++ {
	results := utils.DownloadTest(url, session)
	utils.Pipeline(results)
	// }
	// } else {
	// 	results := utils.Download(url, platName, 1, chromeDriver)
	// 	utils.Pipeline(results)
	// }

	session.Delete()

	// fmt.Println(results)
}
