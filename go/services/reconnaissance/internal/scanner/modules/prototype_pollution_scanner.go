package modules

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

func (x *XSSScanner) ScanPrototypePollution(ctx context.Context, url string) {
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url+"?__proto__[polluted]=true"),
		chromedp.Evaluate(`Object.prototype.polluted`, &res),
	)

	if err != nil {
		log.Printf("Error scanning for prototype pollution: %v", err)
		return
	}

	if res == "true" {
		log.Printf("Prototype pollution found on %s", url)
	}
}
