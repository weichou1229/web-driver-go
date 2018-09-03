package main

import (
	"fmt"
	"github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	"log"
	"testing"
)

func TestDriver(t *testing.T) {
	gomega.RegisterTestingT(t)
	fmt.Println("Start !")

	goToTarget()

}

func TestDriver2(t *testing.T) {
	fmt.Println("Start !")
	for i := 1; i < 10000; i++ {
		testByPhantomJS(i)
	}
}

func testByPhantomJS(i int) {
	// init driver
	var driver *agouti.WebDriver = agouti.PhantomJS()

	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start Selenium:", err)
	}

	testTarget(i, driver)
	stopDriver(driver)
}

func TestDriver3(t *testing.T) {
	fmt.Println("Start !")
	for i := 1; i < 10000; i++ {
		testByChromeDriver(i)
	}
}

func testByChromeDriver(i int) {
	// init driver
	chormOpetion1 := agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox"})
	//chormOpetion2 := agouti.ChromeOptions("prefs", "{'intl.accept_languages': 'en,es'}")
	var driver *agouti.WebDriver = agouti.ChromeDriver(chormOpetion1)

	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start Selenium:", err)
	}

	testTarget(i, driver)
	stopDriver(driver)
}
