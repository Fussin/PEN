package modules

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"time"
)

func captureScreenshotWithChromedp(url string) []byte {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(3*time.Second),
		chromedp.FullScreenshot(&buf, 90),
	)
	if err != nil {
		return []byte{}
	}
	return buf
}

func triggerDOMEvents(ctx context.Context, selector string) error {
	tasks := chromedp.Tasks{
		chromedp.Focus(selector),
		chromedp.Click(selector),
		chromedp.MouseOver(selector),
	}
	return chromedp.Run(ctx, tasks)
}

func fuzzInteraction(url, payload string) bool {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var foundSelector string
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *chromedp.EventDialogOpening:
			if ev.Message == "XSS" || ev.Message == "1337" {
				foundSelector = "detected"
			}
		}
	})

	var nodes []*cdp.Node
	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(500 * time.Millisecond),
		chromedp.Nodes(`//*[contains(text(), "`+payload+`")]`, &nodes, chromedp.AtLeast(0)),
	}
	if err := chromedp.Run(ctx, tasks); err != nil {
		return false
	}

	for _, node := range nodes {
		if node.FullXPath != "" {
			triggerDOMEvents(ctx, node.FullXPath)
		}
	}
	return foundSelector == "detected"
}
