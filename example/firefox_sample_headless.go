// main.go
package main

import (
	"log"

	"github.com/sclevine/agouti"
)

func main() {
	driver := agouti.Selenium()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("firefox"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	for i := 0; i < 20; i++ {
		if err := page.Navigate("http://www.yahoo.co.jp/"); err != nil {
			log.Fatalf("Failed to navigate:%v", err)
		}
		page.ClearCookies()
	}
}
