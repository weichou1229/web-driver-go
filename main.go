package main

import (
	"flag"
	"fmt"
	"github.com/sclevine/agouti"
	"log"
)

const (
	CHROME          = "chrome"
	CHROME_HEADLESS = "chrome_headless"
	PHANTOMJS       = "phantomJS"
)

func stopDriver(driver *agouti.WebDriver) {
	if err := driver.Stop(); err != nil {
		log.Fatal("Failed to close pages and stop WebDriver:", err)
	}
}

func initDriver(driverType string) *agouti.WebDriver {
	var driver *agouti.WebDriver
	switch driverType {
	case CHROME:
		driver = agouti.ChromeDriver()
	case CHROME_HEADLESS:
		driver = agouti.ChromeDriver(agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox"}))
	case PHANTOMJS:
		driver = agouti.PhantomJS()
	default:
		driver = agouti.PhantomJS()
	}

	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start Selenium:", err)
	}

	return driver
}

func initPage(driver *agouti.WebDriver) *agouti.Page {
	var url = "https://www.google.com.tw/search?q=胖護士+減肥"

	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("Failed to open page:", err)
	}

	if err := page.Navigate(url); err != nil {
		log.Fatal("Failed to navigate:", err)
	}

	return page
}

func initPageWithPageUrl(driver *agouti.WebDriver, pageUrl string) *agouti.Page {
	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("Failed to open page:", err)
	}

	if err := page.Navigate(pageUrl); err != nil {
		log.Fatal("Failed to navigate:", err)
	}

	return page
}

func main() {
	var driverType = flag.String("driverType", PHANTOMJS, "the web driver type")
	flag.Parse()

	fmt.Println("Start ! ", *driverType)

	goToTarget(*driverType)
}

func goToPage(driverType string) {
	for i := 0; i < 1000; i++ {
		var driver = initDriver(driverType)

		pageUrl := "http://weiwei0923.pixnet.net/blog/post/299526739-%E5%A6%82%E5%A6%82"
		page := initPageWithPageUrl(driver, pageUrl)

		//title, _ := page.FindByXPath("//*[@id='article-299526739']/h2/a").Text()

		log.Printf("[%v] Open taget", i)

		//
		page.CloseWindow()

		stopDriver(driver)
	}
}

func goToTarget(driverType string) {
	for i := 0; i < 1000; i++ {
		var driver = initDriver(driverType)
		page := initPageWithPageUrl(driver, "https://www.google.com.tw/search?q=胖護士+減肥")

		var urls = []string{
			"http://weiwei0923.pixnet.net/blog/post/299526739-如如",
			"http://weiwei0923.pixnet.net/blog/post/314385229-最重要的決定",
			"http://weiwei0923.pixnet.net/blog/post/299525662-規劃婚前減重目標，三個月成功達陣",
		}

		for _, url := range urls {
			if err := page.Navigate(url); err != nil {
				log.Fatal("Failed to navigate:", err)
			}
		}

		//var target = page.Find("a[href*='http://weiwei0923.pixnet.net/blog/post/299526739-%E5%A6%82%E5%A6%82']")
		//
		//target.Click()

		//title, _ := page.FindByXPath("//*[@id='article-299526739']/h2/a").Text()
		//
		//log.Printf("[%v] Open taget %v", i, title)
		log.Printf("[%v] Open taget", i)
		//
		page.CloseWindow()

		stopDriver(driver)
	}
}

func testTarget(i int, driver *agouti.WebDriver) {
	// init page
	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("Failed to open page:", err)
	}

	// navigate page
	var url = "http://weiwei0923.pixnet.net/blog/post/299526739-%E5%A6%82%E5%A6%82"
	if err := page.Navigate(url); err != nil {
		log.Fatal("Failed to navigate:", err)
	}
	log.Printf("Open url %v", url)
	title, _ := page.FindByXPath("//*[@id='article-299526739']/h2/a").Text()

	log.Printf("[%v] Open taget %v", i, title)

	page.CloseWindow()
}
