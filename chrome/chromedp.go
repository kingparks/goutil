package chrome

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

func GetResult(url, wait string) (html string) {
	userAgent := "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.37"
	allocContext, _ := chromedp.NewExecAllocator(
		context.Background(),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-default-browser-check", true),
		chromedp.Flag("disable-plugins", true),
		chromedp.Flag("ignore-certificate-errors", "1"),
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(userAgent),
	)
	ctx, cancel := chromedp.NewContext(
		allocContext,
	)
	defer cancel()

	// 设置timeout
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(wait),
		chromedp.OuterHTML(`document.querySelector("html")`, &html, chromedp.ByJSPath),
	)
	if err != nil {
		return
	}
	return
}
