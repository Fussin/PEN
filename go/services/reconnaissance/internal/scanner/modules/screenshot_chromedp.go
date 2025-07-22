package modules

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
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
	}
	return chromedp.Run(ctx, tasks)
}

func fuzzInteraction(url, payload string) bool {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var foundSelector string
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev.(type) {
		case *runtime.EventExceptionThrown:
			foundSelector = "detected"
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
		if node.FullXPath() != "" {
			triggerDOMEvents(ctx, node.FullXPath())
		}
	}
	return foundSelector == "detected"
}

func simulatePayloadExecution(url string, payload string) bool {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var alertFired bool
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if _, ok := ev.(*runtime.EventExceptionThrown); ok {
			alertFired = true
		}
	})

	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(2 * time.Second),
	}

	err := chromedp.Run(ctx, tasks)
	if err != nil {
		return false
	}
	return alertFired
}

func captureElementScreenshot(url, sel string) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Screenshot(sel, &buf, chromedp.NodeVisible),
	)
	return buf, err
}
