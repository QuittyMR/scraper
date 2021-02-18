# Scraper
Is a straightforward Go web-scraper with a simple, flexible interface, inspired by [BeautifulSoup](https://www.crummy.com/software/BeautifulSoup/bs4/doc/).

## Quickstart
1. Create a `Scraper` from any `io.ReadCloser` compatible type:
   ```
   // http.Response.Body
   response, _ := http.Get("URL goes here")
   page, _ := scraper.NewFromBuffer(response.Body)
   
   // os.File
   fileHandle, _ := os.Open("file name goes here")
   page, _ := NewFromBuffer(fileHandle)
   ```

2. Construct a `Scraper.Filter` with one or more criteria:
   ```
   filter := scraper.Filter{
      Tag: "div",
      Attributes: scraper.Attributes{
         "id":    "div-1",
         "class": "tp-modal",
      },
   }
   ```
3. Use the `Filter` to run a concurrent search on your `Scraper` page.    
   Every returned element is a `Scraper` page that can be searched:
   ```
   for element := range page.FindAll(filter) {
      for link := range element.FindAll(Filter{Tag:"a"}) {
         fmt.Printf("URL: %v found under %v", link.Attributes()["href"], element.Type())
      }
   }
   ```

## Next steps
* ~~Find and FindOne implementations~~
* ~~Concurrent scraping~~
* ~~Resilience for broken pages (BeautifulSoup-esque)~~
* Support for wildcards in attributes
* Tests
* Full documentation
