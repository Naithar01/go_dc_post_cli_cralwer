package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Post struct {
	Id     string
	Title  string
	Writer string
}

var (
	URL string = "https://gall.dcinside.com/board/lists/?id=programming&page=%d"
)

func CrawlerSite(page int, docChan chan *goquery.Document) error {
	url := fmt.Sprintf(URL, page)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return err
	}

	docChan <- doc

	return nil
}

func main() {
	// 한 페이지에 게시글은 총 51개
	var posts []Post
	docChan := make(chan *goquery.Document)

	go func() {
		for page := 1; page <= 22; page++ {
			CrawlerSite(page, docChan)
		}
		close(docChan)
	}()

	for doc := range docChan {
		doc.Find("tr.ub-content").Each(func(i int, s *goquery.Selection) {
			id := s.Find("td.gall_num").Text()
			title := s.Find("td.gall_tit > a").Text()
			writer := s.Find("td.gall_writer").AttrOr("data-nick", "ㅇㅇ")
			post := Post{Id: id, Title: title, Writer: writer}
			posts = append(posts, post)
		})
	}

	fmt.Println(posts)
}
