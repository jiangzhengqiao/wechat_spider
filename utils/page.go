package utils

import (
	"fmt"
	"github.com/fedesog/webdriver"
	"log"
	"strconv"
)

func GetPageNum(url, platName string, chromeDriver *webdriver.ChromeDriver) int {
	url = fmt.Sprintf(url, platName, 1)

	session, err := chromeDriver.NewSession(webdriver.Capabilities{}, webdriver.Capabilities{})
	if err != nil {
		log.Println(err)
	}

	err = session.Url(url)

	if err != nil {
		log.Println(err)
	}

	var page int = 1
	p, err := session.FindElement(webdriver.ID, "pagebar_container")
	if err != nil {
		return page
	}

	mun, _ := p.FindElement(webdriver.ClassName, "mun")
	scdnum, _ := mun.FindElement(webdriver.ID, "scd_num")
	pageText, _ := scdnum.Text()
	pageCount, _ := strconv.Atoi(pageText)

	if pageCount > 10 {
		page = pageCount/10 + 1
	}
	session.Delete()
	return page
}
