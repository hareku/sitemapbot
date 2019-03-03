package main

import (
	"log"
	"os"

	sitemapbot "github.com/hareku/sitemapbot/pkg/crawler"
	"github.com/urfave/cli"
	"golang.org/x/xerrors"
)

func main() {
	app := cli.NewApp()

	app.Name = "sitemapbot"
	app.Usage = "This tool crawls URLs on sitemap.xml"
	app.Version = "0.1.1"
	app.Action = appAction

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("sitemapbot error: %+v", err)
	}
}

func appAction(context *cli.Context) error {
	if context.NArg() == 0 {
		return xerrors.New("must pass sitemap.xml URLs to arguments")
	}

	for _, sitemapURL := range context.Args() {
		log.Printf("crawling started: %s", sitemapURL)

		successN, errorN, err := sitemapbot.Run(sitemapURL)
		if err != nil {
			return xerrors.Errorf("crawling error: %w", err)
		}

		log.Printf("crawling finished: %s, success:%d, error:%d", sitemapURL, successN, errorN)
	}

	return nil
}
