package main

import (
  "fmt"
  "log"
  "os"
  "strings"
  "github.com/PuerkitoBio/goquery"
)

func RelScrape(url string) {
  doc, err := goquery.NewDocument(url) 
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("URL: %s\r\n", url)
  fmt.Printf("Title: %s\r\n", doc.Find("head > title").Text()) //doc.Find("title").Text()

  doc.Find("link[rel]").Each(func(i int, s *goquery.Selection) {
    rel, _ := s.Attr("rel")
    url, _ := s.Attr("href")
    fmt.Printf("Rel: %s %s\r\n", rel, url)
  })

  doc.Find("link[rev]").Each(func(i int, s *goquery.Selection) {
    rev, _ := s.Attr("rev")
    url, _ := s.Attr("href")
    fmt.Printf("Rel: %s %s\r\n", rev, url)
  })

  doc.Find("meta[name]").Each(func(i int, s *goquery.Selection) {
    name, _ := s.Attr("name")
    content, _ := s.Attr("content")
    fmt.Printf("Meta: %s %s\r\n", name, content)
  })

  doc.Find("meta[property]").Each(func(i int, s *goquery.Selection) {
    property, _ := s.Attr("property")
    content, _ := s.Attr("content")
    fmt.Printf("Meta: %s %s\r\n", property, content)
  })

  doc.Find("*[itemscope]").Each(func(i int, s *goquery.Selection) {
    itemtype, _ := s.Attr("itemtype")
    fmt.Printf("Schema: %s\r\n", itemtype)

    //hack!
    a := s.Find("a").First()
    href, _ := a.Attr("href")
    fmt.Printf("\thref %s\r\n", href)

    s.Find("link[itemprop]").Each(func(i int, s *goquery.Selection) {
      itemprop, _ := s.Attr("itemprop")
      href, _ := s.Attr("href")
      fmt.Printf("\t%s %s\r\n", itemprop, href)
    })

    //sometimes doesn't get the image because data-original="http://media2.s-nbcnews.com/j/msnbc/components/video/__new/150303-barbershop-chung-vid-tease2.nbcnews-fp-320-240.jpg" instead of src=""
    s.Find("img[itemprop]").Each(func(i int, s *goquery.Selection) {
      itemprop, _ := s.Attr("itemprop")
      src, _ := s.Attr("src")
      if(src == "") {
        src="missingUrl"
      }
      alt, _ := s.Attr("alt")
      fmt.Printf("\t%s %s %s\r\n", itemprop, src, alt)
    })

    s.Find("time[itemprop]").Each(func(i int, s *goquery.Selection) {
      itemprop, _ := s.Attr("itemprop")
      datetime, _ := s.Attr("datetime")
      fmt.Printf("\t%s %s\r\n", itemprop, datetime)
    })

    s.Find("*[itemprop]").Each(func(i int, s *goquery.Selection) {
      itemprop, _ := s.Attr("itemprop")
      itemval := strings.TrimSpace(s.Text())
      if(itemval != "") {
        fmt.Printf("\t%s %s\r\n", itemprop, itemval)
      }
    })

  })
}

func main() {
  if len(os.Args) != 2 {
     fmt.Printf("Usage : %s URL \n ", os.Args[0]) 
     os.Exit(1) 
  }

  RelScrape(os.Args[1])
}
