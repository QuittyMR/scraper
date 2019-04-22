# Scraper
Is a straightforward Go web-scraper with a simple, flexible interface, inspired by [BeautifulSoup](https://www.crummy.com/software/BeautifulSoup/bs4/doc/).

## Quickstart
* Search for CSS imports using a URL:
```
myScraper, _ := scraper.NewFromURI("Your URL goes here")
parameters := scraper.Parameters{"rel": "stylesheet"}
fmt.Println(myScraper.FindAll(scraper.Filters{Tag:"link", Parameters:parameters}))
```

* Implement your own http request and search the response for tables:
```
response, _ := http.Get(url)
myScraper, _ := scraper.NewFromResponse(response)
fmt.Println(myScraper.FindAll(scraper.Filters{Tag:"table"}))
```

* Search for all headers of all tables having a certain class:
```
myScraper, _ := scraper.NewFromURI("Your URL goes here")
parameters := scraper.Parameters{"class":"someClass"}
for _, table := range myScraper.FindAll(scraper.Filters{Tag:"table", Parameters:parameters}) {
    fmt.Println(table.FindAll(scraper.Filters{Tag:"th"}))
}
```

* Render the HTML of Github's code-block:
```
myScraper, _ := scraper.NewFromURI("Some Github code page")
myScraper.Content()
```

## Next steps
* ~~Find and FindOne implementations~~
* Concurrent scraping
* Tests
* Resilience for broken pages (BeautifulSoup-esque)
