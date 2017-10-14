package main

import (
	"log"

	"github.com/sclevine/agouti"
)

// START OMIT
func main() {
	options := agouti.Desired(agouti.Capabilities{
		"platformName":    "iOS",
		"platformVersion": "11.0",
		"automationName":  "XCUITest",
		"deviceName":      "iPhone X",
		"browserName":     "Safari",
		"language":        "ja",
		"maxInstances":    1,
		"platform":        "MAC",
	})
	command := []string{"cd", "."}        // 特に何もしないコマンドを入れています
	url := "http://127.0.0.1:4723/wd/hub" // Appiumに接続
	driver := agouti.NewWebDriver(url, command, options)
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()
	// END OMIT
	page, err := driver.NewPage(agouti.Browser("safari"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	if err := page.SetImplicitWait(10); err != nil {
		log.Fatalf("Failed to timeout:%v", err)
	}

	for i := 0; i < 10; i++ {
		if err := page.Navigate("http://www.yahoo.co.jp/"); err != nil {
			log.Fatalf("Failed to navigate:%v", err)
		}
		page.ClearCookies()
	}
}
