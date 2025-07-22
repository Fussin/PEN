package browser

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

type Chromedp struct{}

func NewChromedp() *Chromedp {
	return &Chromedp{}
}

func (c *Chromedp) Screenshot(url string) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(3*time.Second),
		chromedp.FullScreenshot(&buf, 90),
	)
	return buf, err
}

func (c *Chromedp) FuzzDOMEvents(url, selector string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(2 * time.Second),
		chromedp.Focus(selector),
		chromedp.Click(selector),
		chromedp.MouseOver(selector),
	}
	return chromedp.Run(ctx, tasks)
}

func (c *Chromedp) CheckPayloadExecution(url string) (bool, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var alertFired bool
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if _, ok := ev.(*chromedp.EventDialogOpening); ok {
			alertFired = true
		}
	})

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(2*time.Second),
	)
	return alertFired, err
}
