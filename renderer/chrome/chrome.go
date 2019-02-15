package chrome

import (
	"context"
	"io"
	"log"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/client"
	"github.com/pkg/errors"

	"github.com/harborian/page-render/renderer"
)

type chromeRenderer struct {
	client *chromedp.CDP
	ctx    context.Context
	cancel context.CancelFunc
}

// New - creates new renderer instance
func New() renderer.Renderer {
	ctx, cancel := context.WithCancel(context.Background())

	// create chrome instance
	c, err := chromedp.New(ctx, chromedp.WithTargets(client.New().WatchPageTargets(ctx)), chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	return &chromeRenderer{
		ctx:    ctx,
		client: c,
		cancel: cancel,
	}
}

// Render page
func (r *chromeRenderer) Render(url string) (io.Reader, error) {
	var html string

	err := r.client.Run(r.ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.InnerHTML(`html`, &html, chromedp.ByQuery),
		chromedp.ActionFunc(func(context.Context, cdp.Executor) error {
			return nil
		}),
	})

	if err != nil {
		return nil, err
	}

	log.Printf("Page - %s", html)

	reader := strings.NewReader(html)

	return reader, nil
}

func (r *chromeRenderer) Close() error {
	defer r.cancel()

	err := r.client.Shutdown(r.ctx)
	if err != nil {
		return errors.Wrap(err, "shutdown")
	}

	err = r.client.Wait()
	if err != nil {
		return errors.Wrap(err, "wait closed")
	}

	return nil
}
