package modules

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

func (x *XSSScanner) ScanDOMXSS(ctx context.Context, url string) {
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Evaluate(`
			const sinks = {
				'document.write': [0],
				'document.writeln': [0],
				'innerHTML': [0],
				'outerHTML': [0],
				'eval': [0],
				'setTimeout': [1],
				'setInterval': [1],
			};

			let tainted = location.hash.substring(1);
			let results = [];

			for (const sink in sinks) {
				const sinkFunc = eval(sink);
				if (typeof sinkFunc === 'function') {
					const originalFunc = sinkFunc;
					eval(sink + ` = function(...args) {
						for (const pos of sinks[sink]) {
							if (args.length > pos && args[pos] === tainted) {
								results.push(sink);
							}
						}
						return originalFunc.apply(this, args);
					}`);
				}
			}

			// Trigger some events to try and activate the sinks
			document.dispatchEvent(new Event('DOMContentLoaded'));
			window.dispatchEvent(new Event('load'));

			JSON.stringify(results);
		`, &res),
	)

	if err != nil {
		log.Printf("Error scanning for DOM XSS: %v", err)
		return
	}

	log.Printf("DOM XSS scan results: %s", res)
}
