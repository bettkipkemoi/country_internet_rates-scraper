Scraping the cost of internet across different countries from worldpopulationreview.com

- make a directory and cd into it
`mkdir myproject && cd myproject`

- initialize the go directory
`go mod init github.com/username/myproject`

- install go/colly
`go get github.com/gocolly/colly`

- create your scraper
[main.go](/main.go)

- run your scraper
you can run through either:
    - go run .
    - go run main.go
    - go build ./myproject