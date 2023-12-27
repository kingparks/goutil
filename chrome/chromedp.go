package chrome

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

func GetResult(remoteChrome, url, wait string) (html string, err error) {
	userAgent := "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.37"
	allocContext, cancel := chromedp.NewExecAllocator(
		context.Background(),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-default-browser-check", true),
		chromedp.Flag("disable-plugins", true),
		chromedp.Flag("ignore-certificate-errors", "1"),
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(userAgent),
	)
	if remoteChrome != "" {
		allocContext, cancel = chromedp.NewRemoteAllocator(
			allocContext, remoteChrome,
		)
	}
	allocContext, cancel = chromedp.NewContext(
		allocContext,
	)
	defer cancel()

	allocContext, cancel = context.WithTimeout(allocContext, 65*time.Second)
	defer cancel()

	if wait == "" {
		err = chromedp.Run(allocContext,
			chromedp.Navigate(url),
			chromedp.OuterHTML(`document.querySelector("html")`, &html, chromedp.ByJSPath),
		)
	} else {
		err = chromedp.Run(allocContext,
			chromedp.Navigate(url),
			chromedp.WaitVisible(wait),
			chromedp.OuterHTML(`document.querySelector("html")`, &html, chromedp.ByJSPath),
		)
	}
	return
}
