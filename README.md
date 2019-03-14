# sitemap.xml Crawler

**sitemapbot** crawls URLs on `sitemap.xml`.

## Install

`go install github.com/hareku/sitemapbot/cmd/sitemapbot`

## Usage

`sitemapbot [sitemap URLs...]`

```
GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Example

```
$ sitemapbot https://example.com/sitemap.xml
2019/03/03 19:02:41 crawling started: https://example.com/sitemap.xml
 50 / 50 [=================================] 100.00% 11s
2019/03/03 19:02:53 crawling finished: https://example.com/sitemap.xml, success:50, error:0
```
