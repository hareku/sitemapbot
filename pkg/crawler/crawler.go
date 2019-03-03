package crawler

import (
	"net/http"
	"sync"

	sm "github.com/yterajima/go-sitemap"
	"golang.org/x/xerrors"
	"gopkg.in/cheggaaa/pb.v1"
)

// Run sitemap.xml crawling.
func Run(sitemapURL string) (successN int, errorN int, err error) {
	var successNLock, errorNLock sync.Mutex

	sitemap, err := sm.Get(sitemapURL, nil)
	if err != nil {
		err = xerrors.Errorf("error url %q: %w", sitemapURL, err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(sitemap.URL))

	bar := pb.StartNew(len(sitemap.URL))

	for _, URL := range sitemap.URL {
		go func(loc string) {
			defer wg.Done()
			defer bar.Increment()

			isSuccess := Access(loc)
			if isSuccess {
				successNLock.Lock()
				defer successNLock.Unlock()
				successN++
			} else {
				errorNLock.Lock()
				defer errorNLock.Unlock()
				errorN++
			}
		}(URL.Loc)
	}

	wg.Wait()
	bar.Finish()

	return
}

// Access a URL, and return whether http.Get() is success or not.
func Access(url string) bool {
	_, err := http.Get(url)
	return err == nil
}
